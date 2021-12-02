package config

import (
	"github.com/spf13/viper"
	"log"
)

// Tables Binding Dynamo Tables
type Tables struct {
	ApiMaster string `yaml:"api_master"`
}

type Configurations struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	AWS struct {
		Region string `yaml:"region"`
		Tables Tables `yaml:"tables"`
		Key    string `yaml:"key"`
		Secret string `yaml:"secret"`
	} `yaml:"aws"`
}

var conf *viper.Viper

func LoadConfig(environment string) {

	// name of the config file => environment
	conf = viper.New()
	conf.SetConfigType("yaml")

	// set the file name of the configurations file
	// here, environment => configuration name
	conf.SetConfigName(environment)

	// set the path to look for the configurations file
	conf.AddConfigPath("env/")

	// Enable VIPER to read Environment Variables
	conf.AutomaticEnv()

	// find and read the configuration file
	if err := conf.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config Configurations

	err := conf.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func GetConfigurations() *viper.Viper {
	return conf
}
