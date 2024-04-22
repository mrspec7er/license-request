package main

import (
	"github.com/joho/godotenv"
	"github.com/mrspec7er/license-request/services/form/internal/db"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	DB := db.StartConnection()

	SeedDatabase(DB)

}

func SeedDatabase(DB *gorm.DB) {
	DB.Create([]*db.Form{
		{
			BaseModel: db.BaseModel{ID: 1},
			Name:      "Driving License",
			Category:  "General",
			Sections: []*db.Section{
				{
					BaseModel: db.BaseModel{ID: 1},
					Name:      "Biodata",
					Fields: []*db.Field{
						{
							BaseModel:  db.BaseModel{ID: 1},
							Label:      "Name",
							Type:       "String",
							FieldOrder: 1,
						},
						{
							BaseModel:  db.BaseModel{ID: 2},
							Label:      "Birth Date",
							Type:       "Date",
							FieldOrder: 2,
						},
						{
							BaseModel:  db.BaseModel{ID: 3},
							Label:      "Gender",
							Type:       "Option",
							FieldOrder: 3,
						},
					},
				},
				{
					BaseModel: db.BaseModel{ID: 2},
					Name:      "General Knowledge",
					Fields: []*db.Field{
						{
							BaseModel:  db.BaseModel{ID: 4},
							Label:      "Technical Skill (1-10)",
							Type:       "Number",
							FieldOrder: 1,
						},
						{
							BaseModel:  db.BaseModel{ID: 5},
							Label:      "Practical Experience (1-10)",
							Type:       "Number",
							FieldOrder: 2,
						},
						{
							BaseModel:  db.BaseModel{ID: 6},
							Label:      "Note",
							Type:       "Text",
							FieldOrder: 3,
						},
					},
				},
			},
		},
	})

	DB.Create([]*db.User{
		{
			BaseModel: db.BaseModel{ID: 1},
			UID:       "ADMIN-1",
			Name:      "User Admin",
			Email:     "admin@email.com",
			Password:  "",
			Role:      "ADMIN",
		},
		{
			BaseModel: db.BaseModel{ID: 2},
			UID:       "User-1",
			Name:      "Basic User",
			Email:     "user@email.com",
			Password:  "",
			Role:      "USER",
		},
	})

	DB.Create([]*db.Application{
		{
			Number: "DRV-001",
			UserID: 1,
			FormID: 1,
			Responses: []*db.Response{
				{
					BaseModel: db.BaseModel{ID: 1},
					FieldID:   1,
					Value:     "Basic User",
				},
				{
					BaseModel: db.BaseModel{ID: 2},
					FieldID:   2,
					Value:     "2001-28-03",
				},
				{
					BaseModel: db.BaseModel{ID: 3},
					FieldID:   3,
					Value:     "Female",
				},
				{
					BaseModel: db.BaseModel{ID: 4},
					FieldID:   4,
					Value:     "8",
				},
				{
					BaseModel: db.BaseModel{ID: 5},
					FieldID:   5,
					Value:     "7",
				},
				{
					BaseModel: db.BaseModel{ID: 6},
					FieldID:   6,
					Value:     "Little bit too aggressive in corner but has very good handling and fast response",
				},
			},
		},
		{
			Number: "DRV-002",
			UserID: 1,
			FormID: 1,
			Responses: []*db.Response{
				{
					BaseModel: db.BaseModel{ID: 7},
					FieldID:   1,
					Value:     "Basic User",
				},
				{
					BaseModel: db.BaseModel{ID: 8},
					FieldID:   3,
					Value:     "Female",
				},
				{
					BaseModel: db.BaseModel{ID: 9},
					FieldID:   4,
					Value:     "8",
				},
			},
		},
	})
}
