package settings

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		RunMode      string        `yaml:"run_mode" envconfig:"SERVER_RUN_MODE"`
		Port         int           `yaml:"port" envconfig:"SERVER_PORT"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
	} `yaml:server`
	Database struct {
		ConnectionString string `yaml:"connection_string" envconfig:"DATABASE_CONN_STRING"`
	} `yaml:database`
	Service struct {
		FirebaseCredentials string `yaml:"firebase_credentials" envconfig:"FIREBASE_CREDENTIALS"`
	} `yaml:service`
}

var ServiceConfig = &Config{}

func InitConfig() {
	readFile(ServiceConfig)
	readEnv(ServiceConfig)
}

func readFile(cfg *Config) {
	f, err := os.Open("config.yaml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
	log.Debug("Read configuration. ServiceConfig ", ServiceConfig)
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	log.Fatal(err)
	os.Exit(2)
}
