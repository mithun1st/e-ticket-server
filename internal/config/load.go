package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() (*AppConfigModel, error) {

	// // Get the path to the current file (server/dev/main.go)
	// _, currentFile, _, _ := runtime.Caller(0)
	// currentDir := filepath.Dir(currentFile)

	// // Navigate up to project root (adjust the number of ".." as needed)
	// projectRoot := filepath.Join(currentDir, "..", "..")
	// envPath := filepath.Join(projectRoot, ".env")

	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	server := AppServer{
		IP:   os.Getenv("APP_IP"),
		Port: os.Getenv("APP_PORT"),
	}

	db := AppDB{
		DbHost: os.Getenv("DB_HOST_DEV"),
		DbPort: os.Getenv("DB_PORT_DEV"),

		DbUser: os.Getenv("DB_USER_DEV"),
		DbPass: os.Getenv("DB_PASSWORD_DEV"),
		DbName: os.Getenv("DB_NAME_DEV"),
	}

	// expiry, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRY_TIME"), 10, 64)

	jwt := AppKey{
		JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
		// JwtExpiryTime: int(expiry),
	}

	appConfig := AppConfigModel{
		Server: server, Db: db, Keys: jwt,
	}

	return &appConfig, nil
}
