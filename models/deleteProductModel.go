package models

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	// Driver de SQL Server
)

// DeleteProduct elimina un producto de la base de datos
func DeleteProduct(db *sqlx.DB, id int) error {
	query := "DELETE FROM Products WHERE id = @id"
	_, err := db.Exec(query, sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("error al eliminar el producto: %w", err)
	}
	return nil
}
