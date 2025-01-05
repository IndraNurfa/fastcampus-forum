package configs

type (
	Config struct {
		Service  Service  `maspstructure:"service"`
		Database Database `mapstructure:"database"`
	}

	Service struct {
		Port string `mapstructure:"port"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourcename"`
	}
)
