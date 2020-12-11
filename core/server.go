package core

import (
	"sync"
	"time"
	"tsetmc-gopher/global"
	"tsetmc-gopher/service"
	"tsetmc-gopher/utils"
	"tsetmc-gopher/workers"
)

func RunWindowsServer() {
	utils.UpdataSymbolDbs()
	once := &sync.Once{}
	for {
		go func() {
			once.Do(func() {
				workers.Ticker(global.BRC_CONFIG.System.NumOfThread)
			})
		}()
		time.Sleep(1 * time.Second)
		service.Input("please write sym name: ")
	}

}
