package accountdomain

import (
	"strings"
	"time"
)

type Account struct {
	id        string
	name      string
	email     string
	password  string
	status    Status
	createdAt *time.Time
}

func (a *Account) GetID() string {
	return a.id
}

func (a *Account) GetName() string {
	return a.name
}

func (a *Account) GetEmail() string {
	return a.email
}

func (a *Account) GetPassword() string {
	return a.password
}

func (a *Account) GetStatus() Status {
	return a.status
}

func (a *Account) GetCreatedAt() time.Time {
	return *a.createdAt
}

func NewAccount(id, name, email, password string, status Status, createdAt *time.Time) (*Account, error) {
	return &Account{
		id:        id,
		name:      name,
		email:     email,
		password:  password,
		status:    status,
		createdAt: createdAt,
	}, nil
}

type Status int

const (
	StatusActivated Status = iota
	StatusBanned
)

func (r Status) String() string {
	switch r {
	case StatusActivated:
		return "activated"
	case StatusBanned:
		return "banned"
	default:
		return "unknown"
	}
}

func Enum(s string) Status {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "activated":
		return StatusActivated
	default:
		return StatusBanned
	}
}
