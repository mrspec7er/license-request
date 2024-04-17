package main

import (
	"github.com/joho/godotenv"
	"github.com/mrspec7er/license-request/server/internal/db"
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
	DB.Create(&[]*db.Form{
		{
			Model:    gorm.Model{ID: 1},
			Name:     "Driving License",
			Category: "General",
			Sections: []*db.Section{
				{
					Model: gorm.Model{ID: 1},
					Name:  "Biodata",
					Fields: []*db.Field{
						{
							Model:      gorm.Model{ID: 1},
							Label:      "Name",
							Type:       "String",
							FieldOrder: 1,
						},
						{
							Model:      gorm.Model{ID: 2},
							Label:      "Birth Date",
							Type:       "Date",
							FieldOrder: 2,
						},
						{
							Model:      gorm.Model{ID: 3},
							Label:      "Gender",
							Type:       "Option",
							FieldOrder: 3,
						},
					},
				},
				{
					Model: gorm.Model{ID: 2},
					Name:  "General Knowledge",
					Fields: []*db.Field{
						{
							Model:      gorm.Model{ID: 4},
							Label:      "Technical Skill (1-10)",
							Type:       "Number",
							FieldOrder: 1,
						},
						{
							Model:      gorm.Model{ID: 5},
							Label:      "Practical Experience (1-10)",
							Type:       "Number",
							FieldOrder: 2,
						},
						{
							Model:      gorm.Model{ID: 6},
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
