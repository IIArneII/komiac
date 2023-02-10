package def

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()

	DBHost = os.Getenv("KOMIAC_DB_HOST")
	DBPort = intGetenv("KOMIAC_DB_PORT", 5432)
	DBUser = os.Getenv("KOMIAC_DB_USER")
	DBPass = os.Getenv("KOMIAC_DB_PASS")
	DBName = os.Getenv("KOMIAC_DB_NAME")

	Port = intGetenv("KOMIAC_PORT", 80)
	Host = strGetenv("KOMIAC_HOST", "localhost")

	GooseDir = "./migrations"
)

func intGetenv(name string, def int) int {
	value := os.Getenv(name)

	if value == "" {
		return def
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		log.WithError(err).Errorf("failed to parse %q=%q", name, value)
		return def
	}

	return i
}

func strGetenv(name, def string) string {
	value := os.Getenv(name)

	if value == "" {
		return def
	}

	return value
}
