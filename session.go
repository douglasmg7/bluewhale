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

// User id from each session, different sessions can point to same user id.
var userIdMap = map[string]int{}

// Measure execution time to retrive the session.
// var timeToGetSession time.Time

// Keep data that is retrive on every request when user is logged.
// So keep it small and cached.
// Every time some data into db that is keeped in the cache is changed,
// the session data must be deleted, to be update on new request.
type SessionData struct {
  UserId   int
  UserName string
}

// Session data from user id.
var sessionDataMap = map[int]SessionData{}

// Writing a cookie on client and keep a reletion of cookie -> user id.
func NewSession(w http.ResponseWriter, userId int) error {
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
  userIdMap[sUUIDString] = userId
  return nil
}

// Remove the session.
func RemoveSession(w http.ResponseWriter, req *http.Request) error {
  c, err := req.Cookie("sessionUUID")
  // log.Println("cookie:", c)
  // log.Println("cookie-err:", err)
  // No cookie.
  if err != nil && err != http.ErrNoCookie {
    // log.Println("No cookie")
    return err
  } else {
    // Delete cookie.
    c.MaxAge = -1
    log.Println("changed cookie:", c)
    http.SetCookie(w, c)
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return nil
  }
}

// Retrive session data.
func GetSessionData(req *http.Request) (sData *SessionData, err error) {
  // timeToGetSession = time.Now()
  userId, err := userIdfromSessionUUID(req)
  // Some error.
  if err != nil {
    return nil, err
    // No user id.
  } else if userId == 0 {
    return nil, nil
    // Found user.
  } else {
    sData, err := sessionDataFromUserId(userId)
    if err != nil {
      log.Fatal(err)
    }
    // log.Println("Time to get session:", time.Since(timeToGetSession))
    return sData, err
    // return sessionDataFromUserId(userId)
  }
}

// Get user id from session uuid.
// Try the cache first.
func userIdfromSessionUUID(req *http.Request) (int, error) {
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
    userId := userIdMap[sessionUUID]
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
        userIdMap[sessionUUID] = userId
        return userId, nil
      }
    }
  }
  // No cookie
  return 0, nil
}

// Session data from cache.
// If not cached, cache it.
func sessionDataFromUserId(userId int) (sData *SessionData, err error) {
  sDateTemp := sessionDataMap[userId]
  sData = &sDateTemp
  if sData.UserName != "" {
    // log.Println("Get from cache:", sData.UserName)
    return sData, nil
  } else {
    return cacheSession(userId)
  }
}

// Cache session data and return it.
// func cacheSession(userId int) (sData *SessionData, err error) {
func cacheSession(userId int) (*SessionData, error) {
  // Get data from db(s).
  var sData SessionData
  err := db.QueryRow("select id, name from user where id = ?", userId).Scan(&sData.UserId, &sData.UserName)
  // log.Println("Retrive from db:", sData.UserName)
  if err != nil {
    return &sData, err
  }
  // Cache it.
  if sData.UserName != "" {
    sessionDataMap[userId] = sData
  }
  return &sData, nil
}
