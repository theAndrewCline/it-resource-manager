package postgres

import (
	"fmt"

	"github.com/akyoto/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/theAndrewCline/it-resource-manager/types"
)

// ComputerStore struct for computers table
type ComputerStore struct {
	*sqlx.DB
}

// Computer gets a computer by ID
func (s *ComputerStore) Computer(id uuid.UUID) (types.Computer, error) {
	var c types.Computer
	err := s.Get(&c, `SELECT * FROM computers WHERE id = $1`, id)
	if err != nil {
		return types.Computer{}, fmt.Errorf("error getting computer: %w", err)
	}
	return c, nil
}

// Computers gets all computers
func (s *ComputerStore) Computers() ([]types.Computer, error) {
	var cc []types.Computer
	err := s.Select(&cc, `SELECT * FROM computers`)
	if err != nil {
		return nil, fmt.Errorf("error getting computers: %w", err)
	}
	return cc, nil
}

// CreateComputer creates a computer with given struct
func (s *ComputerStore) CreateComputer(c *types.Computer) error {
	err := s.Get(c, `INSERT INTO computers VALUES ($1, $2, $3) RETURNING *`,
		c.ID,
		c.OwnerID,
		c.Description)
	if err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}
	return nil
}

// UpdateComputer updates a computer with given struct
func (s *ComputerStore) UpdateComputer(c *types.Computer) error {
	err := s.Get(c, `UPDATE INTO computers SET owner_id = $2 description = $3 WHERE id = $1 RETURNING *`,
		c.ID,
		c.OwnerID,
		c.Description)
	if err != nil {
		return fmt.Errorf("error updating computer: %w", err)
	}
	return nil
}

// DeleteComputer deletes a computer with given ID
func (s *ComputerStore) DeleteComputer(id uuid.UUID) error {
	_, err := s.Exec(`DELETE FROM computers WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting computer: %w", err)
	}
	return nil
}
