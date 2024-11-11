package models

import (
	"time"
)

// VerificationToken represents the verification_token table
type VerificationToken struct {
	Identifier string    `gorm:"primaryKey;not null;type:text"`
	Token      string    `gorm:"primaryKey;not null;type:text"`
	Expires    time.Time `gorm:"not null"`
}

// Account represents the accounts table
type Account struct {
	ID                int64   `gorm:"primaryKey;autoIncrement"` 
	UserID            int64   `gorm:"not null;column:userId"`
	Type              string  `gorm:"size:255;not null"`
	Provider          string  `gorm:"size:255;not null"`
	ProviderAccountID string  `gorm:"size:255;not null;column:providerAccountId"`
	RefreshToken      string `gorm:"type:text"`
	AccessToken       string `gorm:"type:text"`
	ExpiresAt         int64
	IDToken           string `gorm:"type:text"`
	Scope             string `gorm:"type:text"`
	SessionState      string `gorm:"type:text"`
	TokenType         string `gorm:"type:text"`
}

// Session represents the sessions table
type Session struct {
	ID           int64     `gorm:"primaryKey;autoIncrement"`
	UserID       int64     `gorm:"not null;column:userId"`
	Expires      time.Time `gorm:"not null"`
	SessionToken string    `gorm:"size:255;not null;column:sessionToken"`
	User         User
}

// User represents the users table
type User struct {
	ID            int64      `gorm:"primaryKey;autoIncrement"`
	Name          string    `gorm:"size:255,unique" json:"name"` 
	Email         string    `gorm:"size:255" json:"email"` 
	EmailVerified time.Time `gorm:"column:emailVerified" json:"emailVerified"` 
	Image         string    `gorm:"type:text" json:"image"` 
	Website       string    `gorm:"type:text" json:"website"`
	GithubURL     string    `gorm:"type:text" json:"githubURL"`  
	LinkedInURL   string    `gorm:"type:text" json:"linkedInURL"`  
	XUrl          string    `gorm:"type:text" json:"xUrl"` 
	Birthday      string    `gorm:"type:text" json:"birthday"`  
	Location      string    `gorm:"type:text" json:"location"` 
}

type Experience struct {
	ID     int64 `gorm:"primaryKey;autoIncrement"`
	UserID int
	Name   string
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Solution struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	ProblemID int64     `gorm:"not null"`
	Language  string   `gorm:"size:50;not null"`
	Code      string   `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

type Keymaps struct {
	ID     int64 `gorm:"primaryKey;autoIncrement"`
	UserID int
	Choice *string `gorm:"size:50"`
	User   User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Problem struct {
	ID         int64      `gorm:"primaryKey;autoIncrement"`
	Title      string     `gorm:"size:255;not null"`
	Difficulty string     `gorm:"size:50;not null"`
	Tags       []string   `gorm:"type:text[]"`
	Markdown   string     `gorm:"type:text;not null"`
	Solutions  []Solution `gorm:"foreignKey:ProblemID;constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
}
