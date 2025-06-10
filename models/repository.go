package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(b string) (*DogBreed, error)
	GetDogOfMonthById(id int) (*DogOfMonth, error)
}

type mysqlRepository struct {
	Db *sql.DB
}

func newMysqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{
		Db: conn,
	}
}

type testRepository struct {
	Db *sql.DB
}

func newTestRepository(conn *sql.DB) Repository {
	return &testRepository{
		Db: conn,
	}
}
