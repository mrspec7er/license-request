package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/mrspec7er/license-request/server/internal/db"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	DB := db.StartConnection()

	PushSchema(DB)
	RunSeeders(DB)

}

func PushSchema(DB *sqlx.DB) {
	DB.MustExec(db.FormSchema)

	DB.MustExec(db.SectionSchema)

	DB.MustExec(db.FieldSchema)

	DB.MustExec(db.AccountSchema)

	DB.MustExec(db.ApplicationSchema)

	DB.MustExec(db.ResponseSchema)
}

func RunSeeders(DB *sqlx.DB) {
	txInsertNewForm := DB.MustBegin()
	txInsertNewForm.MustExec("INSERT INTO Form (id, name, category) VALUES ($1, $2, $3)", 1, "Driving License", "General")
	txInsertNewForm.MustExec("INSERT INTO Form (id, name, category) VALUES ($1, $2, $3)", 2, "Medical License", "Profession")
	txInsertNewForm.Commit()

	txInsertNewSection := DB.MustBegin()
	txInsertNewSection.MustExec("INSERT INTO Section (id, form_id, name) VALUES ($1, $2, $3)", 1, 1, "Biodata")
	txInsertNewSection.MustExec("INSERT INTO Section (id, form_id, name) VALUES ($1, $2, $3)", 2, 1, "Knowledge")
	txInsertNewSection.MustExec("INSERT INTO Section (id, form_id, name) VALUES ($1, $2, $3)", 3, 2, "Practical Experience")
	txInsertNewSection.Commit()

	txInsertNewField := DB.MustBegin()
	txInsertNewField.MustExec("INSERT INTO Field (id, section_id, label, type, field_order) VALUES ($1, $2, $3, $4, $5)", 1, 1, "Name", "String", 1)
	txInsertNewField.MustExec("INSERT INTO Field (id, section_id, label, type, field_order) VALUES ($1, $2, $3, $4, $5)", 2, 1, "Date of Birth", "Date", 2)
	txInsertNewField.MustExec("INSERT INTO Field (id, section_id, label, type, field_order) VALUES ($1, $2, $3, $4, $5)", 3, 1, "Gender", "Option", 3)
	txInsertNewField.MustExec("INSERT INTO Field (id, section_id, label, type, field_order) VALUES ($1, $2, $3, $4, $5)", 4, 2, "Technical Knowledge (1-10)", "Number", 1)
	txInsertNewField.MustExec("INSERT INTO Field (id, section_id, label, type, field_order) VALUES ($1, $2, $3, $4, $5)", 5, 2, "Practical Skill (1-10)", "Number", 2)
	txInsertNewField.MustExec("INSERT INTO Field (id, section_id, label, type, field_order) VALUES ($1, $2, $3, $4, $5)", 6, 2, "Note", "Text", 3)
	txInsertNewField.MustExec("INSERT INTO Field (id, section_id, label, type, field_order) VALUES ($1, $2, $3, $4, $5)", 7, 3, "Description", "Text", 1)
	txInsertNewField.Commit()

	txInsertNewAccount := DB.MustBegin()
	txInsertNewAccount.MustExec("INSERT INTO Account (id, uid, name, email, password, role) VALUES ($1, $2, $3, $4, $5, $6)", 1, "unknown-123", "Admin", "admin@email.com", "mrc201", "ADMIN")
	txInsertNewAccount.MustExec("INSERT INTO Account (id, uid, name, email, password, role) VALUES ($1, $2, $3, $4, $5, $6)", 2, "unknown-3.14", "User", "user@email.com", "mrc201", "USER")
	txInsertNewAccount.Commit()

	txInsertNewApplication := DB.MustBegin()
	txInsertNewApplication.MustExec("INSERT INTO Application (number, form_id, account_id) VALUES ($1, $2, $3)", "GNR-001", 1, 2)
	txInsertNewApplication.MustExec("INSERT INTO Application (number, form_id, account_id) VALUES ($1, $2, $3)", "PRF-001", 2, 2)
	txInsertNewApplication.Commit()

	txInsertNewResponse := DB.MustBegin()
	txInsertNewResponse.MustExec("INSERT INTO Response (id, application_number, field_id, value) VALUES ($1, $2, $3, $4)", 1, "GNR-001", 1, "First User")
	txInsertNewResponse.MustExec("INSERT INTO Response (id, application_number, field_id, value) VALUES ($1, $2, $3, $4)", 2, "GNR-001", 2, "18-05-2006")
	txInsertNewResponse.MustExec("INSERT INTO Response (id, application_number, field_id, value) VALUES ($1, $2, $3, $4)", 3, "GNR-001", 3, "Female")
	txInsertNewResponse.MustExec("INSERT INTO Response (id, application_number, field_id, value) VALUES ($1, $2, $3, $4)", 4, "GNR-001", 4, "9")
	txInsertNewResponse.MustExec("INSERT INTO Response (id, application_number, field_id, value) VALUES ($1, $2, $3, $4)", 5, "GNR-001", 5, "8")
	txInsertNewResponse.MustExec("INSERT INTO Response (id, application_number, field_id, value) VALUES ($1, $2, $3, $4)", 6, "GNR-001", 6, "Very good at at corner and straight line, need to improve on start and endurance")
	txInsertNewResponse.MustExec("INSERT INTO Response (id, application_number, field_id, value) VALUES ($1, $2, $3, $4)", 7, "PRF-001", 7, "Already practice for 2 year at second class hospital and being professor assistant for one year")
	txInsertNewResponse.Commit()
}
