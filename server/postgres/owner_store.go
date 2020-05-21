package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewOwnerStore(db *sqlx.DB) *OwnerStore {
	return &OwnerStore{
		DB: db,
	}
}

type OwnerStore struct {
	*sqlx.DB
}

func (s *OwnerStore) Owner(id uuid.UUID) (main.Owner, error) {
	var o main.Owner
	if err := s.Get(&o, `SELECT * FROM owners WHERE id = $1`, id); err != nil {
		return main.Owner{}, fmt.Errorf("error getting owner: %w", err)
	}
	return o, nil
}

func (s *OwnerStore) Owners() ([]main.Owner, error) {
	var oo []main.Owner
	if err := s.Select(&oo, `SELECT * FROM owners`); err != nil {
		return []main.Owner{}, fmt.Errorf("error getting owners: %w", err)
	}
	return oo, nil
}

func (s *OwnerStore) CreateOwner(o *main.Owner) error {
	err := s.Get(o, `INSERT INTO owners VALUES ($1, $2) RETURNING *`,
	o.ID
	o.name); err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}
	return nil
}

func (s *OwnerStore) UpateOwner(o *main.Owner) error {
	err := s.Get(o, `UPDATE INTO owners SET name = $2 WHERE id = $1 RETURNING *`,
	o.ID
	o.name); err != nil {
		return fmt.Errorf("error updating thread: %w", err)
	}
	return nil
}

func (s *OwnerStore) DeleteOwner(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM threads WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting thread: %w", err)
	}
	return nil
}
