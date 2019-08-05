package config


type Web struct {
	
}

type DB struct {
	DB2250 MySQL
	DB1580 MySQL
}

type MySQL struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
}
