package models

import "gorm.io/gorm"

type BazarStatus struct {
	gorm.Model
	CountOfBuyQueue     int64
	CountOfSellQueue    int64
	ValueOfAllBuyQueue  int64
	ValueOfAllSellQueue int64
	CountOfHogogiSell   int64
	CountOfHogogiBuy    int64
	HogogiBuyVal        int64
	HogogiSellVal       int64
}
