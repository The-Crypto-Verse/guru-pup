package lib

import (
	"github.com/The-Crypto-Verse/guru-pup/types"
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

// Create a new Deta instance.
func DetaDB() (*deta.Deta, error) {
	d, err := deta.New(deta.WithProjectKey(DETA_PROJECT_KEY))
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

// gets all of the users in db
func GetAllUser() ([]types.UserProps, error) {
	var users []types.UserProps

	usersBase, _ := UsersBase()

	_, err := usersBase.Fetch(
		&base.FetchInput{
			Q:     base.Query{},
			Dest:  &users,
			Limit: 0,
		},
	)
	if err != nil {
		return []types.UserProps{}, err
	}

	return users, nil
}
