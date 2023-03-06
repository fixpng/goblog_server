package config

type Email struct {
	Host             string `json:"host" yaml:"host"`
	Port             int    `json:"port" yaml:"port"`
	User             string `json:"user" yaml:"user"` // 发件人邮箱
	Password         string `json:"password" yaml:"password"`
	DefaultFromEmail string `json:"default_from_email" yaml:"default_from_email-from-email"` // 默认的发件人名字
	UseSSL           bool   `json:"use_ssl" yaml:"use_ssl"`                                  // 是否使用ssl
	UserTls          bool   `json:"user_tls" yaml:"user_tls"`
}
