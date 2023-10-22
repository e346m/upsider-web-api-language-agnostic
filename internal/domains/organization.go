package domains

import "time"

type Organization struct {
	ID              string
	Name            string
	Rrepresentative string
	PhoneNumber     string
	Address         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
