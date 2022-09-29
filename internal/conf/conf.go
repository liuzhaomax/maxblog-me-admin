package conf

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"maxblog-me-admin/internal/core"
	"sync"
)

var cfg *Config
var once sync.Once

func init() {
	once.Do(func() {
		cfg = &Config{}
	})
}

func GetInstanceOfConfig() *Config {
	return cfg
}

type Config struct {
	mux        sync.Mutex
	RunMode    string     `mapstructure:"run_mode" json:"run_mode"`
	Logger     Logger     `mapstructure:"logger" json:"logger"`
	App        App        `mapstructure:"app" json:"app"`
	Server     Server     `mapstructure:"server" json:"server"`
	Upstream   Upstream   `mapstructure:"upstream" json:"upstream"`
	Downstream Downstream `mapstructure:"downstream" json:"downstream"`
}

type Logger struct {
	Color bool `mapstructure:"color" json:"color"`
}

type App struct {
	AppName string `mapstructure:"app_name" json:"app_name"`
	Version string `mapstructure:"version" json:"version"`
}

type Server struct {
	Host            string `mapstructure:"host" json:"host"`
	Port            int    `mapstructure:"port" json:"port"`
	ShutdownTimeout int    `mapstructure:"shutdown_timeout" json:"shutdown_timeout"`
}

type Downstream struct {
	MaxblogBEUser Address `mapstructure:"maxblog_be_user" json:"maxblog_be_user"`
	MaxblogBEDemo Address `mapstructure:"maxblog_be_demo" json:"maxblog_be_demo"`
}

type Upstream struct {
	MaxblogFEAdmin AddressHttp `mapstructure:"maxblog_fe_admin" json:"maxblog_fe_admin"`
}

type Address struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type AddressHttp struct {
	Protocol string `mapstructure:"protocol" json:"protocol"`
	Domain   string `mapstructure:"domain" json:"domain"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Secure   bool   `mapstructure:"secure" json:"secure"`
}

func (cfg *Config) Load(configFile string) {
	configPath := configFile
	v := viper.New()
	v.SetConfigFile(configPath)
	err := v.ReadInConfig()
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": core.GetFuncName(),
		}).Panic(core.FormatError(900, err).Error())
	}
	err = v.Unmarshal(cfg)
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": core.GetFuncName(),
		}).Panic(core.FormatError(901, err).Error())
	}
	// 录入相关微服务地址
	ctx := core.GetInstanceOfContext()
	cfg.RegisterUpStreamToContext(ctx)
	cfg.RegisterDownstreamsToContext(ctx)
	// 设置密钥或加密方式
	core.SetKeys()
	core.SetPwdEncodingOpts()
	core.SetJWTSecret("liuzhaomax")
}

func (cfg *Config) RegisterUpStreamToContext(ctx *core.Context) {
	ctx.Upstream.MaxblogFEAdmin.Host = cfg.Upstream.MaxblogFEAdmin.Host
	ctx.Upstream.MaxblogFEAdmin.Port = cfg.Upstream.MaxblogFEAdmin.Port
	ctx.Upstream.MaxblogFEAdmin.Protocol = cfg.Upstream.MaxblogFEAdmin.Protocol
	ctx.Upstream.MaxblogFEAdmin.Domain = cfg.Upstream.MaxblogFEAdmin.Domain
	ctx.Upstream.MaxblogFEAdmin.Secure = cfg.Upstream.MaxblogFEAdmin.Secure
}

func (cfg *Config) RegisterDownstreamsToContext(ctx *core.Context) {
	ctx.Downstream.MaxblogBEUser.Host = cfg.Downstream.MaxblogBEUser.Host
	ctx.Downstream.MaxblogBEUser.Port = cfg.Downstream.MaxblogBEUser.Port
	ctx.Downstream.MaxblogBEDemo.Host = cfg.Downstream.MaxblogBEDemo.Host
	ctx.Downstream.MaxblogBEDemo.Port = cfg.Downstream.MaxblogBEDemo.Port
}
