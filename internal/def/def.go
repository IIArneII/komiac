package def

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New().WithField("module", "DTO")

	DBHost = "postgres"
	DBPort = 5432
	DBUser = "postgres"
	DBPass = "postgres"
	DBName = "postgres"

	Port = 80
	Host = "localhost"

	GooseDir = "./migrations"
)

func init() {
	content, err := ioutil.ReadFile(".env")
	if err != nil {
		log.WithError(err).Fatal("Failed to read contents of file '.env'")
	}

	lines := strings.Split(strings.ReplaceAll(string(content), " ", ""), "\n")

	for _, l := range lines {
		l = strings.ReplaceAll(l, "\r", "")
		if len(l) >= 3 {
			if v := strings.Split(l, "="); len(v) == 2 {
				os.Setenv(v[0], v[1])
			}
		}
	}

	setEnvs()
}

func setEnvs() {
	DBHost = os.Getenv("KOMIAC_DB_HOST")
	DBPort = intGetenv("KOMIAC_DB_PORT", 5432)
	DBUser = os.Getenv("KOMIAC_DB_USER")
	DBPass = os.Getenv("KOMIAC_DB_PASS")
	DBName = os.Getenv("KOMIAC_DB_NAME")

	Port = intGetenv("KOMIAC_PORT", 80)
	Host = strGetenv("KOMIAC_HOST", "localhost")
}

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
