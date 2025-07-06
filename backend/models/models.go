package models

type User struct {
    UserID      int    `json:"user_id"`
    Name        string `json:"name"`
    Email       string `json:"email"`
    Password    string `json:"password"`
    TimeBalance int    `json:"time_balance"`
}
