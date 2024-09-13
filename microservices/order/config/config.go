package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}
	return port
}

func getEnvironmentValue(key string) string {
	envValue := os.Getenv(key)
	if envValue == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}

	return envValue
}

func GetPaymentServiceURL() string {
	return getEnvironmentValue("PAYMENT_SERVICE_URL")
}
