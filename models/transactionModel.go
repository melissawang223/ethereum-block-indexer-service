package models

// Transaction data structure
type Transaction struct {
	ID          int64  `gorm:"column:id; primary_key"`
	BlockNumber uint64 `gorm:"column:block_number" json:"blockNumber"`
	Hash        string `gorm:"column:tx_hash" json:"hash"`
	Value       string `gorm:"column:value" json:"value"`
	Gas         uint64 `gorm:"column:gas" json:"gas"`
	GasPrice    uint64 `gorm:"column:gas_price" json:"gasPrice"`
	Nonce       uint64 `gorm:"column:nonce" json:"nonce"`
	To          string `gorm:"column:to" json:"to"`
	From        string `gorm:"column:from" json:"from"`
	Pending     bool   `gorm:"column:pending" json:"pending"`
	Data        []byte `gorm:"column:data" json:"data"`
}
