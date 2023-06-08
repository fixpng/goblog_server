package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site-info"`
	QQ       QQ       `yaml:"qq"`
	QiNiu    QiNiu    `yaml:"qiniu"`
	Email    Email    `yaml:"email"`
	Jwt      Jwt      `yaml:"jwt"`
	Upload   Upload   `yaml:"upload"`
	Redis    Redis    `json:"redis"`
	ES       ES       `json:"es"`
}
