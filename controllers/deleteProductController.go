package controllers

import (
	"delete-products/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// DeleteProductHandler maneja la eliminación de productos
func DeleteProductHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del producto desde la URL
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		// Eliminar el producto de la base de datos
		err = models.DeleteProduct(db, id)
		if err != nil {
			http.Error(w, "Error al eliminar el producto", http.StatusInternalServerError)
			return
		}

		// Responder con éxito
		w.WriteHeader(http.StatusNoContent)
	}
}
