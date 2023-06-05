package models

type Link struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  URL  string `json:"url"`
  HashedUrl  string `json:"hashed_url"`
}
