package Database

import "database/sql"

type DatabaseManager struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
	Ssl      string
}

func (dbm *DatabaseManager) CreateConnectionString() string {
	connectString := "user=" + dbm.User
	connectString += "dbname=" + dbm.DbName
	connectString += "password=" + dbm.Password
	connectString += "host=" + dbm.Host
	connectString += "sslmode=" + dbm.Ssl

	return connectString
}

func (dbm *DatabaseManager) Connect() *sql.DB {
	db, err := sql.Open("postgres", dbm.CreateConnectionString())
	if err != nil {
		panic(err.Error())
	}

	return db
}
