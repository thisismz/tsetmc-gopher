package main

import (
	"tsetmc-gopher/core"
	"tsetmc-gopher/global"
	"tsetmc-gopher/initialize"
)

func main() {
	global.BRC_VP = core.Viper("config.yaml")
	global.BRC_LOG = core.Zap()
	global.BRC_DB = initialize.Gorm()
	initialize.MysqlTables(global.BRC_DB)
	db, _ := global.BRC_DB.DB()
	defer db.Close()
	core.RunWindowsServer()
}
