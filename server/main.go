package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Owner struct {
	ID   uuid.UUID `db:"id"`
	name string    `db:"name"`
}

type Computer struct {
	ID          uuid.UUID `db:"id"`
	OwnerID     uuid.UUID `db:"owner_id"`
	Description string    `db:"description"`
}

type Part struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	ModelNumber string    `db:"model number"`
}

type OwnerStore interface {
	Owner(id uuid.UUID) (Owner, error)
	Owners() ([]Owner, error)
	CreateOwner(c *Owner) error
	UpdateOwner(c *Owner) error
	DeleteOwner(id uuid.UUID) error
}

type ComputerStore interface {
	Computer(id uuid.UUID) (Computer, error)
	Computers() ([]Computer, error)
	CreateComputer(c *Computer) error
	UpdateComputer(c *Computer) error
	DeleteComputer(id uuid.UUID) error
}

type PartStore interface {
	Part(id uuid.UUID) (Part, error)
	Parts() ([]Part, error)
	CreatePart(c *Part) error
	UpdatePart(c *Part) error
	DeletePart(id uuid.UUID) error
}

func main() {
	fmt.Println("Hello, It Manager")
}
