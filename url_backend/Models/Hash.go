package models

import (
    "time"
)

type User struct {
  ID           uint
  URL         string
  Hash         string
  CreatedAt    time.Time
  UpdatedAt    time.Time
}
