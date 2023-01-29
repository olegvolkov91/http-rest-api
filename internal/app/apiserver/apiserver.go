package apiserver

import (
	"database/sql"
	"github.com/olegvolkov91/http-rest-api/internal/app/store/sqlstore"
	"net/http"
)

// Start ...
func Start(config *Config) error {
	db, err := NewDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	s := NewServer(store)
	s.logger.Info("starting api server")
	return http.ListenAndServe(config.BindAddr, s)
}

func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
