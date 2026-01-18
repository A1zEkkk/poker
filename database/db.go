package database

import (
	"database/sql"
	"poker/config"

	_ "github.com/jackc/pgx/v5/stdlib"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	Postgres *sql.DB
	Mongo    *mongo.Client
}

func NewDB(cfg *config.Config) (*DB, error) {
	// Postgres
	pg, err := sql.Open("pgx", cfg.PostgresDNS)
	if err != nil {
		return nil, err
	}

	if err := pg.Ping(); err != nil {
		return nil, err
	}
	// Mongo
	/*mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoAddr))
	  if err != nil {
	      log.Fatal(err)
	  }*/

	return &DB{
		Postgres: pg,
		/*Mongo:    mongoClient,*/
	}, nil
}
