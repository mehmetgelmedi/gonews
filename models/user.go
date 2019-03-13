package models

type User struct {
    ID int64 
    Username string `json:"username"`
    Password string `json:"password"`
    IsAdmin bool
}

type UserLogin struct {
    Username string
    Password string
}