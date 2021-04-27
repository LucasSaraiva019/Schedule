package models

import "time"

type User struct {
	ID          uint64    `json:"id,omitempty"`
	Name        string    `json:"nome,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phoneNumber, omitempty"`
	Password    string    `json:"password,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}
