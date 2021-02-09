package config

import (
	"github.com/joho/godotenv"
	"os"
	config "primero_test/properties"
	"strconv"
)

var NAME_FILE_ENVIRONMENT = ".env"
var instancePropertiesFactory *PropertiesFactory

func NewPropertiesFactory() PropertiesFactory {
	if instancePropertiesFactory == nil {
		instancePropertiesFactory = &PropertiesFactory{}
	}
	err := godotenv.Load(NAME_FILE_ENVIRONMENT)
	if err != nil {
		panic(err)
	}

	return *instancePropertiesFactory
}

type PropertiesFactory struct{}

func (r *PropertiesFactory) PropertiesSystemFactory() config.PropertiesSystem {

	valuePort, _ := strconv.ParseInt(os.Getenv("DEFAULT_PORT"), 10, 32)

	return config.PropertiesSystem{
		int(valuePort),
	}
}
