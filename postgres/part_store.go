package postgres

import (
	"fmt"

	"github.com/akyoto/uuid"
	"github.com/jmoiron/sqlx"
	manager "github.com/theAndrewCline/it-resource-manager"
)

// PartStore Implements manager.PartStore interface
type PartStore struct {
	*sqlx.DB
}

// Part returns Part with the matching ID passed to it
func (s *PartStore) Part(id uuid.UUID) (manager.Part, error) {
	var p manager.Part
	err := s.Get(&p, `SELECT * FROM parts WHERE id = $1`, id)
	if err != nil {
		return manager.Part{}, fmt.Errorf("error getting part: %w", err)
	}
	return p, nil
}

// Parts returns all the parts in the database
func (s *PartStore) Parts() ([]manager.Part, error) {
	var pp []manager.Part
	err := s.Select(&pp, `SELECT * FROM parts`)
	if err != nil {
		return []manager.Part{}, fmt.Errorf("error getting parts: %w", err)
	}
	return pp, nil
}

// CreatePart From struct passed to it
func (s *PartStore) CreatePart(p *manager.Part) error {
	err := s.Get(p, `INSERT INTO parts VALUES ($1, $2, $3, $4) RETURNING *`,
		p.ID,
		p.ComputerID,
		p.Name,
		p.ModelNumber)
	if err != nil {
		return fmt.Errorf("error creating part: %w", err)
	}
	return nil
}

// UpdatePart updates Part based on struct passed to it
func (s *PartStore) UpdatePart(p *manager.Part) error {
	sql := `UPDATE 
				INTO parts 
				SET 
					computer_id = $2
					name = $3
					model_number = $4
				WHERE id = $1
				RETURNING *`

	err := s.Get(p, sql,
		p.ID,
		p.ComputerID,
		p.Name,
		p.ModelNumber)
	if err != nil {
		return fmt.Errorf("error updating part: %w", err)
	}
	return nil
}

// DeletePart deletes part
func (s *PartStore) DeletePart(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM part WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting part: %w", err)
	}
	return nil
}
