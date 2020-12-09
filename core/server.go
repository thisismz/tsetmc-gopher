package core

import (
	"tsetmc-gopher/global"
	"tsetmc-gopher/workers"
)

func RunWindowsServer() {
	workers.Ticker(global.BRC_CONFIG.System.NumOfThread)
}
