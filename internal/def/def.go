package def

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pkg/errors"

	"github.com/powerman/structlog"
)

var (
	log = structlog.New()

	DBHost = os.Getenv("KOMIAC_DB_HOST")
	DBPort = intGetenv("KOMIAC_DB_PORT", 5431)
	DBUser = os.Getenv("KOMIAC_DB_USER")
	DBPass = os.Getenv("KOMIAC_DB_PASS")
	DBName = os.Getenv("KOMIAC_DB_NAME")

	Port = intGetenv("KOMIAC_PORT", 5000)
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
		log.Err(errors.Errorf("failed to parse %q=%q as int: %v", name, value, err))
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
