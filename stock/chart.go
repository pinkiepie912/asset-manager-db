package stock

import "github.com/shopspring/decimal"

type Chart struct {
	ID        uint             `json:"id" gorm:"primaryKey;autoIncrement"`
	Stock     Stock           `json:"stock" gorm:"embedded"`
	Open      decimal.Decimal `json:"open"`
	Low       decimal.Decimal `json:"low"`
	High      decimal.Decimal `json:"high"`
	Close     decimal.Decimal `json:"close"`
	AdjClose  decimal.Decimal `json:"adjClose"`
	Volume    int             `json:"volume"`
	Timestamp int             `json:"timestamp" gorm:"index,sort:desc"`
}
