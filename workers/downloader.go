package workers

import (
	"tsetmc-gopher/global"
	"tsetmc-gopher/models"
	tool "tsetmc-gopher/tools"
	"fmt"
	"go.uber.org/zap"
)

// For recovery
func handlepanic() {
	if a := recover(); a != nil {
		//fmt.Println("RECOVERED", a)
		global.BRC_LOG.Error("downloader :", zap.Any("err", a))
		//debug.PrintStack()
	}
}

//RealTimeDataDownloaderName get symbol data with name of symbol in struct
func RealTimeDataDownloaderName(symName string) *models.RealTimeData {
	defer handlepanic()
	var sym models.Symbols
	reslut := global.BRC_DB.Where("name = ?", "symName").Find(&sym)
	println(reslut.RowsAffected)
	s1 := fmt.Sprintf(global.BRC_CONFIG.Link.TSE_ISNT_INFO_URL, sym.Key)
	urlData := tool.Responser(s1)
	temp := tool.RealTimeDataInterpreter(urlData)
	return temp
}

//RealTimeDataDownloaderKey get symbol data with kay
func RealTimeDataDownloaderKey(sym models.Symbols) *models.RealTimeData {
	defer handlepanic()
	s1 := fmt.Sprintf(global.BRC_CONFIG.Link.TSE_ISNT_INFO_URL, sym.Key)
	urlData := tool.Responser(s1)
	temp := tool.RealTimeDataInterpreter(urlData)
	return temp
}

//GetAllSymbolRealtimeDataInMap get all data in map format for save in json if link is broken get symbol name
func GetAllSymbolRealtimeDataInMap() ([]models.RealTimeData) {
	defer handlepanic()
	var temp []models.RealTimeData
	sym := []models.Symbols{}
	for i := 0; i < len(sym); i++ {
		symData := RealTimeDataDownloaderKey(sym[i])
		if symData != nil {
			symData.SymbolsID = sym[i].ID
			temp = append(temp,*symData)
			//println("org: ",result.Error)
			if symData == nil {
				global.BRC_LOG.Error("GetAllSymbolRealtimeDataInMap")
			}

		}
	}
	return temp
}

