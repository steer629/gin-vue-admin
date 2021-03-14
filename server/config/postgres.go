package config

type Postgres struct {
	Dbhost     string `mapstructure:"dbhost" json:"dbhost" yaml:"dbhost"`
	Dbuser     string `mapstructure:"dbuser " json:"dbuser " yaml:"dbuser "`
	Dbpassword string `mapstructure:"dbpassword" json:"dbpassword" yaml:"dbpassword"`
	Dbname     string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Dbport     string `mapstructure:"dbport" json:"dbport" yaml:"dbport"`
}
