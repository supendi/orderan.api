package env

import "os"

//GetJWTSecretKey return a string of JWT secret key from env
func GetJWTSecretKey() string {
	return os.Getenv("JWT_KEY")
}

//GetDBTestConstring return database test connection string from env
func GetDBTestConstring() string {
	return os.Getenv("ORDERAN_DB_TEST")
}

//GetDBConstring return database connection string from env
func GetDBConstring() string {
	return os.Getenv("ORDERAN_DB")
}
