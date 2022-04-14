package transactiondao

import (
	"errors"
	"fmt"
	"github.com/ethereum-block-indexer-service/models"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const table = "transaction_record"

// queryOption set query condition, used by queryChain()
type queryOption struct {
	BlockNum string
	txHash   string
	OrderBy  []string // ex: {"transaction_record.id desc", "transaction_record.updated_at asc"}
	Limit    int
	Offset   int
}

// NewQueryOption generate a queryOption with given default value
func NewQueryOption() queryOption {
	query := queryOption{}
	return query
}

// New a row
func New(tx *gorm.DB, bankModel *models.Transaction) error {
	err := tx.Table(table).
		Create(bankModel).Error

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 { // new duplicate
		fmt.Println("1062", err)
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

// Get return a record as raw-data-form
func Get(tx *gorm.DB, query *queryOption) *models.Transaction {
	result := models.Transaction{}
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scopes(limitScope(query.Limit, query.Offset)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}

	return &result
}

// Gets return records as raw-data-form
func Gets(tx *gorm.DB, query *queryOption) *[]models.Transaction {
	result := []models.Transaction{}
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scopes(limitScope(query.Limit, query.Offset)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}

	return &result
}

func queryChain(query *queryOption) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(blockNumberEqualScope(query.BlockNum)).
			Scopes(txHashEqualScope(query.txHash)).
			Scopes(orderByScope(query.OrderBy))
	}
}

func txHashEqualScope(txHash string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if txHash != "" {
			return db.Where(table+".tx_hash = ?", txHash)
		}
		return db
	}
}

func blockNumberEqualScope(blockNumber string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if blockNumber != "" {
			return db.Where(table+".block_number = ?", blockNumber)
		}
		return db
	}
}

func orderByScope(orderBy []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(orderBy) != 0 {
			order := orderBy[0]
			for _, o := range orderBy[1:] {
				order = order + ", " + o
			}
			return db.Order(order)
		}
		return db
	}
}

func limitScope(limit int, offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit > 0 {
			return db.Limit(limit).Offset(offset)
		}
		return db
	}
}
