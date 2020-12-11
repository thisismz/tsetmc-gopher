package models

import (
	"gorm.io/gorm"
)

type Symbols struct {
	gorm.Model
	Name          string //2
	FullName      string //3
	Key           string //0
	IsniId        string //1
	IsCompany     bool
	SymbolStatues []SymbolStatus
	RealTimeDatas []RealTimeData
}
type SymbolStatus struct {
	gorm.Model
	SymbolsID        uint
	InBuyQueue       bool
	InSellQueue      bool
	ValueOfBuyQueue  int64
	ValueOfSellQueue int64
	PowerOfBuyQueue  int64
	PowerOfSellQueue int64
}
type SymbolDailyData struct {
	gorm.Model
	SymbolsID      uint
	YesterdayPrice int64
	LastPrice      int64
	ClosePrice     int64
	DayMinPrice    int64
	DayMaxPrice    int64
	Value          int64

	HagigiBuy     int64
	HagigiSell    int64
	CntHagigiBuy  int64
	CntHagigiSell int64

	HogogiBuy     int64
	HogogiSell    int64
	CntHogogiBuy  int64
	CntHogogiSell int64
	//RealTimeSymbols []RealTimeData
}
type RealTimeData struct {
	gorm.Model
	SymbolsID       uint
	ServerDataTime  int64
	CountOfExchange int64
	LastPrice       int64
	ClosePrice      int64
	Value           int64
	DayMinPrice     int64
	DayMaxPrice     int64

	HagigiBuy     int64
	HagigiSell    int64
	CntHagigiBuy  int64
	CntHagigiSell int64

	HogogiBuy     int64
	HogogiSell    int64
	CntHogogiBuy  int64
	CntHogogiSell int64

	Sumbuy              int64
	SumSel              int64
	RowOneCountOfBuyer  int64 // تعداد خریدار
	RowOneDemandVol     int64 //حجم خرید
	RowOneDemandPrice   int64 // ردیف اول مظنه خرید
	RowOneCountOfSeller int64 // تعداد فروشنده
	RowOneSupplyVol     int64 // حجم فروش
	RowOneSupplyPrice   int64 // ردیف اول مظنه فروش

	RowTwoCountOfBuyer  int64 // تعداد خریدار
	RowTwoDemandVol     int64 //حجم خرید
	RowTwoDemandPrice   int64 // ردیف دوم مظنه خرید
	RowTwoCountOfSeller int64 // تعداد فروشنده
	RowTwoSupplyVol     int64 // حجم فروش
	RowTwoSupplyPrice   int64 // ردیف دوم مظنه فروش

	RowThreeCountOfBuyer  int64 // تعداد خریدار
	RowThreeDemandVol     int64 //حجم خرید
	RowThreeDemandPrice   int64 // ردیف سوم مظنه خرید
	RowThreeCountOfSeller int64 // تعداد فروشنده
	RowThreeSupplyVol     int64 // حجم فروش
	RowThreeSupplyPrice   int64 // ردیف سوم مظنه فروش
}

//func (realtimesymbol *RealTimeData) CreateRealTimeTable() {
//	global.BRC_DB.Create(realtimesymbol)
//	if realtimesymbol.ID <= 0 {
//		println("Error with Creat, realtimesymbol Table")
//	}
//}
//func (symboldailydata *SymbolDailyData) CreateSymbolsTable() {
//	global.BRC_DB.Create(symboldailydata)
//	if symboldailydata.ID <= 0 {
//		println("Error with Creat, symbol Table")
//	}
//}
//func (symbolstatus *SymbolStatus) CreateStatusTable() {
//	global.BRC_DB.Create(symbolstatus)
//	if symbolstatus.ID <= 0 {
//		println("Error with Creat, symbolstatus Table")
//	}
//}
//func Getsymbol(name string) *Symbols {
//
//	symStatus := &Symbols{}
//	err :=global.BRC_DB.Model(&SymbolStatus{}).Select("")
//	if err != nil {
//		return nil
//	}
//	return symStatus
//}
//func GetsymbolStatus(name string) *SymbolStatus {
//	//type result struct {
//	//	Id string
//	//	Name  string
//	//	CreatedAt  string
//	//	LastPrice string
//	//}
//	var res SymbolStatus
//	global.BRC_DB.Table("real_time_data").Select("real_time_data.id, symbols.name,real_time_data.created_at,real_time_data.last_price").Joins("INNER JOIN symbols ON real_time_data.symbols_id=symbols.id").Where("symbols.name = "+"'"+name+"'").Find(&res)
//	return &res
//}
