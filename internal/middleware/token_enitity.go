package middleware

import "time"

type TokenEnitty struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  *string   `json:"lastName"`
	Phone     string    `json:"phone"`
	Email     *string   `json:"email"`
	Time      time.Time `json:"time"`
}
