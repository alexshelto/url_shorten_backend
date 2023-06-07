package models


/*
Object created for the db
*/
type Url struct {
    // Id              uint   `json:"id"`
    Id              int    `gorm:"primaryKey;default:auto_random()"`
    Url             string `json:"url"`
    Hashed          string `json:"hashed"`
    Message         string `json:"message"`
    ShowMsg         bool   `json:"show_msg"`
}
