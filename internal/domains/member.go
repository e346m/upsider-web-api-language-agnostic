package domains

import (
	"time"
)

type Member struct {
	ID        string
	FullName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	*Organization
}
