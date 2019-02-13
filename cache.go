package main

import (
  // _ "github.com/mattn/go-sqlite3"
  // "github.com/satori/go.uuid"
  // "log"
  "time"
)

// uuid session -> user_id
var session = map[string]string{}

// get user from uuid session
func GetUserId(uuid string) (string, error) {
  userId := session[uuid]
  if userId != "" {
    return userId, nil
  } else {
    err = db.QueryRow("select user_id from session where uuid = ?", uuid).Scan(&userId)
    if err != nil {
      return "", nil
    }
    if uuid != "" {
      session[uuid] = userId
      return userId, nil
    }
  }
  return "", nil
}

func SetSession(uuid string, userId string) error {
  stmt, err := db.Prepare(`INSERT INTO session(uuid, user_id, created) VALUES( ?, ?, ?)`)
  if err != nil {
    return err
  }
  defer stmt.Close()
  _, err = stmt.Exec(uuid, userId, time.Now().String())
  if err != nil {
    return err
  }
  session[uuid] = userId
  return nil
}
