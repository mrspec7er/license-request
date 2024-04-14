package dto

type Form struct {
	ID       int64      `json:"id" db:"id"`
	Name     string     `json:"name" db:"name"`
	Category string     `json:"category" db:"category"`
	Sections []*Section `json:"sections"`
}

type Section struct {
	ID     int64    `json:"id" db:"id"`
	FormID int64    `json:"form_id" db:"form_id"`
	Name   string   `json:"name" db:"name"`
	Form   *Form    `json:"form"`
	Fields []*Field `json:"fields"`
}

type Field struct {
	ID         int64       `json:"id" db:"id"`
	SectionID  int64       `json:"section_id" db:"section_id"`
	Label      string      `json:"label" db:"label"`
	Type       string      `json:"type" db:"type"`
	FieldOrder int32       `json:"order" db:"field_order"`
	Section    *Section    `json:"section"`
	Responses  []*Response `json:"responses"`
}

type Account struct {
	ID           int64          `json:"id" db:"id"`
	UID          string         `json:"uid" db:"uid"`
	Name         string         `json:"name" db:"name"`
	Email        string         `json:"email" db:"email"`
	Password     string         `json:"password" db:"password"`
	Role         string         `json:"role" db:"role"`
	Applications []*Application `json:"applications"`
}

type Application struct {
	Number    string      `json:"number" db:"number"`
	FormID    int64       `json:"form_id" db:"form_id"`
	AccountID int64       `json:"account_id" db:"account_id"`
	Form      *Form       `json:"form"`
	Account   *Account    `json:"account"`
	Responses []*Response `json:"responses"`
}

type Response struct {
	ID                int64        `json:"id" db:"id"`
	ApplicationNumber string       `json:"application_number" db:"application_number"`
	FieldID           string       `json:"field_id" db:"field_id"`
	Value             string       `json:"value" db:"value"`
	Application       *Application `json:"application"`
	Field             *Field       `json:"field"`
}
