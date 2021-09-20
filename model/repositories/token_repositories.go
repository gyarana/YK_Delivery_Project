package repositories

import (
	"database/sql"
	"nix_education/model"
)

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

type TokenRepositoryI interface {
	GetByUID(uID string) (int, error)
}

type TokenRepository struct {
	db *sql.DB
}

func (t TokenRepository) GetByUID(uID string) (int, error) {
	tokensID := model.TokenIDs{}
	rows, err := t.db.Query("SELECT user_id FROM uids WHERE uid=?", uID)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		err = rows.Scan(&tokensID.UserID)
		if err != nil {
			return 0, err
		}
	}
	err = rows.Close()
	if err != nil {
		return 0, err
	}
	return tokensID.UserID, nil
}
