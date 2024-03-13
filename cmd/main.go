package main

import (
	"database/sql"
	"pocket-link/api"
	"pocket-link/config"
	db "pocket-link/db/sqlc"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("cannot load configs")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Msgf("cannot connect to db: %v", err)
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server instance")
	}

	if err = server.Start(); err != nil {
		panic(err)
	}

	log.Info().Msg("listening to :8080")
}
