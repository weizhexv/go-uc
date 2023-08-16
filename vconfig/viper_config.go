package vconfig

import (
	dgi18n "dghire.com/libs/go-i18n"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"strings"
)

type config struct {
	Host    map[string]string `json:"host"`
	Service map[string]string `json:"service"`
	Email   map[string]string `json:"email"`
	Mysql   map[string]string `json:"mysql"`
	Url     map[string]string `json:"url"`
	App     map[string]string `json:"app"`
	Monitor map[string]string `json:"monitor"`
	Veriff  map[string]string `json:"veriff"`
}

func (c *config) String() string {
	j, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(j)
}

var viperConfig *config

func init() {
	initViperConfig()
	initI18nConfig()
}

func initViperConfig() {
	rootPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := rootPath + "/vconfig"
	testPath := strings.ReplaceAll(path, "/test", "")
	fmt.Printf("path: %s | test path: %s\n", path, testPath)

	viper.AddConfigPath(path)
	viper.AddConfigPath(testPath)
	viper.SetConfigName(activeConfig())
	viper.SetConfigType("yml")

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	c := new(config)
	err = viper.Unmarshal(c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("viper config: %v\n", c)
	viperConfig = c
}

func initI18nConfig() {
	dgi18n.Localize(dgi18n.WithBundle(&dgi18n.BundleCfg{
		DefaultLanguage:  language.Chinese,
		AcceptLanguage:   []language.Tag{language.English, language.Chinese},
		RootPath:         "vconfig/i18n",
		FormatBundleFile: "yaml",
		UnmarshalFunc:    yaml.Unmarshal,
	}))
}

func activeConfig() string {
	env, ok := os.LookupEnv("profile")
	if !ok || len(env) == 0 {
		panic("missing os variable: profile")
	}
	name := strings.ToLower("app-" + env)

	fmt.Printf("active config filename: %s\n", name)
	return name
}

func HostCookie() string {
	return viperConfig.Host["cookie"]
}

func ServiceName() string {
	return viperConfig.Service["name"]
}

func ServicePort() int {
	if port, err := strconv.Atoi(viperConfig.Service["port"]); err != nil {
		panic(err)
	} else {
		return port
	}
}

func HostInfra() string {
	return viperConfig.Host["infra"]
}

func EmailHrefA() string {
	return viperConfig.Email["href-a"]
}

func EmailHrefB() string {
	return viperConfig.Email["href-b"]
}

func EmailHrefS() string {
	return viperConfig.Email["href-s"]
}

func EmailHrefC() string {
	return viperConfig.Email["href-c"]
}

func EmailHrefInvite() string {
	return viperConfig.Email["href-invite"]
}

func MysqlURL() string {
	return viperConfig.Mysql["url"]
}

func MysqlMaxOpenConns() int {
	num, err := strconv.Atoi(viperConfig.Mysql["max-open-conns"])
	if err != nil {
		panic(err)
	}
	return num
}

func MysqlMaxIdleConns() int {
	num, err := strconv.Atoi(viperConfig.Mysql["max-idle-conns"])
	if err != nil {
		panic(err)
	}
	return num
}

func MysqlMaxIdleTime() int {
	num, err := strconv.Atoi(viperConfig.Mysql["max-idle-time"])
	if err != nil {
		panic(err)
	}
	return num
}

func MysqlMaxLifeTime() int {
	num, err := strconv.Atoi(viperConfig.Mysql["max-life-time"])
	if err != nil {
		panic(err)
	}
	return num
}

func UrlFgw() string {
	return viperConfig.Url["fgw"]
}

func AppName() string {
	return viperConfig.App["name"]
}

func MonitorPort() int {
	port, err := strconv.Atoi(viperConfig.Monitor["port"])
	if err != nil {
		panic(err)
	}
	return port
}

func Veriff() map[string]string {
	return viperConfig.Veriff
}
