package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Db       Db
		App      App
		Client   Client
		Jwt      Jwt
		Google   Google
		Facebook Facebook
		Sessions Sessions
	}

	App struct {
		Host  string
		Port  int
		Stage string
	}

	Client struct {
		Host string
		Port int
	}

	// Database
	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}

	Jwt struct {
		AccessSecretKey  string
		AccessDuration   int64
		RefreshSecretKey string
		RefreshDuration  int64
	}

	Google struct {
		ClientID     string
		ClientSecret string
	}

	Facebook struct {
		ClientID     string
		ClientSecret string
	}

	Sessions struct {
		Secret string
		MaxAge int
	}
)

func LoadConfig(path string) Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file")
	}
	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal("Error loading .env file: app's port is invalid")
	}

	clientPort, err := strconv.Atoi(os.Getenv("CLIENT_PORT"))
	if err != nil {
		log.Fatal("Error loading .env file: client's port is invalid")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Error loading .env file: db's port is invalid")
	}

	jwtAccessDuration, err := strconv.ParseInt(os.Getenv("JWT_ACCESS_DURATION"), 10, 64)
	if err != nil {
		log.Fatal(`Error loading .env file: db's "jwt access duration" is invalid`)
	}

	jwtRefreshDuration, err := strconv.ParseInt(os.Getenv("JWT_REFRESH_DURATION"), 10, 64)
	if err != nil {
		log.Fatal(`Error loading .env file: db's "jwt refresh duration" is invalid`)
	}

	sessionsMaxAge, err := strconv.Atoi(os.Getenv("SESSIONS_MAX_AGE"))
	if err != nil {
		log.Fatal(`Error loading .env file: db's "sessions max age" is invalid`)
	}

	return Config{
		App: App{
			Host:  os.Getenv("APP_HOST"),
			Port:  appPort,
			Stage: os.Getenv("APP_HOST"),
		},

		Client: Client{
			Host: os.Getenv("CLIENT_HOST"),
			Port: clientPort,
		},

		Db: Db{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
			TimeZone: os.Getenv("DB_TIMEZONE"),
		},

		Jwt: Jwt{
			AccessSecretKey:  os.Getenv("JWT_ACCESS_SECRET_KEY"),
			AccessDuration:   jwtAccessDuration,
			RefreshSecretKey: os.Getenv("JWT_REFRESH_SECRET_KEY"),
			RefreshDuration:  jwtRefreshDuration,
		},

		Google: Google{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		},

		Facebook: Facebook{
			ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
			ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		},

		Sessions: Sessions{
			Secret: os.Getenv("SESSIONS_SECRET"),
			MaxAge: sessionsMaxAge,
		},
	}
}
