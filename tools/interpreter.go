package tools

import (
	"tsetmc-gopher/global"
	"tsetmc-gopher/models"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

// For recovery
func interpetereHandlepanic() {
	if a := recover(); a != nil {
		//fmt.Println("RECOVERED", a)
		global.BRC_LOG.Error("interpetereHandlepanic", zap.Any("err", a))
		//debug.PrintStack()
	}
}
func RealTimeDataInterpreter(tseInfoResult string) *models.RealTimeData {
	defer interpetereHandlepanic()
	var err error
	sym := &models.RealTimeData{}
	//dataChecker(tseInfoResult)
	allData := strings.Split(tseInfoResult, ";")
	if len(allData) >= 7 {
		Ov := strings.Split(allData[0], ",")
		if len(Ov) > 12 {
			if sym.CountOfExchange, err = strconv.ParseInt(Ov[8], 10, 64); err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(CountOfExchange): ", zap.Any("err", err))
			}
			if sym.DayMinPrice, err = strconv.ParseInt(Ov[7], 10, 64); err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(DayMinPrice): ", zap.Any("err", err))
			}
			if sym.DayMaxPrice, err = strconv.ParseInt(Ov[6], 10, 64); err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(DayMaxPrice): ", zap.Any("err", err))
			}
			if sym.LastPrice, err = strconv.ParseInt(Ov[2], 10, 64); err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(LastPrice): ", zap.Any("err", err))
			}
			if sym.ClosePrice, err = strconv.ParseInt(Ov[3], 10, 64); err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(ClosePrice): ", zap.Any("err", err))
			}
			if sym.Value, err = strconv.ParseInt(Ov[9], 10, 64); err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(Value): ", zap.Any("err", err))
			}
		}
		if len(allData) >= 2 {
			tableRows := strings.Split(allData[2], ",") // تابلو
			if len(tableRows) >= 4 {
				rowOne := strings.Split(tableRows[0], "@")

				if sym.RowOneCountOfBuyer, err = strconv.ParseInt(rowOne[0], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowOneCountOfBuyer): ",zap.Any("err", err))
				}
				if sym.RowOneDemandVol, err = strconv.ParseInt(rowOne[1], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowOneDemandVol): ", zap.Any("err", err))
				}
				if sym.RowOneDemandPrice, err = strconv.ParseInt(rowOne[2], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowOneDemandPrice): ", zap.Any("err", err))
				}
				if sym.RowOneCountOfSeller, err = strconv.ParseInt(rowOne[5], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowOneCountOfSeller): ", zap.Any("err", err))
				}
				if sym.RowOneSupplyVol, err = strconv.ParseInt(rowOne[4], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowOneSupplyVol): ", zap.Any("err", err))
				}
				if sym.RowOneSupplyPrice, err = strconv.ParseInt(rowOne[3], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowOneSupplyPrice): ", zap.Any("err", err))
				}
				// secound raw
				rowTwo := strings.Split(tableRows[1], "@")
				if sym.RowTwoCountOfBuyer, _ = strconv.ParseInt(rowTwo[0], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowTwoCountOfBuyer): ", zap.Any("err", err))
				}
				if sym.RowTwoDemandVol, _ = strconv.ParseInt(rowTwo[1], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowTwoDemandVol): ", zap.Any("err", err))
				}
				if sym.RowTwoDemandPrice, _ = strconv.ParseInt(rowTwo[2], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowTwoDemandPrice): ", zap.Any("err", err))
				}
				if sym.RowTwoCountOfSeller, _ = strconv.ParseInt(rowTwo[5], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowTwoCountOfSeller): ", zap.Any("err", err))
				}
				if sym.RowTwoSupplyVol, _ = strconv.ParseInt(rowTwo[4], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowTwoSupplyVol): ", zap.Any("err", err))
				}
				if sym.RowTwoSupplyPrice, _ = strconv.ParseInt(rowTwo[3], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowTwoSupplyPrice): ", zap.Any("err", err))
				}
				// third raw
				rowThree := strings.Split(tableRows[2], "@")
				if sym.RowThreeCountOfBuyer, _ = strconv.ParseInt(rowThree[0], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowThreeCountOfBuyer): ", zap.Any("err", err))
				}
				if sym.RowThreeDemandVol, _ = strconv.ParseInt(rowThree[1], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowThreeDemandVol): ", zap.Any("err", err))
				}
				if sym.RowThreeDemandPrice, _ = strconv.ParseInt(rowThree[2], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowThreeDemandPrice): ",zap.Any("err", err))
				}
				if sym.RowThreeCountOfSeller, _ = strconv.ParseInt(rowThree[5], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowThreeCountOfSeller): ", zap.Any("err", err))
				}
				if sym.RowThreeSupplyVol, _ = strconv.ParseInt(rowThree[4], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowThreeSupplyVol): ", zap.Any("err", err))
				}
				if sym.RowThreeSupplyPrice, _ = strconv.ParseInt(rowThree[3], 10, 64); err != nil {
					global.BRC_LOG.Error("RealTimeDataInterpreter(RowThreeSupplyPrice): ", zap.Any("err", err))
				}
			}
		}

		// خرید حقیقی حقوقی
		eData := strings.Split(allData[4], ",")
		if len(eData) > 9 {
			buyHagigi, err := strconv.ParseInt(eData[0], 10, 64) // خرید حقیقی
			if err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(buyHagigi): ", zap.Any("err", err))
			}
			buyHogogi, err := strconv.ParseInt(eData[1], 10, 64) // خرید حقوقی
			if err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(buyHogogi): ", zap.Any("err", err))
			}
			buydata2, err := strconv.ParseInt(eData[2], 10, 64) // مجهول
			if err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(buydata2): ", zap.Any("err", err))
			}
			cntOfBuyHagigi, err := strconv.ParseInt(eData[5], 10, 64)
			if err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(cntOfBuyHagigi): ", zap.Any("err", err))
			}
			cntOfBuyHogogi, err := strconv.ParseInt(eData[6], 10, 64)
			if err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(cntOfBuyHogogi): ", zap.Any("err", err))
			}
			sellHagigi, err := strconv.ParseInt(eData[3], 10, 64) //فروش حقیقی
			if err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(sellHagigi): ", zap.Any("err", err))
			}
			sellHogogi, err := strconv.ParseInt(eData[4], 10, 64) // فروش حقوقی
			if err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(sellHogogi): ", zap.Any("err", err))
			}
			cntOfSellHagigi, err := strconv.ParseInt(eData[8], 10, 64)
			if err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(cntOfSellHagigi): ", zap.Any("err", err))
			}
			cntOFSellHogogi, err := strconv.ParseInt(eData[9], 10, 64)
			if err != nil {
				global.BRC_LOG.Error("RealTimeDataInterpreter(cntOFSellHogogi): ", zap.Any("err", err))
			}
			sumBuy := buyHagigi + buyHogogi + buydata2
			sumSell := sellHagigi + sellHogogi
			sym.HagigiBuy = buyHagigi
			sym.HagigiSell = sellHagigi
			sym.HogogiBuy = buyHogogi
			sym.HagigiSell = sellHogogi

			sym.CntHagigiBuy = cntOfBuyHagigi
			sym.CntHogogiBuy = cntOfBuyHogogi
			sym.CntHagigiSell = cntOfSellHagigi
			sym.CntHogogiBuy = cntOFSellHogogi
			sym.Sumbuy = sumBuy
			sym.SumSel = sumSell
		}

		return sym
	}
	return nil
}
