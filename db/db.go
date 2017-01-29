package db

import (
	"flag"
	"io/ioutil"

	yaml "gopkg.in/yaml.v1"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB                 *gorm.DB
	defaultConfigFile  = "dbconfig.yml"
	defaultEnvironment = "development"
)

type DBConfigure struct {
	Dialect    string `yaml:"dialect"`
	DataSource string `yaml:"datasource"`
	Dir        string `yaml:"dir"`
	TableName  string `yaml:"table"`
	SchemaName string `yaml:"schema"`
}

func readConfig(configFile string, environment string) (*DBConfigure, error) {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	config := make(map[string]*DBConfigure)
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config[environment], nil
}

func init() {
	configFile := flag.String("dbconfig", defaultConfigFile, "--dbconfig=dbconfig.yml")
	environment := flag.String("environment", defaultEnvironment, "--env=development")

	dbConfig, err := readConfig(*configFile, *environment)
	if err != nil {
		panic(err)
	}

	DB, err = gorm.Open(dbConfig.Dialect, dbConfig.DataSource)
	if err != nil {
		panic(err)
	}
}
