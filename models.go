package main

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"unique"`
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}

type Terminal struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string
	Location  string
	CreatedAt time.Time
}
