package main

import "github.com/jinzhu/gorm"

type Wallet struct {
	gorm.Model
}

// User
type User struct {
	gorm.Model
	Email         string
	Password      string `json:",omitempty"`
	AccountType   int
	FirstName     string
	LastName      string
	PhoneNumber   string
	RecoveryEmail string
}

