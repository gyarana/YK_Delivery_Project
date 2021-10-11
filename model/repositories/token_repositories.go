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
	GetUIByUID(uID string) (int, error)
	UpdateUI(userID int, uID string) (int, error)
	DeleteUI(userID int) error
}

type TokenRepository struct {
	db *sql.DB
}

func (t TokenRepository) GetUIByUID(uID string) (int, error) {
	tokensID := model.TokenIDs{}
	rows, err := t.db.Query("SELECT user_id FROM uid WHERE uid=?", uID)
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
func (t TokenRepository) UpdateUI(userID int, uID string) error {
	_, err := t.db.Exec("UPDATE uid SET user_id=?, uid=?", userID, uID)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func (t TokenRepository) DeleteUI(userID int) error {
	_, err := t.db.Exec("UPDATE uid SET uid=NULL WHERE user_id=?", userID)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
