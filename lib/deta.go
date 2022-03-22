package lib

import (
	"os"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

// Create a new Deta instance.
func DetaDB() (*deta.Deta, error) {
	d, err := deta.New(deta.WithProjectKey(os.Getenv("DETA_PROJECT_KEY")))
	if err != nil {
		return nil, err
	}

	return d, nil
}

// returns the users deta base
func UsersBase() (*base.Base, error) {
	db, err := DetaDB()
	if err != nil {
		return nil, err
	}

	return base.New(db, "users")
}
