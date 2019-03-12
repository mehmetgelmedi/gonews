package models

type User struct {
    ID int64
    Username string
    Password string
    IsAdmin bool
}

type UserLogin struct {
    Username string
    Password string
}