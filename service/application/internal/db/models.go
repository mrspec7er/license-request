package db

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Form struct {
	BaseModel
	Name     string     `json:"name" gorm:"type:varchar(63); index:priority:1"`
	Category string     `json:"category" gorm:"type:varchar(63)"`
	Sections []*Section `json:"sections"`
}

type Section struct {
	BaseModel
	FormID int64    `json:"formId" gorm:"type:bigint"`
	Name   string   `json:"name" gorm:"type:varchar(63)"`
	Form   *Form    `json:"form"`
	Fields []*Field `json:"fields"`
}

type Field struct {
	BaseModel
	SectionID  int64     `json:"sectionId" gorm:"type:bigint"`
	Label      string    `json:"label" gorm:"type:varchar(255)"`
	Type       string    `json:"type" gorm:"type:varchar(63)"`
	FieldOrder int32     `json:"order" gorm:"type:int"`
	Section    *Section  `json:"section"`
	Responses  *Response `json:"responses"`
}

type User struct {
	BaseModel
	UID          string         `json:"uid" gorm:"type:varchar(255)"`
	Name         string         `json:"name" gorm:"type:varchar(63)"`
	Email        string         `json:"email" gorm:"type:varchar(63)"`
	Password     string         `json:"password" gorm:"type:text"`
	Role         string         `json:"role" gorm:"type:varchar(63)"`
	Applications []*Application `json:"applications"`
}

type Application struct {
	Number    string      `json:"number" gorm:"primaryKey; index:priority:1; type:varchar(63)"`
	FormID    int64       `json:"formId" gorm:"type:bigint"`
	UserID    int64       `json:"userId" gorm:"type:bigint"`
	Form      *Form       `json:"form"`
	User      *User       `json:"user"`
	Responses []*Response `json:"responses" gorm:"foreignKey:ApplicationNumber;references:Number"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Response struct {
	BaseModel
	ApplicationNumber string       `json:"applicationNumber" gorm:"type:varchar(63)"`
	FieldID           int64        `json:"field_id" gorm:"type:bigint"`
	Value             string       `json:"value" gorm:"type:text"`
	Application       *Application `json:"application"`
	Field             *Field       `json:"field"`
}
