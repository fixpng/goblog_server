package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_Line"`      // 是否显示文件行号
	LogInConsole bool   `yaml:"log_In_Console"` // 是否显示打印路径
}
