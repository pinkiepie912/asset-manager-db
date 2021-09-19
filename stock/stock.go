package stock

import "time"

type Stock struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Symbol       string    `json:"symbol" gorm:"unique"`
	Firm         Firm      `json:"firm" gorm:"embedded"`
	ExpenseRatio float32   `json:"expenseRatio" gorm:"expenseRatio"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime"`

	// ChartMeta
	Currency             string `json:"currency"`
	ExchangeName         string `json:"exchangeName"`
	QuoteType            string `json:"instrumentType"`
	FirstTradeDate       int    `json:"firstTradeDate"`
	Gmtoffset            int    `json:"gmtoffset"`
	Timezone             string `json:"timezone"`
	ExchangeTimezoneName string `json:"exchangeTimezoneName"`
}
