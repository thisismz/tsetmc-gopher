
package config

type System struct {
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	NumOfThread int    `mapstructure:"thread" json:"thread" yaml:"thread"`
}
