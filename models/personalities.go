package models

import "gorm.io/gorm"

type Personality struct {
	gorm.Model
	Id      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}