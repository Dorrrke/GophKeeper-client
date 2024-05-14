package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	errText "github.com/Dorrrke/GophKeeper-client/internal/domain/errors"
	"github.com/Dorrrke/GophKeeper-client/internal/domain/models"
	"github.com/mattn/go-sqlite3"
)

var (
	ErrUserAlredyExist  = errors.New(errText.UserExistsError)
	ErrUserNotExist     = errors.New(errText.UserNotExistError)
	ErrCardAlredyExist  = errors.New(errText.CardExistsError)
	ErrLoginAlredyExist = errors.New(errText.LoginExistsError)
	ErrTextAlredyExist  = errors.New(errText.TextExistsError)
	ErrBinAlredyExist   = errors.New(errText.BinDataExistsError)
	ErrCardNotExist     = errors.New(errText.CardNotExistsError)
	ErrLoginNotExist    = errors.New(errText.LoginNotExistsError)
	ErrTextNotExist     = errors.New(errText.TextNotExistsError)
	ErrBinDataNotExist  = errors.New(errText.BinDataNotExistsError)
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUser(ctx context.Context, user models.UserModel) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO users(login, hash) VALUES(?,?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, user.Login, user.Hash)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, ErrUserAlredyExist
		}
		return 0, err
	}
	cID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return cID, nil
}

func (s *Storage) GetUserHash(ctx context.Context, login string) (int64, string, error) {
	stmt, err := s.db.Prepare("SELECT uId, hash FROM users WHERE login = ?")
	if err != nil {
		return -1, "", err
	}
	row := stmt.QueryRowContext(ctx, login)
	var uID int64
	var hash string
	err = row.Scan(&uID, &hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, "", ErrUserNotExist
		}
		return -1, "", err
	}
	return uID, hash, nil
}

func (s *Storage) SaveCard(ctx context.Context, card models.CardModel, uID int64) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO cards(name, number, date, cvv, uId, deleted, last_update) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}

	t := time.Now()
	res, err := stmt.ExecContext(ctx, card.Name, card.Number, card.Date, card.CVVCode, uID, false, t.Format(time.RFC3339))
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			fmt.Println("exist")
			return 0, ErrCardAlredyExist
		}

		return 0, err
	}
	cID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return cID, nil
}

func (s *Storage) SaveLogin(ctx context.Context, loginData models.LoginModel, uID int64) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO logins(name, login, password, uId, deleted, last_update) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	t := time.Now()
	res, err := stmt.ExecContext(ctx, loginData.Name, loginData.Login, loginData.Password, uID, false, t.Format(time.RFC3339))
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, ErrLoginAlredyExist
		}
		return 0, err
	}
	cID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return cID, nil
}

func (s *Storage) SaveText(ctx context.Context, textData models.TextDataModel, uID int64) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO text_data(name, data, uId, deleted, last_update) VALUES(?,?,?,?,?)")
	if err != nil {
		return 0, err
	}

	t := time.Now()
	res, err := stmt.ExecContext(ctx, textData.Name, textData.Data, uID, false, t.Format(time.RFC3339))
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, ErrTextAlredyExist
		}
		return 0, err
	}
	cID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return cID, nil
}

func (s *Storage) SaveBin(ctx context.Context, binData models.BinaryDataModel, uID int64) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO binares_data(name, data, uId, deleted, last_update) VALUES(?,?,?,?,?)")
	if err != nil {
		return 0, err
	}

	t := time.Now()
	res, err := stmt.ExecContext(ctx, binData.Name, binData.Data, uID, false, t.Format(time.RFC3339))
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, ErrBinAlredyExist
		}
		return 0, err
	}
	cID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return cID, nil
}

func (s *Storage) GetAllCards(ctx context.Context, uID int64) ([]models.CardModel, error) {
	stmt, err := s.db.Prepare("SELECT name, number, date, cvv FROM cards WHERE uId=? AND deleted=0")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cards []models.CardModel
	for rows.Next() {
		var card models.CardModel
		err := rows.Scan(&card.Name, &card.Number, &card.Date, &card.CVVCode)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (s *Storage) GetAllLogins(ctx context.Context, uID int64) ([]models.LoginModel, error) {
	stmt, err := s.db.Prepare("SELECT name, login, password FROM logins WHERE uId=? AND deleted=0")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var logins []models.LoginModel
	for rows.Next() {
		var login models.LoginModel
		err := rows.Scan(&login.Name, &login.Login, &login.Password)
		if err != nil {
			return nil, err
		}
		logins = append(logins, login)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return logins, nil
}

func (s *Storage) GetAllTextData(ctx context.Context, uID int64) ([]models.TextDataModel, error) {
	stmt, err := s.db.Prepare("SELECT name, data FROM text_data WHERE uId=? AND deleted=0")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tData []models.TextDataModel
	for rows.Next() {
		var data models.TextDataModel
		err := rows.Scan(&data.Name, &data.Data)
		if err != nil {
			return nil, err
		}
		tData = append(tData, data)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tData, nil
}

func (s *Storage) GetAllBin(ctx context.Context, uID int64) ([]models.BinaryDataModel, error) {
	stmt, err := s.db.Prepare("SELECT name, data FROM binares_data WHERE uId=? AND deleted=0")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var bData []models.BinaryDataModel
	for rows.Next() {
		var data models.BinaryDataModel
		err := rows.Scan(&data.Name, &data.Data)
		if err != nil {
			return nil, err
		}
		bData = append(bData, data)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return bData, nil
}

func (s *Storage) GetCardByName(ctx context.Context, name string, uID int64) (models.CardModel, error) {
	stmt, err := s.db.Prepare("SELECT name, number, date, cvv FROM cards WHERE name = ? AND uId = ? AND deleted=0")
	if err != nil {
		return models.CardModel{}, err
	}

	row := stmt.QueryRowContext(ctx, name, uID)
	var card models.CardModel
	err = row.Scan(&card.Name, &card.Number, &card.Date, &card.CVVCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.CardModel{}, ErrCardNotExist
		}
		return models.CardModel{}, err
	}
	return card, nil
}

func (s *Storage) GetLoginByName(ctx context.Context, name string, uID int64) (models.LoginModel, error) {
	stmt, err := s.db.Prepare("SELECT name, login, password FROM logins WHERE name = ? AND uId = ? AND deleted=0")
	if err != nil {
		return models.LoginModel{}, err
	}

	row := stmt.QueryRowContext(ctx, name, uID)
	var login models.LoginModel
	err = row.Scan(&login.Name, &login.Login, &login.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.LoginModel{}, ErrLoginNotExist
		}
		return models.LoginModel{}, err
	}
	return login, nil
}

func (s *Storage) GetTextDataByName(ctx context.Context, name string, uID int64) (models.TextDataModel, error) {
	stmt, err := s.db.Prepare("SELECT name, data FROM text_data WHERE name = ? AND uId = ? AND deleted=0")
	if err != nil {
		return models.TextDataModel{}, err
	}

	row := stmt.QueryRowContext(ctx, name, uID)
	var data models.TextDataModel
	err = row.Scan(&data.Name, &data.Data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.TextDataModel{}, ErrTextNotExist
		}
		return models.TextDataModel{}, err
	}
	return data, nil
}

func (s *Storage) GetBinByName(ctx context.Context, name string, uID int64) (models.BinaryDataModel, error) {
	stmt, err := s.db.Prepare("SELECT name, data FROM binares_data WHERE name = ? AND uId = ? AND deleted=0")
	if err != nil {
		return models.BinaryDataModel{}, err
	}

	row := stmt.QueryRowContext(ctx, name, uID)
	var data models.BinaryDataModel
	err = row.Scan(&data.Name, &data.Data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.BinaryDataModel{}, ErrBinDataNotExist
		}
		return models.BinaryDataModel{}, err
	}
	return data, nil
}

func (s *Storage) DeleteCard(ctx context.Context, name string, uID int64) error {
	stmt, err := s.db.Prepare("UPDATE cards SET deleted = 1, last_update = ? WHERE name = ? AND uId = ?")
	if err != nil {
		return err
	}
	t := time.Now()
	_, err = stmt.ExecContext(ctx, t.Format(time.RFC3339), name, uID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteLogin(ctx context.Context, name string, uID int64) error {
	stmt, err := s.db.Prepare("UPDATE logins SET deleted = 1, last_update = ?  WHERE name = ? AND uId = ?")
	if err != nil {
		return err
	}
	t := time.Now()
	_, err = stmt.ExecContext(ctx, t.Format(time.RFC3339), name, uID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteText(ctx context.Context, name string, uID int64) error {
	stmt, err := s.db.Prepare("UPDATE text_data SET deleted = 1, last_update = ?  WHERE name = ? AND uId = ?")
	if err != nil {
		return err
	}
	t := time.Now()
	_, err = stmt.ExecContext(ctx, t.Format(time.RFC3339), name, uID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteBin(ctx context.Context, name string, uID int64) error {
	stmt, err := s.db.Prepare("UPDATE binares_data SET deleted = 1, last_update = ?  WHERE name = ? AND uId = ?")
	if err != nil {
		return err
	}
	t := time.Now()
	_, err = stmt.ExecContext(ctx, t.Format(time.RFC3339), name, uID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateCard(ctx context.Context, card models.CardModel, uID int64) error {
	stmt, err := s.db.Prepare("UPDATE cards SET name = ?, number = ?, date = ?, cvv = ?, last_update = ? WHERE name = ? and uId = ?")
	if err != nil {
		return err
	}
	lTime := time.Now().Format(time.RFC3339)
	_, err = stmt.ExecContext(ctx, card.Name, card.Number, card.Date, card.CVVCode, lTime, card.Name, uID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateLogin(ctx context.Context, auth models.LoginModel, uID int64) error {
	stmt, err := s.db.Prepare("UPDATE logins SET name = ?, login = ?, password = ?, last_update = ? WHERE name = ? and uId = ?")
	if err != nil {
		return err
	}
	lTime := time.Now().Format(time.RFC3339)
	_, err = stmt.ExecContext(ctx, auth.Name, auth.Login, auth.Password, lTime, auth.Name, uID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateText(ctx context.Context, data models.TextDataModel, uID int64) error {
	stmt, err := s.db.Prepare("UPDATE text_data SET name = ?, data = ?, last_update = ? WHERE name = ? and uId = ?")
	if err != nil {
		return err
	}
	lTime := time.Now().Format(time.RFC3339)
	_, err = stmt.ExecContext(ctx, data.Name, data.Data, lTime, data.Name, uID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateBin(ctx context.Context, data models.BinaryDataModel, uID int64) error {
	stmt, err := s.db.Prepare("UPDATE binares_data SET name = ?, data = ?, last_update = ? WHERE name = ? and uId = ?")
	if err != nil {
		return err
	}
	lTime := time.Now().Format(time.RFC3339)
	_, err = stmt.ExecContext(ctx, data.Name, data.Data, lTime, data.Name, uID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetAllSaves(ctx context.Context, uID int64) (models.SyncModel, error) {
	stmt, err := s.db.Prepare("SELECT name, data, uId, deleted, last_update FROM binares_data WHERE uId = ?")
	if err != nil {
		return models.SyncModel{}, err
	}
	binRows, err := stmt.QueryContext(ctx, uID)
	if err != nil {
		return models.SyncModel{}, err
	}
	defer binRows.Close()
	var binData []models.SyncBinaryDataModel
	for binRows.Next() {
		var data models.SyncBinaryDataModel
		err := binRows.Scan(&data.Name, &data.Data, &data.UserID, &data.Deleted, &data.Updated)
		if err != nil {
			return models.SyncModel{}, err
		}
		binData = append(binData, data)
	}

	stmt, err = s.db.Prepare("SELECT name, data, uId, deleted, last_update FROM text_data WHERE uId = ?")
	if err != nil {
		return models.SyncModel{}, err
	}
	textRows, err := stmt.QueryContext(ctx, uID)
	if err != nil {
		return models.SyncModel{}, err
	}
	defer textRows.Close()
	var textData []models.SyncTextDataModel
	for textRows.Next() {
		var data models.SyncTextDataModel
		err := textRows.Scan(&data.Name, &data.Data, &data.UserID, &data.Deleted, &data.Updated)
		if err != nil {
			return models.SyncModel{}, err
		}
		textData = append(textData, data)
	}

	stmt, err = s.db.Prepare("SELECT name, login, password, uId, deleted, last_update FROM logins WHERE uId = ?")
	if err != nil {
		return models.SyncModel{}, err
	}
	authRows, err := stmt.QueryContext(ctx, uID)
	if err != nil {
		return models.SyncModel{}, err
	}
	defer authRows.Close()
	var loginData []models.SyncLoginModel
	for authRows.Next() {
		var data models.SyncLoginModel
		err := authRows.Scan(&data.Name, &data.Login, &data.Password, &data.UserID, &data.Deleted, &data.Updated)
		if err != nil {
			return models.SyncModel{}, err
		}
		loginData = append(loginData, data)
	}

	stmt, err = s.db.Prepare("SELECT name, number, date, cvv, uId, deleted, last_update FROM cards WHERE uId = ?")
	if err != nil {
		return models.SyncModel{}, err
	}
	cardsRows, err := stmt.QueryContext(ctx, uID)
	if err != nil {
		return models.SyncModel{}, err
	}
	defer cardsRows.Close()
	var cardData []models.SyncCardModel
	for cardsRows.Next() {
		var data models.SyncCardModel
		err := cardsRows.Scan(&data.Name, &data.Number, &data.Date, &data.CVVCode, &data.UserID, &data.Deleted, &data.Updated)
		if err != nil {
			return models.SyncModel{}, err
		}
		cardData = append(cardData, data)
	}

	return models.SyncModel{
		Cards: cardData,
		Auth:  loginData,
		Texts: textData,
		Bins:  binData,
	}, err
}

func (s *Storage) ClearDB(ctx context.Context, uId int64) error {
	stmt, err := s.db.Prepare("DELETE FROM logins WHERE deleted = true AND uId = ?")
	if err != nil {
		return err
	}
	if _, err := stmt.ExecContext(ctx, uId); err != nil {
		return err
	}
	stmt, err = s.db.Prepare("DELETE FROM text_data WHERE deleted = true AND uId = $1")
	if err != nil {
		return err
	}
	if _, err := stmt.ExecContext(ctx, uId); err != nil {
		return err
	}
	stmt, err = s.db.Prepare("DELETE FROM binares_data WHERE deleted = true AND uId = $1")
	if err != nil {
		return err
	}
	if _, err := stmt.ExecContext(ctx, uId); err != nil {
		return err
	}
	stmt, err = s.db.Prepare("DELETE FROM cards WHERE deleted = true AND uId = $1")
	if err != nil {
		return err
	}
	if _, err := stmt.ExecContext(ctx, uId); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Sync(ctx context.Context, model models.SyncModel) error {
	stmt, err := s.db.Prepare(`INSERT INTO cards (name, number, date, cvv, uId, deleted, last_update) VALUES (?,?,?,?,?,?,?)
	ON CONFLICT (name) DO UPDATE SET name =?, number = ?, date = ?, cvv = ?, uId = ?, deleted = ?, last_update = ? WHERE name = ? and uId = ?`)
	if err != nil {
		return err
	}
	for _, card := range model.Cards {
		_, err := stmt.ExecContext(ctx, card.Name, card.Number, card.Date, card.CVVCode, card.UserID, card.Deleted, card.Updated,
			card.Name, card.Number, card.Date, card.CVVCode, card.UserID, card.Deleted, card.Updated, card.Name, card.UserID)
		if err != nil {
			return err
		}
	}

	stmt, err = s.db.Prepare(`INSERT INTO logins (name, login, password, uId, deleted, last_update) VALUES(?,?,?,?,?,?)
	ON CONFLICT (name) DO UPDATE SET name =?, login = ?, password = ?, uId = ?, deleted = ?, last_update = ? WHERE name = ? and uId = ?`)
	if err != nil {
		return err
	}
	for _, auth := range model.Auth {
		_, err := stmt.ExecContext(ctx, auth.Name, auth.Login, auth.Password, auth.UserID, auth.Deleted, auth.Updated,
			auth.Name, auth.Login, auth.Password, auth.UserID, auth.Deleted, auth.Updated, auth.Name, auth.UserID)
		if err != nil {
			return err
		}
	}

	for _, text := range model.Texts {
		stmt, err = s.db.Prepare(`INSERT INTO text_data(name, data, uId, deleted, last_update) VALUES(?,?,?,?,?)
		ON CONFLICT (name) DO UPDATE SET name = ?, data = ?, uId = ?, deleted = ?, last_update = ? WHERE name = ? and uId = ?`)
		if err != nil {
			return err
		}
		_, err := stmt.ExecContext(ctx, text.Name, text.Data, text.UserID, text.Deleted, text.Updated,
			text.Name, text.Data, text.UserID, text.Deleted, text.Updated, text.Name, text.UserID)
		if err != nil {
			return err
		}
	}

	stmt, err = s.db.Prepare(`INSERT INTO binares_data(name, data, uId, deleted, last_update) VALUES(?,?,?,?,?)
	ON CONFLICT (name) DO UPDATE SET name =?, data = ?, uId = ?, deleted = ?, last_update = ? WHERE name = ? and uId = ?`)
	if err != nil {
		return err
	}
	for _, bin := range model.Bins {
		_, err := stmt.ExecContext(ctx, bin.Name, bin.Data, bin.UserID, bin.Deleted, bin.Updated,
			bin.Name, bin.Data, bin.UserID, bin.Deleted, bin.Updated, bin.Name, bin.UserID)
		if err != nil {
			return err
		}
	}

	return nil
}
