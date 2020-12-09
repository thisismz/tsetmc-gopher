package workers

import (
	"tsetmc-gopher/models"
)

func Analyser(data *models.RealTimeData) *models.SymbolStatus {
	symdata := &models.SymbolStatus{}
	symdata.InBuyQueue, symdata.InSellQueue, symdata.PowerOfSellQueue, symdata.PowerOfBuyQueue = queueAnalyser(data)
	symdata.ValueOfBuyQueue, symdata.ValueOfSellQueue = valueCalculator(data)
	return symdata
}

func queueAnalyser(data *models.RealTimeData) (safekharid bool, safeforosh bool, powerOfSeller int64, powerOfbuyer int64) {
	defer handlepanic()
	var kharid = false
	var forosh = false
	var powerOfSell int64
	var powerOfbuy int64
	if data.DayMinPrice <= data.RowOneSupplyPrice && (data.RowOneSupplyVol*2) > data.RowOneDemandVol {
		if data.RowOneSupplyVol != 0 && data.RowOneDemandVol != 0 {
			powerOfSell = (data.RowOneSupplyVol / data.RowOneDemandVol) / 100
		}
		forosh = true
	} else if data.DayMaxPrice >= data.RowOneDemandPrice && (data.RowOneDemandVol*2) > data.RowOneSupplyVol {
		if data.RowOneSupplyVol != 0 && data.RowOneDemandVol != 0 {
			powerOfbuy = (data.RowOneDemandVol / data.RowOneSupplyVol) / 100
		}
		kharid = true
	} else {
		if data.RowOneSupplyVol != 0 && data.RowOneDemandVol != 0 {
			powerOfbuyer = (data.RowOneDemandVol / data.RowOneSupplyVol) / 100
			powerOfSell = (data.RowOneSupplyVol / data.RowOneDemandVol) / 100
		}
		kharid = false
		forosh = false
	}
	return kharid, forosh, powerOfSell, powerOfbuy
}
func valueCalculator(data *models.RealTimeData) (buyVol int64, sellVol int64) {
	valueOfBuyQueue := data.RowOneDemandVol * data.RowOneDemandPrice
	valueOfSellQueue := data.RowOneSupplyVol * data.RowOneSupplyPrice
	return valueOfBuyQueue, valueOfSellQueue
}
func BazarAnalyser(data *models.SymbolStatus) {

}
