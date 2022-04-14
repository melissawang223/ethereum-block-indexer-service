package models

// event logs data structure
type EventLogsModel struct {
	ID          int64  `gorm:"column:id; primary_key"`
	BlockNumber uint64 `gorm:"column:block_number" json:"blockNumber"`
	Hash        string `gorm:"column:tx_hash" json:"hash"`
	Index       uint   `gorm:"column:index" json:"index"`
	Data        string `gorm:"column:data" json:"data"`
}
