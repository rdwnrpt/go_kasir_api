package repositories

import (
	"database/sql"
	"fmt"
	"go_kasir_api/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range items {
		var productName string
		var productPrice, stock int
		
		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id=$1", item.ProductID).
			Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		// Check stock availability
		if stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s", productName)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		// Update product stock
		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2",
			item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	// Insert transaction
	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions(total_amount) VALUES($1) RETURNING id", totalAmount).
		Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	// Insert transaction details
	for i := range details {
		details[i].TransactionID = transactionID
		_, err = tx.Exec(
			"INSERT INTO transaction_details(transaction_id, product_id, quantity, subtotal) VALUES($1, $2, $3, $4)",
			transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}