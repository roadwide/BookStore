package queries

import (
	"backend/ent"
	"backend/ent/user"
	"context"
)

func (db *Query) CrateUser(userName, hash, email string) (*ent.User, error) {
	return db.User.Create().
		SetID(userName).
		SetPassword(hash).
		SetEmail(email).
		Save(context.Background())
}

func (db *Query) GetUser(userName string) (*ent.User, error) {
	return db.User.Get(context.Background(), userName)
}

func (db *Query) GetUserHash(userName string) (string, error) {
	return db.User.Query().
		Where(user.ID(userName)).
		Select(user.FieldPassword).
		String(context.Background())
}
