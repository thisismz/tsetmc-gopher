package initialize

import (
	"tsetmc-gopher/global"
	"tsetmc-gopher/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func Gorm() *gorm.DB {
	switch global.BRC_CONFIG.System.DbType {
	case "mysql":
		println("Error IF U Want Mysql uncommant line 36-57")
		return nil
	default:
		return GormPostgress()
	}
}

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.Symbols{},
		models.SymbolDailyData{},
		models.RealTimeData{},
		models.SymbolStatus{},
	)
	if err != nil {
		global.BRC_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.BRC_LOG.Info("register table success")
}
//func GormMysql() *gorm.DB {
//	m := global.BRC_CONFIG.Mysql
//	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Dbname + "?" + m.Config
//	mysqlConfig := mysql.Config{
//		DSN:                       dsn,   // DSN data source name
//		DefaultStringSize:         191,   // Default length of string type field
//		DisableDatetimePrecision:  true,  // Disable datetime precision, which is not supported by databases before MySQL 5.6
//		DontSupportRenameIndex:    true,  // When renaming indexes, delete and create new ones. Databases before MySQL 5.7 and MariaDB do not support renaming indexes
//		DontSupportRenameColumn:   true,  // Use `change` to rename the column, the database before MySQL 8 and MariaDB do not support renaming the column
//		SkipInitializeWithVersion: false, // Automatic configuration based on version
//	}
//	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
//		global.BRC_LOG.Error("MySQL", zap.Any("err", err))
//		os.Exit(0)
//		return nil
//	} else {
//		sqlDB, _ := db.DB()
//		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
//		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
//		return db
//	}
//}
func GormPostgress() *gorm.DB {
	m := global.BRC_CONFIG.PostgresSql
	dsn := "host="+m.Host+" user="+m.Username+" dbname="+m.Dbname+" sslmode=disable password="+m.Password
	if db, err := gorm.Open(postgres.Open(dsn), gormConfig(m.LogMode)); err != nil {
		global.BRC_LOG.Error("postgres ",zap.Any("err", err))
		println("Connecting error: ",err)
		os.Exit(0)
		return nil
	} else {
		println("Successfully connecting ")
		return db
	}
}

func gormConfig(mod bool) *gorm.Config {
	if mod {
		return &gorm.Config{
			Logger:logger.Default.LogMode(logger.Info),

		}
	} else {
		return &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),

		}
	}
}