package models

type (
	LogConfig struct {
		Path  				string  `json:"path" env:"LOG_PATH"`
		File  				string  `json:"file" env:"LOG_FILE"`
		Level 				int     `json:"level" env:"LOG_LEVEL"`
		OutputConsole       bool    `json:"output_console"`
		OutputFile          bool    `json:"output_file"`
	}

	Configuration struct {
		Log              LogConfig               			 `json:"log"`
	}

)