package config
type Links struct {
	TSE_TICKER_EXPORT_DATA_ADDRESS string `mapstructure:"TSE_TICKER_EXPORT_DATA_ADDRESS" json:"TSE_TICKER_EXPORT_DATA_ADDRESS" yaml:"TSE_TICKER_EXPORT_DATA_ADDRESS"`
	TSE_TICKER_ADDRESS             string `mapstructure:"TSE_TICKER_ADDRESS" json:"TSE_TICKER_ADDRESS" yaml:"TSE_TICKER_ADDRESS"`
	TSE_ISNT_INFO_URL              string `mapstructure:"TSE_ISNT_INFO_URL" json:"TSE_ISNT_INFO_URL" yaml:"TSE_ISNT_INFO_URL"`
	TSE_CLIENT_TYPE_DATA_URL       string `mapstructure:"TSE_CLIENT_TYPE_DATA_URL" json:"TSE_CLIENT_TYPE_DATA_URL" yaml:"TSE_CLIENT_TYPE_DATA_URL"`
	TSE_SYMBOL_ID_URL              string `mapstructure:"TSE_SYMBOL_ID_URL" TSE_SYMBOL_ID_URL:"host" TSE_SYMBOL_ID_URL:"TSE_SYMBOL_ID_URL"`
	TSE_DAILY_PRICE_INFO           string `mapstructure:"host" json:"host" yaml:"host"`
	TSE_CHART_IMAGE_URL            string `mapstructure:"TSE_CHART_IMAGE_URL" json:"TSE_CHART_IMAGE_URL" yaml:"TSE_CHART_IMAGE_URL"`
	TSE_FILTER_ALL_DATA_NAME       string `mapstructure:"TSE_FILTER_ALL_DATA_NAME" json:"TSE_FILTER_ALL_DATA_NAME" yaml:"TSE_FILTER_ALL_DATA_NAME"`
}