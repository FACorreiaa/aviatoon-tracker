package queries

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/create-go-app/net_http-go-template/app/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
	"strings"
)

type TaxQueries struct {
	*sqlx.DB
}

func (q *TaxQueries) CreateTaxTable() error {
	log.Println("Creating country table.")
	if _, err := q.Exec(
		`CREATE TABLE tax (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), tax_id varchar(255), tax_name varchar(255), iata_code varchar(255))`); err != nil {
		return fmt.Errorf("error creating aircraft table: %w", err)
	}
	return nil
}

func (q *CountryQueries) CreateAviationTax(t *models.Tax) error {
	// Start a transaction.

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	taxId, err := StringToInt(t.TaxId)
	if err != nil {
		return fmt.Errorf("error converting tax to int: %w", err)
	}

	if _, err := tx.ExecContext(context.Background(),
		`INSERT INTO tax VALUES ($1, $2, $3, $4, $5, $6)`,
		t.ID,
		taxId,
		t.TaxName,
		t.IataCode,
		t.CreatedAt,
		t.UpdatedAt); err != nil {
		return fmt.Errorf("error inserting values: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (q *CountryQueries) GetAviationTax() ([]models.Tax, error) {
	var tax []models.Tax

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Send query to database.
	rows, err := tx.Query(`SELECT * FROM tax ORDER BY tax_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Tax
		err := rows.Scan(
			&t.ID,
			&t.TaxId,
			&t.TaxName,
			&t.IataCode,
			&t.CreatedAt,
			&t.UpdatedAt)

		if err != nil {
			return nil, err
		}
		tax = append(tax, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return tax, nil
}

func (q *CountryQueries) GetAviationTaxByID(id string) (models.Tax, error) {
	var tax models.Tax

	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return tax, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(context.Background(), `SELECT
			id,
			tax_id,
			tax_name,
			iata_code,
			created_at,
			updated_at
		FROM tax
		WHERE tax_id = $1 LIMIT 1`, id)
	err = row.Scan(
		&tax.ID,
		&tax.TaxId,
		&tax.TaxName,
		&tax.IataCode,
		&tax.CreatedAt,
		&tax.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tax, fmt.Errorf("tax with ID %s not found: %w", id, err)
		}
		return tax, fmt.Errorf("failed to scan tax: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return tax, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return tax, nil
}

func (q *CountryQueries) DeleteAviationTaxByID(id string) error {
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(context.Background(), "DELETE FROM tax WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete tax: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *CountryQueries) UpdateAviationTaxByID(id string, updates map[string]interface{}) error {
	tx, err := q.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	var setColumns []string
	var args []interface{}

	for key, value := range updates {
		setColumns = append(setColumns, fmt.Sprintf("%s = $%d", key, len(args)+1))
		args = append(args, value)
	}
	args = append(args, id)

	stmt := fmt.Sprintf("UPDATE tax SET %s WHERE id = $%d", strings.Join(setColumns, ", "), len(args))
	_, err = tx.ExecContext(context.Background(), stmt, args...)
	if err != nil {
		return fmt.Errorf("failed to update tax: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q *CountryQueries) GetNumberOfAviationTax() (int, error) {
	tx, err := q.BeginTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var count int
	err = tx.QueryRowContext(context.TODO(), "SELECT COUNT(*) FROM tax").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no tax found")
		}
		return 0, fmt.Errorf("failed to get number of tax: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return count, nil
}
