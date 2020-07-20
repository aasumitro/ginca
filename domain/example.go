package domain

import "time"

type Example struct {
	ID        	uint 		`json:"id" gorm:"primary_key" sql:"index"`
	Name		string		`json:"name"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	DeletedAt 	*time.Time 	`json:"deleted_at" sql:"index"`
}

type ExampleService interface {
	Fetch() ([]Example, error)
	Find(id int) (Example, error)
	//Store()
	//Update()
	//Delete()
}

type ExampleRepository interface {
	Fetch() (data []Example, error error)
	Find(id int) (Example, error)
	//Store()
	//Update()
	//Delete()
}