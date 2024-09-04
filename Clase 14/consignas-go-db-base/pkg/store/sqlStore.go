package store

import (
	"database/sql"
	"fmt"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

// Read devuelve un producto por su id
func (s *sqlStore) Read(id int) (domain.Product, error) {
	return domain.Product{}, nil
}

// Create agrega un nuevo producto
func (s *sqlStore) Create(product domain.Product) error {
	query := "INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	fmt.Println("stmt", stmt)
	if err != nil {
		return err
	}
	fmt.Println("stmt", stmt)

	res, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price)
	if err != nil {
		return err
	}
	fmt.Println("res", res)

	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("numero de filas", num)

	return nil
}

// Update actualiza un producto
func (s *sqlStore) Update(product domain.Product) error {
	return nil
}

// Delete elimina un producto
func (s *sqlStore) Delete(id int) error {
	return nil
}

// Exists verifica si un producto existe
func (s *sqlStore) Exists(codeValue string) bool {
	return true
}
