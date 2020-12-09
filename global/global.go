package global

import (
	"tsetmc-gopher/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"go.uber.org/zap"
)

var(
	BRC_DB     *gorm.DB
	BRC_VP     *viper.Viper
	BRC_CONFIG config.Server
	BRC_LOG    *zap.Logger
)