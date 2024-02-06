package model

import "gorm.io/gorm"

//customer is structure
type Customer struct {
	gorm.Model
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Dateofbirth  string `json:"dateofbirth"`
	Mobilenumber string `json:"mobilenumber"`
}
