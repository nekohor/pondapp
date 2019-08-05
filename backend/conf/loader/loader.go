package loader

import (
	"github.com/nekohor/pondapp/backend/conf/config"
	"github.com/spf13/viper"
)

// load config
func LoadConfig(fpath string) (c *config.Config, err error) {
	v := viper.New()
	v.SetConfigFile(fpath)
	v.SetConfigType("toml")
	if err1 := v.ReadInConfig(); err1 != nil {
		err = err1
		return
	}
	c = &config.Config{}
	c.Web.StaticPath = v.GetString("web.static_path")
	c.DB.MYSQL1580 = v.GetString("web.domain")

	return
}
