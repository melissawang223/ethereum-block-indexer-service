package models

import "time"

// BlockModel is raw-data-form exactly as db data
type BlockModel struct {
	ID                int64     `gorm:"column:id; primary_key"`
	BlockNumber       uint64    `gorm:"column:block_number"`
	ParentHash        string    `gorm:"column:parent_hash"`
	Difficulty        uint64    `gorm:"column:difficulty"`
	Hash              string    `gorm:"column:hash"`
	TransactionsCount int       `gorm:"column:transactions_count"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
	Timestamp         uint64    `gorm:"column:timestamp"`

	/*ParentHash:  header.ParentHash.String(),
	UncleHash:   header.UncleHash.String(),
	Coinbase:    header.Coinbase.String(),
	Root:        header.Root.String(),
	TxHash:      header.TxHash.String(),
	ReceiptHash: header.ReceiptHash.String(),
	Bloom:       header.Bloom.Bytes(),
	Difficulty:  header.Difficulty.Int64(),
	Number:      header.Number.Int64(),
	GasLimit:    header.GasLimit,
	GasUsed:     header.GasUsed,
	Time:        header.Time.Uint64(),
	ExtraData:   header.Extra,
	MixDigest:   header.MixDigest.String(),
	Nonce:       header.Nonce.Uint64(),*/
}

// Block data structure
type Block struct {
	BlockNumber       int64  `json:"blockNumber"`
	Timestamp         uint64 `json:"timestamp"`
	Hash              string `json:"hash"`
	ParentHash        string `json:"parentHash"`
	Difficulty        uint64 `json:"column:difficulty"`
	TransactionsCount int    `json:"column:transactionsCount"`
}

type BlockWithTx struct {
	BlockNumber       uint64        `json:"blockNumber"`
	Timestamp         uint64        `json:"timestamp"`
	Hash              string        `json:"hash"`
	ParentHash        string        `json:"parentHash"`
	Difficulty        uint64        `json:"column:difficulty"`
	TransactionsCount int           `json:"column:transactionsCount"`
	Transactions      []Transaction `json:"transactions"`
}
