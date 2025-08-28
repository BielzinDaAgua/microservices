package db

import (
	"database/sql"
	"fmt"
)

type InventoryRepository struct {
	DB *sql.DB
}

func NewInventoryRepository(db *sql.DB) *InventoryRepository {
	return &InventoryRepository{DB: db}
}

func (r *InventoryRepository) ItemExists(itemID string) (bool, error) {
	var exists bool
	err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM inventory WHERE item_id = $1)", itemID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking item existence: %w", err)
	}
	return exists, nil
}
