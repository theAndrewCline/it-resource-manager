package postgres

import (
	"fmt"

	"github.com/akyoto/uuid"
	"github.com/jmoiron/sqlx"
	manager "github.com/theAndrewCline/it-resource-manager"
)

// OwnerStore type
type OwnerStore struct {
	*sqlx.DB
}

// Owner gets owner by id
func (s *OwnerStore) Owner(id uuid.UUID) (manager.Owner, error) {
	var o manager.Owner
	err := s.Get(&o, `SELECT * FROM owners WHERE id = $1`, id)
	if err != nil {
		return manager.Owner{}, fmt.Errorf("error getting owner: %w", err)
	}
	return o, nil
}

// Owners gets all owner
func (s *OwnerStore) Owners() ([]manager.Owner, error) {
	var oo []manager.Owner
	if err := s.Select(&oo, `SELECT * FROM owners`); err != nil {
		return []manager.Owner{}, fmt.Errorf("error getting owners: %w", err)
	}
	return oo, nil
}

// CreateOwner creates owner with given struct
func (s *OwnerStore) CreateOwner(o *manager.Owner) error {
	if err := s.Get(o, `INSERT INTO owners VALUES ($1, $2) RETURNING *`,
		o.ID,
		o.Name); err != nil {
		return fmt.Errorf("error creating owner: %w", err)
	}
	return nil
}

// UpdateOwner updates owner with given stuct
func (s *OwnerStore) UpdateOwner(o *manager.Owner) error {
	if err := s.Get(o, `UPDATE INTO owners SET name = $2 WHERE id = $1 RETURNING *`,
		o.ID,
		o.Name); err != nil {
		return fmt.Errorf("error updating owner: %w", err)
	}
	return nil
}

// DeleteOwner deletes owner by id
func (s *OwnerStore) DeleteOwner(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM owners WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting owner: %w", err)
	}
	return nil
}
