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
    RefreshToken      *string `gorm:"type:text"`
    AccessToken       *string `gorm:"type:text"`
    ExpiresAt         *int64
    IDToken           *string `gorm:"type:text"`
    Scope             *string `gorm:"type:text"`
    SessionState      *string `gorm:"type:text"`
    TokenType         *string `gorm:"type:text"`
}

// Session represents the sessions table
type Session struct {
    ID           int64     `gorm:"primaryKey;autoIncrement"`
    UserID       int64     `gorm:"not null;column:userId"`
    Expires      time.Time `gorm:"not null"`
    SessionToken string    `gorm:"size:255;not null;column:sessionToken"`
}

// User represents the users table
type User struct {
    ID             int64      `gorm:"primaryKey;autoIncrement"`
    Name           *string    `gorm:"size:255"`
    Email          *string    `gorm:"size:255"`
    EmailVerified  *time.Time `gorm:"column:emailVerified"`
    Image          *string    `gorm:"type:text"`
}




