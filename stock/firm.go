package stock

import "time"

type Firm struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"not null"`
	TradingFee uint      `json:"tradingFee" gorm:"not null"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
