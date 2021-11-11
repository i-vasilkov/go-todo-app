package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Mongo MongoConfig
	Http  HttpConfig
	Auth  AuthConfig
	Jwt   JwtConfig
}

type MongoConfig struct {
	DbName   string `mapstructure:"MONGODB_DATABASE"`
	UserName string `mapstructure:"MONGO_INITDB_ROOT_USERNAME"`
	Password string `mapstructure:"MONGO_INITDB_ROOT_PASSWORD"`
	Host     string `mapstructure:"MONGO_HOST"`
	Port     string `mapstructure:"MONGO_PORT"`
}

func (mc *MongoConfig) GetURI() string {
	return fmt.Sprintf("%s:%s", mc.Host, mc.Port)
}

type HttpConfig struct {
	Host         string        `mapstructure:"HTTP_HOST"`
	Port         string        `mapstructure:"HTTP_PORT"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
}

func (hc *HttpConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", hc.Host, hc.Port)
}

type AuthConfig struct {
	PwdSalt string `mapstructure:"PASSWORD_SALT"`
}

type JwtConfig struct {
	Signature string        `mapstructure:"JWT_SIGN"`
	Ttl       time.Duration `mapstructure:"ttl"`
}

func Init(paths ...string) (Config, error) {
	if err := ReadConfigFiles(paths); err != nil {
		return Config{}, err
	}

	cfg, err := UnmarshalConfig()
	return cfg, err
}

func ReadConfigFiles(paths []string) error {
	for _, path := range paths {
		if err := InitCfgFile(path); err != nil {
			return err
		}
	}
	return nil
}

func InitCfgFile(cfgFile string) error {
	if cfgFile == "" {
		return errors.New("received empty cfg file path")
	}

	viper.SetConfigFile(cfgFile)
	return viper.MergeInConfig()
}

func UnmarshalConfig() (Config, error) {
	var cfg Config

	if err := UnmarshalMongoCfg(&cfg); err != nil {
		return cfg, err
	}

	if err := UnmarshalHttpCfg(&cfg); err != nil {
		return cfg, err
	}

	if err := UnmarshalAuthCfg(&cfg); err != nil {
		return cfg, err
	}

	if err := UnmarshalJwtCfg(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func UnmarshalMongoCfg(cfg *Config) error {
	return viper.Unmarshal(&cfg.Mongo)
}

func UnmarshalHttpCfg(cfg *Config) error {
	if err := viper.Unmarshal(&cfg.Http); err != nil {
		return err
	}
	return viper.UnmarshalKey("http", &cfg.Http)
}

func UnmarshalAuthCfg(cfg *Config) error {
	return viper.Unmarshal(&cfg.Auth)
}

func UnmarshalJwtCfg(cfg *Config) error {
	if err := viper.Unmarshal(&cfg.Jwt); err != nil {
		return err
	}
	return viper.UnmarshalKey("jwt", &cfg.Jwt)
}
