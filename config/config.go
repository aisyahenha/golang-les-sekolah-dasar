package config

import (
	"fmt"

	"github.com/aisyahenha/golang-les-sekolah-dasar/utils"
)

type ApiConfig struct {
	ApiHost string
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	Password string
	User     string
	Driver   string
}

type FileConfig struct {
	Env      string
	FilePath string
}

type Config struct {
	ApiConfig
	DbConfig
	FileConfig
}

func (c *Config) ReadConfig() error {
	fmt.Print("config masuk")
	vp := utils.NewViperUtil("environment", "dev", "env")
	err := vp.LoadEnv()
	fmt.Print("ini erornya", err)
	if err != nil {
		return err
	}

	c.DbConfig = DbConfig{
		Host:     vp.GetEnv("DB_HOST", "localhost"),
		Port:     vp.GetEnv("DB_PORT", "5432"),
		Name:     vp.GetEnv("DB_NAME", "postgres"),
		Password: vp.GetEnv("DB_PASSWORD", "P@ssw0rd"),
		User:     vp.GetEnv("DB_USER", "postgres"),
		Driver:   vp.GetEnv("DB_DRIVER", "postgres"),
	}

	c.ApiConfig = ApiConfig{
		ApiHost: vp.GetEnv("API_HOST", "localhost"),
		ApiPort: vp.GetEnv("API_PORT", "8888"),
	}

	c.FileConfig = FileConfig{
		Env:      vp.GetEnv("MIGRATION", "migration"),
		FilePath: vp.GetEnv("FILE_PATH", "logger.log"),
	}
	if c.DbConfig.Host == "" || c.DbConfig.Port == "" ||
		c.DbConfig.Name == "" || c.DbConfig.User == "" || c.DbConfig.Password == "" ||
		c.ApiConfig.ApiHost == "" || c.ApiConfig.ApiPort == "" || c.FileConfig.Env == "" ||
		c.FileConfig.FilePath == "" {
		return fmt.Errorf("missing required environment variables")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.ReadConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}
