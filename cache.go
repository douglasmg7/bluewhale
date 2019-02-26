// todo - concurret access to session map can cause incosistence
package main

import (
	// _ "github.com/mattn/go-sqlite3"
	// "github.com/satori/go.uuid"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"time"
)

// Measure execution time to retrive the session.
// var timeToGetSession time.Time

// Data need for each session.
type Session struct {
	UserId   int
	UserName string
}

// Cached data.
type Cache struct {
	// UserId from uuidSession.
	userIds map[string]int
	// Session from userId.
	sessions map[int]Session
}

// Create a session, writing a cookie on client and keep a reletion of cookie -> user id.
func (cache *Cache) CreateSession(w http.ResponseWriter, userId int) error {
	// create cookie
	sUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}
	sUUIDString := sUUID.String()
	// Save cookie.
	http.SetCookie(w, &http.Cookie{
		Name:  "sessionUUID",
		Value: sUUIDString,
		Path:  "/",
		// Secure: true, // to use only in https
		// HttpOnly: true, // Can't be used into js client
	})
	// Save session UUID on db.
	stmt, err := db.Prepare(`INSERT INTO sessionUUID(uuid, user_id, created) VALUES( ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sUUIDString, userId, time.Now().String())
	if err != nil {
		return err
	}
	// Save on cache.
	cache.userIds[sUUIDString] = userId
	return nil
}

// Remove the session.
func (cache *Cache) RemoveSession(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("sessionUUID")
	// No cookie.
	if err == http.ErrNoCookie {
		// log.Println("No cookie")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		// Some error.
	} else if err != nil {
		log.Fatal(err)
		// Remove cookie.
	} else {
		c.MaxAge = -1
		c.Path = "/"
		// log.Println("changed cookie:", c)
		http.SetCookie(w, c)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		// Delete userId session.
		delete(cache.userIds, c.Value)
	}
}

// Get session.
func (cache *Cache) GetSession(req *http.Request) (*Session, error) {
	// timeToGetSession = time.Now()
	userId, err := cache.getUserIdfromSessionUUID(req)
	// Some error.
	if err != nil {
		return nil, err
		// No user id.
	} else if userId == 0 {
		return nil, nil
		// Found user.
	} else {
		session, err := cache.getSessionFromUserId(userId)
		if err != nil {
			log.Fatal(err)
		}
		// log.Println("Time to get session:", time.Since(timeToGetSession))
		return session, err
		// return sessionDataFromUserId(userId)
	}
}

// Get user id from session uuid.
// Try the cache first.
func (cache *Cache) getUserIdfromSessionUUID(req *http.Request) (int, error) {
	cookie, err := req.Cookie("sessionUUID")
	// log.Println("Cookie:", cookie.Value)
	// log.Println("Cookie-err:", err)
	// No cookie.
	if err == http.ErrNoCookie {
		return 0, nil
		// some error
	} else if err != nil {
		return 0, err
	}
	// Have a cookie.
	if cookie != nil {
		sessionUUID := cookie.Value
		userId := cache.userIds[sessionUUID]
		// Found on cache.
		if userId != 0 {
			// log.Println("userId from cache", userId)
			return userId, nil
		} else {
			// Get from db.
			err = db.QueryRow("select user_id from sessionUUID where uuid = ?", sessionUUID).Scan(&userId)
			if err != nil {
				// No user id for the sessionUUID.
				return 0, err
			}
			// Found the user id.
			if userId != 0 {
				// log.Println("userId from db", userId)
				cache.userIds[sessionUUID] = userId
				return userId, nil
			}
		}
	}
	// No cookie
	return 0, nil
}

// Get session from cache.
// If not cached, cache it.
func (cache *Cache) getSessionFromUserId(userId int) (*Session, error) {
	sessionTemp := cache.sessions[userId]
	session := &sessionTemp
	if session.UserName != "" {
		// log.Println("Get from cache:", session.UserName)
		return session, nil
	} else {
		return cache.cacheSession(userId)
	}
}

// Cache session data and return it.
func (cache *Cache) cacheSession(userId int) (*Session, error) {
	// Get data from db(s).
	var session Session
	err := db.QueryRow("select id, name from user where id = ?", userId).Scan(&session.UserId, &session.UserName)
	// log.Println("Retrive from db:", session.UserName)
	if err != nil {
		return nil, err
	}
	// Cache it.
	if session.UserName != "" {
		cache.sessions[userId] = session
	}
	return &session, nil
}
