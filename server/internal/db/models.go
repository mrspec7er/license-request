package db

import (
	"time"

	"gorm.io/gorm"
)

type Form struct {
	gorm.Model
	Name     string     `json:"name" gorm:"type:varchar(63); index:priority:1"`
	Category string     `json:"category" gorm:"type:varchar(63)"`
	Sections []*Section `json:"sections"`
}

type Section struct {
	gorm.Model
	FormID int64    `json:"form_id" gorm:"type:bigint"`
	Name   string   `json:"name" gorm:"type:varchar(63)"`
	Form   *Form    `json:"form"`
	Fields []*Field `json:"fields"`
}

type Field struct {
	gorm.Model
	SectionID  int64       `json:"section_id" gorm:"type:bigint"`
	Label      string      `json:"label" gorm:"type:varchar(255)"`
	Type       string      `json:"type" gorm:"type:varchar(63)"`
	FieldOrder int32       `json:"order" gorm:"type:int"`
	Section    *Section    `json:"section"`
	Responses  []*Response `json:"responses"`
}

type Account struct {
	gorm.Model
	UID          string         `json:"uid" gorm:"type:varchar(255)"`
	Name         string         `json:"name" gorm:"type:varchar(63)"`
	Email        string         `json:"email" gorm:"type:varchar(63)"`
	Password     string         `json:"password" gorm:"type:text"`
	Role         string         `json:"role" gorm:"type:varchar(63)"`
	Applications []*Application `json:"applications"`
}

type Application struct {
	Number    string      `json:"number" gorm:"primaryKey; index:priority:1; type:varchar(63)"`
	FormID    int64       `json:"form_id" gorm:"type:bigint"`
	AccountID int64       `json:"account_id" gorm:"type:bigint"`
	Form      *Form       `json:"form"`
	Account   *Account    `json:"account"`
	Responses []*Response `json:"responses" gorm:"foreignKey:ApplicationNumber;references:Number"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Response struct {
	gorm.Model
	ApplicationNumber string       `json:"application_number" gorm:"type:varchar(63)"`
	FieldID           string       `json:"field_id" gorm:"type:bigint"`
	Value             string       `json:"value" gorm:"type:text"`
	Application       *Application `json:"application"`
	Field             *Field       `json:"field"`
}
