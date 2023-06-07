package models


/*
Object created for the db
*/
type Url struct {
    Id              int    `json:"id"` 
    Url             string `json:"url"`
    Hashed          string `json:"hashed"`
    Message         string `json:"message"`
    ShowMsg         bool   `json:"show_msg"`
}
