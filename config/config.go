package config

func MongoDBConfig() string {
	return "mongodb://root:example@localhost:27017"
}

func PostgresConfig() string {
	return "postgres://gabrielvillarinho:strongpassword@localhost:5432/school"
}
