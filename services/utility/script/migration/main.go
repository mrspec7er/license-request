package main

import (
	"github.com/joho/godotenv"
	"github.com/mrspec7er/license-request/service/utility/dto"
	"github.com/mrspec7er/license-request/service/utility/internal/db"

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
	var formCount int64
	DB.Model(&dto.Form{}).Count(&formCount)
	if formCount == 0 {
		DB.Create([]*dto.Form{
			{
				BaseModel: dto.BaseModel{ID: 1},
				Name:      "Driving License",
				Category:  "General",
				Sections: []*dto.Section{
					{
						BaseModel: dto.BaseModel{ID: 1},
						Name:      "Biodata",
						Fields: []*dto.Field{
							{
								BaseModel:  dto.BaseModel{ID: 1},
								Label:      "Name",
								Type:       "String",
								FieldOrder: 1,
							},
							{
								BaseModel:  dto.BaseModel{ID: 2},
								Label:      "Birth Date",
								Type:       "Date",
								FieldOrder: 2,
							},
							{
								BaseModel:  dto.BaseModel{ID: 3},
								Label:      "Gender",
								Type:       "Option",
								FieldOrder: 3,
							},
						},
					},
					{
						BaseModel: dto.BaseModel{ID: 2},
						Name:      "General Knowledge",
						Fields: []*dto.Field{
							{
								BaseModel:  dto.BaseModel{ID: 4},
								Label:      "Technical Skill (1-10)",
								Type:       "Number",
								FieldOrder: 1,
							},
							{
								BaseModel:  dto.BaseModel{ID: 5},
								Label:      "Practical Experience (1-10)",
								Type:       "Number",
								FieldOrder: 2,
							},
							{
								BaseModel:  dto.BaseModel{ID: 6},
								Label:      "Note",
								Type:       "Text",
								FieldOrder: 3,
							},
						},
					},
				},
			},
		})
	}

	var userCount int64
	DB.Model(&dto.User{}).Count(&userCount)
	if userCount == 0 {
		DB.Create([]*dto.User{
			{
				ID:            "ADMIN-001",
				Email:         "admin@email.com",
				Picture:       "https://lh3.googleusercontent.com/a-/ALV-UjUiYmyOrvOAaQsnGfnVDs6QrtTn7sLD5ICee2OcVz32pAi6Pwj1=s96-c",
				VerifiedEmail: true,
				Password:      "",
				Role:          "ADMIN",
			},
			{
				ID:            "USER-001",
				Email:         "user@email.com",
				Picture:       "https://lh3.googleusercontent.com/a-/ALV-UjUiYmyOrvOAaQsnGfnVDs6QrtTn7sLD5ICee2OcVz32pAi6Pwj1=s96-c",
				VerifiedEmail: true,
				Password:      "",
				Role:          "USER",
			},
		})
	}

	var applicationCount int64
	DB.Model(&dto.Application{}).Count(&applicationCount)
	if applicationCount == 0 {
		DB.Create([]*dto.Application{
			{
				Number: "DRV-001",
				UserID: "USER-001",
				FormID: 1,
				Responses: []*dto.Response{
					{
						BaseModel: dto.BaseModel{ID: 1},
						FieldID:   1,
						Value:     "Basic User",
					},
					{
						BaseModel: dto.BaseModel{ID: 2},
						FieldID:   2,
						Value:     "2001-28-03",
					},
					{
						BaseModel: dto.BaseModel{ID: 3},
						FieldID:   3,
						Value:     "Female",
					},
					{
						BaseModel: dto.BaseModel{ID: 4},
						FieldID:   4,
						Value:     "8",
					},
					{
						BaseModel: dto.BaseModel{ID: 5},
						FieldID:   5,
						Value:     "7",
					},
					{
						BaseModel: dto.BaseModel{ID: 6},
						FieldID:   6,
						Value:     "Little bit too aggressive in corner but has very good handling and fast response",
					},
				},
			},
			{
				Number: "DRV-002",
				UserID: "USER-001",
				FormID: 1,
				Responses: []*dto.Response{
					{
						BaseModel: dto.BaseModel{ID: 7},
						FieldID:   1,
						Value:     "Basic User",
					},
					{
						BaseModel: dto.BaseModel{ID: 8},
						FieldID:   3,
						Value:     "Female",
					},
					{
						BaseModel: dto.BaseModel{ID: 9},
						FieldID:   4,
						Value:     "8",
					},
				},
			},
		})
	}

}
