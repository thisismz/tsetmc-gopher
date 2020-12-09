package config

type Server struct {
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	//gorm
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PostgresSql Postgres `mapstructure:"postgres" json:"postgres" yaml:"postgres"`
	//links
	Link Links `mapstructure:"links" json:"links" yaml:"links"`
	//utils
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
}
