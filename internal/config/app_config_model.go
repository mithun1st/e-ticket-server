package config

import "fmt"

type AppConfigModel struct {
	Server AppServer
	Db     AppDB
	Keys   AppKey
}

type AppServer struct {
	IP   string
	Port string
}

func (dbConfig AppDB) DbConnedtionString() string {
	var str string = fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		dbConfig.DbUser,
		dbConfig.DbPass,
		dbConfig.DbHost,
		dbConfig.DbPort,
		dbConfig.DbName,
	)
	return str
}

type AppDB struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

type AppKey struct {
	JwtSecretKey  string
	JwtExpiryTime int
}
