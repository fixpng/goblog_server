package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show-line"`      // 是否显示文件行号
	LogInConsole bool   `yaml:"log-in-console"` // 是否显示打印路径
}
