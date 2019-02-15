// todo - concurret access to session map can cause incosistence
package main

import (
  // _ "github.com/mattn/go-sqlite3"
  // "github.com/satori/go.uuid"
  // "log"
  "github.com/satori/go.uuid"
  "net/http"
  "time"
)

// User id from each session, different sessions can point to same user id.
var userIdMap = map[string]int{}

// Keep data that is retrive on every request when user is logged.
// So keep it small and cached.
// Every time some data into db that is keeped in the cache is changed,
// the session data must be deleted, to be update on new request.
type SessionData struct {
  UserId int
  Name   string
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
  // save cookie
  http.SetCookie(w, &http.Cookie{
    Name:  "sessionUUID",
    Value: sUUIDString,
    // Secure: true, // to use only in https
    HttpOnly: true, // Can't be used into js client
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

// Retrive session data.
func GetSessionData(req *http.Request) (sData *SessionData, err error) {
  userId, err := userIdfromSessionUUID(req)
  // Some error.
  if err != nil {
    return nil, err
    // No user id.
  } else if userId == 0 {
    return nil, nil
    // Found user.
  } else {
    return sessionDataFromUserId(userId)
  }
}

// Get user id from session uuid.
// Try the cache first.
func userIdfromSessionUUID(req *http.Request) (int, error) {
  cookie, err := req.Cookie("session")
  // No cookie.
  if err == http.ErrNoCookie {
    return 0, nil
    // some error
  } else if err != nil {
    return 0, err
  }
  // Have a cookie.
  if cookie != nil {
    sessionUUID := cookie.String()
    userId := userIdMap[sessionUUID]
    // Found on cache.
    if userId != 0 {
      return userId, nil
    } else {
      // Get from db.
      err = db.QueryRow("select user_id from sessionUUID where uuid = ?", cookie.String()).Scan(&userId)
      if err != nil {
        // No user id for the sessionUUID.
        return 0, err
      }
      // Found the user id.
      if userId != 0 {
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
  if sData.Name != "" {
    return sData, nil
  } else {
    return cacheSession(userId)
  }
}

// Cache session data and return it.
func cacheSession(userId int) (sData *SessionData, err error) {
  // Get data from db(s).
  err = db.QueryRow("select name from user where id = ?", userId).Scan(&sData.Name)
  if err != nil {
    return sData, err
  }
  // Cache it.
  if sData.Name != "" {
    sessionDataMap[userId] = sData
  }
  return sData, nil
}
