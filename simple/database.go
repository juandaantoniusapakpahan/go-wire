package simple

type Database struct {
	Name string
}

type DatabasePostgresql Database
type DatabaseMonggoDb Database

func NewDatabasePostgresql() *DatabasePostgresql {
	return (*DatabasePostgresql)(&Database{Name: "Postgresql"})
}

func NewDatabaseMonggoDb() *DatabaseMonggoDb {
	return (*DatabaseMonggoDb)(&Database{Name: "Mongoo"})
}

type DatabaseRepository struct {
	DatabasePostgresql *DatabasePostgresql
	DatabaseMonggoDb   *DatabaseMonggoDb
}

func NewDatabaseRepository(psql *DatabasePostgresql, mongoDb *DatabaseMonggoDb) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgresql: psql, DatabaseMonggoDb: mongoDb}
}
