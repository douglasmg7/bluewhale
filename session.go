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

type session_data struct {
  name string
}

// [session]user_id
var userIdMap = map[string]int{}

// [user_id]session_data
var sessionDataMap = map[int]session_data{}

func data(int) {

}

// get user from uuid session
func UserId(req *http.Request) (int, error) {
  cookie, err := req.Cookie("session")
  // no cookie.
  if err == http.ErrNoCookie {
    return 0, nil
    // some error
  } else if err != nil {
    return 0, err
  }
  // some
  if cookie != nil {
    session := cookie.String()
    userId := userIdMap[session]
    if userId != 0 {
      return userId, nil
    } else {
      // get from db
      err = db.QueryRow("select user_id from session where uuid = ?", cookie.String()).Scan(&userId)
      if err != nil {
        return 0, err
      }
      if userId != 0 {
        userIdMap[session] = userId
        return userId, nil
      }
    }
  }
  // no cookie
  return 0, nil
}

func NewSession(w http.ResponseWriter, userId int) error {
  // create cookie
  session, err := uuid.NewV4()
  if err != nil {
    return err
  }
  sessinoString := session.String()
  // save cookie
  http.SetCookie(w, &http.Cookie{
    Name:  "session",
    Value: session.String(),
    // Secure: true, // to use only in https
    HttpOnly: true, // can't be used into js client
  })
  // save cache on db
  stmt, err := db.Prepare(`INSERT INTO session(uuid, user_id, created) VALUES( ?, ?, ?)`)
  if err != nil {
    return err
  }
  defer stmt.Close()
  _, err = stmt.Exec(session, userId, time.Now().String())
  if err != nil {
    return err
  }
  // set session cache
  userIdMap[sessinoString] = userId
  return nil
}
