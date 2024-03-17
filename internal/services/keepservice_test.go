package services

import (
	"context"
	"errors"
	"testing"

	"github.com/Dorrrke/GophKeeper-client/internal/client"
	"github.com/Dorrrke/GophKeeper-client/internal/domain/models"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang/mock/gomock"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestSaveCard(t *testing.T) {
	type want struct {
		uID int64
		err error
	}
	type test struct {
		name string
		card models.CardModel
		uID  int64
		want want
	}
	tests := []test{
		{
			name: "Test SaveCard function #1; Default call",
			card: models.CardModel{
				Name:    "Card1",
				Number:  "132514562235",
				Date:    "09/33",
				CVVCode: 844,
			},
			uID: 1,
			want: want{
				uID: 4,
				err: nil,
			},
		},
		{
			name: "Test SaveCard function #2; Error call",
			card: models.CardModel{
				Name:    "Card12",
				Number:  "13251454362235",
				Date:    "09/44",
				CVVCode: 854,
			},
			uID: 4,
			want: want{
				uID: -1,
				err: errors.New("Conflict"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m1 := NewMockCardStorage(ctrl)
			m2 := NewMockAuthStorage(ctrl)
			m3 := NewMockBinStorage(ctrl)
			m4 := NewMockStorage(ctrl)
			m5 := NewMockTextStorage(ctrl)
			m6 := NewMockUserStorage(ctrl)
			m1.EXPECT().SaveCard(context.Background(), tc.card, tc.uID).Return(tc.want.uID, tc.want.err)
			service := New(&client.KeeperClient{}, m4, m6, m1, m5, m3, m2)
			res, _ := service.SaveCard(tc.card, tc.uID)
			assert.Equal(t, tc.want.uID, res)
		})
	}
}

func TestSaveLogin(t *testing.T) {
	type want struct {
		uID int64
		err error
	}
	type test struct {
		name  string
		login models.LoginModel
		uID   int64
		want  want
	}
	tests := []test{
		{
			name: "Test SaveLogin function #1; Default call",
			login: models.LoginModel{
				Name:     "Card1",
				Login:    "login1",
				Password: "Pass1",
			},
			uID: 1,
			want: want{
				uID: 4,
				err: nil,
			},
		},
		{
			name: "Test SaveLogin function #2; Error call",
			login: models.LoginModel{
				Name:     "Card1",
				Login:    "login12",
				Password: "Pass13",
			},
			uID: 4,
			want: want{
				uID: -1,
				err: errors.New("Conflict"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m1 := NewMockCardStorage(ctrl)
			m2 := NewMockAuthStorage(ctrl)
			m3 := NewMockBinStorage(ctrl)
			m4 := NewMockStorage(ctrl)
			m5 := NewMockTextStorage(ctrl)
			m6 := NewMockUserStorage(ctrl)
			m2.EXPECT().SaveLogin(context.Background(), tc.login, tc.uID).Return(tc.want.uID, tc.want.err)
			service := New(&client.KeeperClient{}, m4, m6, m1, m5, m3, m2)
			res, _ := service.SaveLogin(tc.login, tc.uID)
			assert.Equal(t, tc.want.uID, res)
		})
	}
}

func TestSaveTextData(t *testing.T) {
	type want struct {
		uID int64
		err error
	}
	type test struct {
		name string
		text models.TextDataModel
		uID  int64
		want want
	}
	tests := []test{
		{
			name: "Test SaveTextData function #1; Default call",
			text: models.TextDataModel{
				Name: "Note 1",
				Data: "Buy notebook lalalalla",
			},
			uID: 1,
			want: want{
				uID: 4,
				err: nil,
			},
		},
		{
			name: "Test SaveTextData function #2; Error call",
			text: models.TextDataModel{
				Name: "Note 2",
				Data: "lalalalla",
			},
			uID: 4,
			want: want{
				uID: -1,
				err: errors.New("Conflict"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m1 := NewMockCardStorage(ctrl)
			m2 := NewMockAuthStorage(ctrl)
			m3 := NewMockBinStorage(ctrl)
			m4 := NewMockStorage(ctrl)
			m5 := NewMockTextStorage(ctrl)
			m6 := NewMockUserStorage(ctrl)
			m5.EXPECT().SaveText(context.Background(), tc.text, tc.uID).Return(tc.want.uID, tc.want.err)
			service := New(&client.KeeperClient{}, m4, m6, m1, m5, m3, m2)
			res, _ := service.SaveTextData(tc.text, tc.uID)
			assert.Equal(t, tc.want.uID, res)
		})
	}
}

func TestSaveBinaryData(t *testing.T) {
	type want struct {
		uID int64
		err error
	}
	type test struct {
		name string
		text models.BinaryDataModel
		uID  int64
		want want
	}
	tests := []test{
		{
			name: "Test SaveBinaryData function #1; Default call",
			text: models.BinaryDataModel{
				Name: "Note 1",
				Data: []byte("Buy notebook lalalalla"),
			},
			uID: 1,
			want: want{
				uID: 4,
				err: nil,
			},
		},
		{
			name: "Test SaveBinaryData function #2; Error call",
			text: models.BinaryDataModel{
				Name: "Note 2",
				Data: []byte("lalalalla"),
			},
			uID: 4,
			want: want{
				uID: -1,
				err: errors.New("Conflict"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m1 := NewMockCardStorage(ctrl)
			m2 := NewMockAuthStorage(ctrl)
			m3 := NewMockBinStorage(ctrl)
			m4 := NewMockStorage(ctrl)
			m5 := NewMockTextStorage(ctrl)
			m6 := NewMockUserStorage(ctrl)
			m3.EXPECT().SaveBin(context.Background(), tc.text, tc.uID).Return(tc.want.uID, tc.want.err)
			service := New(&client.KeeperClient{}, m4, m6, m1, m5, m3, m2)
			res, _ := service.SaveBinaryData(tc.text, tc.uID)
			assert.Equal(t, tc.want.uID, res)
		})
	}
}

func TestGetBins(t *testing.T) {
	type want struct {
		bins []models.BinaryDataModel
		err  error
	}
	type test struct {
		name string
		uID  int64
		want want
	}
	tests := []test{
		{
			name: "Test GetBins function #1; Default call",
			uID:  1,
			want: want{
				bins: []models.BinaryDataModel{
					{
						Name: "Bin1",
						Data: []byte("dafs123gas53dgd"),
					},
					{
						Name: "Bin2",
						Data: []byte("gad345sfg"),
					},
					{
						Name: "Bin3",
						Data: []byte("gasd3215"),
					},
				},
				err: nil,
			},
		},
		{
			name: "Test GetBins function #2; Error call",
			uID:  4,
			want: want{
				bins: nil,
				err:  errors.New("Conflict"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m1 := NewMockCardStorage(ctrl)
			m2 := NewMockAuthStorage(ctrl)
			m3 := NewMockBinStorage(ctrl)
			m4 := NewMockStorage(ctrl)
			m5 := NewMockTextStorage(ctrl)
			m6 := NewMockUserStorage(ctrl)
			m3.EXPECT().GetAllBin(context.Background(), tc.uID).Return(tc.want.bins, tc.want.err)
			service := New(&client.KeeperClient{}, m4, m6, m1, m5, m3, m2)
			res, _ := service.GetBins(tc.uID)
			assert.Equal(t, tc.want.bins, res)
		})
	}
}

func TestGetCards(t *testing.T) {
	type want struct {
		cards []models.CardModel
		err   error
	}
	type test struct {
		name string
		uID  int64
		want want
	}
	tests := []test{
		{
			name: "Test GetBins function #1; Default call",
			uID:  1,
			want: want{
				cards: []models.CardModel{
					{
						Name:    "Card123",
						Number:  "1325145434362235",
						Date:    "09/34",
						CVVCode: 854,
					},
					{
						Name:    "Card125",
						Number:  "13251464754362235",
						Date:    "12/44",
						CVVCode: 854,
					},
					{
						Name:    "Card126",
						Number:  "1325143454362235",
						Date:    "05/34",
						CVVCode: 854,
					},
				},
				err: nil,
			},
		},
		{
			name: "Test GetBins function #2; Error call",
			uID:  4,
			want: want{
				cards: nil,
				err:   errors.New("Conflict"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m1 := NewMockCardStorage(ctrl)
			m2 := NewMockAuthStorage(ctrl)
			m3 := NewMockBinStorage(ctrl)
			m4 := NewMockStorage(ctrl)
			m5 := NewMockTextStorage(ctrl)
			m6 := NewMockUserStorage(ctrl)
			m1.EXPECT().GetAllCards(context.Background(), tc.uID).Return(tc.want.cards, tc.want.err)
			service := New(&client.KeeperClient{}, m4, m6, m1, m5, m3, m2)
			res, _ := service.GetCards(tc.uID)
			assert.Equal(t, tc.want.cards, res)
		})
	}
}

func TestGetLogins(t *testing.T) {
	type want struct {
		logins []models.LoginModel
		err    error
	}
	type test struct {
		name string
		uID  int64
		want want
	}
	tests := []test{
		{
			name: "Test GetBins function #1; Default call",
			uID:  1,
			want: want{
				logins: []models.LoginModel{
					{
						Name:     "C1",
						Login:    "login1",
						Password: "Pass1",
					},
					{
						Name:     "d1",
						Login:    "lo89gin1",
						Password: "Pa44ss1",
					},
					{
						Name:     "rd1",
						Login:    "loghhin1",
						Password: "Palkss1",
					},
				},
				err: nil,
			},
		},
		{
			name: "Test GetBins function #2; Error call",
			uID:  4,
			want: want{
				logins: nil,
				err:    errors.New("Conflict"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m1 := NewMockCardStorage(ctrl)
			m2 := NewMockAuthStorage(ctrl)
			m3 := NewMockBinStorage(ctrl)
			m4 := NewMockStorage(ctrl)
			m5 := NewMockTextStorage(ctrl)
			m6 := NewMockUserStorage(ctrl)
			m2.EXPECT().GetAllLogins(context.Background(), tc.uID).Return(tc.want.logins, tc.want.err)
			service := New(&client.KeeperClient{}, m4, m6, m1, m5, m3, m2)
			res, _ := service.GetLogins(tc.uID)
			assert.Equal(t, tc.want.logins, res)
		})
	}
}

func TestGetTextData(t *testing.T) {
	type want struct {
		text []models.TextDataModel
		err  error
	}
	type test struct {
		name string
		uID  int64
		want want
	}
	tests := []test{
		{
			name: "Test GetBins function #1; Default call",
			uID:  1,
			want: want{
				text: []models.TextDataModel{
					{
						Name: "text1",
						Data: "dafs123gas53dgd",
					},
					{
						Name: "text2",
						Data: "gad345sfg",
					},
					{
						Name: "text3",
						Data: "gasd3215",
					},
				},
				err: nil,
			},
		},
		{
			name: "Test GetBins function #2; Error call",
			uID:  4,
			want: want{
				text: nil,
				err:  errors.New("Conflict"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m1 := NewMockCardStorage(ctrl)
			m2 := NewMockAuthStorage(ctrl)
			m3 := NewMockBinStorage(ctrl)
			m4 := NewMockStorage(ctrl)
			m5 := NewMockTextStorage(ctrl)
			m6 := NewMockUserStorage(ctrl)
			m5.EXPECT().GetAllTextData(context.Background(), tc.uID).Return(tc.want.text, tc.want.err)
			service := New(&client.KeeperClient{}, m4, m6, m1, m5, m3, m2)
			res, _ := service.GetTextData(tc.uID)
			assert.Equal(t, tc.want.text, res)
		})
	}
}

func TestGetCardByName(t *testing.T) {
	type want struct {
		data models.CardModel
		err  error
	}
	type test struct {
		name  string
		cName string
		uID   int64
		want  want
	}
	tests := []test{
		{
			name: "Test GetBins function #1; Default call",
			uID:  1,
			want: want{
				data: models.CardModel{
					Name:    "Card125",
					Number:  "13251464754362235",
					Date:    "12/44",
					CVVCode: 854,
				},
				err: nil,
			},
		},
		{
			name: "Test GetBins function #2; Error call",
			uID:  4,
			want: want{
				data: models.CardModel{},
				err:  errors.New("Conflict"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m1 := NewMockCardStorage(ctrl)
			m2 := NewMockAuthStorage(ctrl)
			m3 := NewMockBinStorage(ctrl)
			m4 := NewMockStorage(ctrl)
			m5 := NewMockTextStorage(ctrl)
			m6 := NewMockUserStorage(ctrl)
			m1.EXPECT().GetCardByName(context.Background(), tc.cName, tc.uID).Return(tc.want.data, tc.want.err)
			service := New(&client.KeeperClient{}, m4, m6, m1, m5, m3, m2)
			res, _ := service.GetCardByName(tc.cName, tc.uID)
			assert.Equal(t, tc.want.data, res)
		})
	}
}

func GetLoginByName(t *testing.T) {
	type want struct {
		data models.LoginModel
		err  error
	}
	type test struct {
		name  string
		cName string
		uID   int64
		want  want
	}
	tests := []test{
		{
			name: "Test GetBins function #1; Default call",
			uID:  1,
			want: want{
				data: models.LoginModel{
					Name:     "Card125",
					Login:    "log1",
					Password: "Passw1",
				},
				err: nil,
			},
		},
		{
			name: "Test GetBins function #2; Error call",
			uID:  4,
			want: want{
				data: models.LoginModel{},
				err:  errors.New("Conflict"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m1 := NewMockCardStorage(ctrl)
			m2 := NewMockAuthStorage(ctrl)
			m3 := NewMockBinStorage(ctrl)
			m4 := NewMockStorage(ctrl)
			m5 := NewMockTextStorage(ctrl)
			m6 := NewMockUserStorage(ctrl)
			m2.EXPECT().GetLoginByName(context.Background(), tc.cName, tc.uID).Return(tc.want.data, tc.want.err)
			service := New(&client.KeeperClient{}, m4, m6, m1, m5, m3, m2)
			res, _ := service.GetLoginByName(tc.cName, tc.uID)
			assert.Equal(t, tc.want.data, res)
		})
	}
}

func TesGetTextDataByName(t *testing.T) {
	type want struct {
		data models.TextDataModel
		err  error
	}
	type test struct {
		name  string
		cName string
		uID   int64
		want  want
	}
	tests := []test{
		{
			name: "Test GetBins function #1; Default call",
			uID:  1,
			want: want{
				data: models.TextDataModel{
					Name: "Card125",
					Data: "gfsgdfgsdg",
				},
				err: nil,
			},
		},
		{
			name: "Test GetBins function #2; Error call",
			uID:  4,
			want: want{
				data: models.TextDataModel{},
				err:  errors.New("Conflict"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m1 := NewMockCardStorage(ctrl)
			m2 := NewMockAuthStorage(ctrl)
			m3 := NewMockBinStorage(ctrl)
			m4 := NewMockStorage(ctrl)
			m5 := NewMockTextStorage(ctrl)
			m6 := NewMockUserStorage(ctrl)
			m5.EXPECT().GetTextDataByName(context.Background(), tc.cName, tc.uID).Return(tc.want.data, tc.want.err)
			service := New(&client.KeeperClient{}, m4, m6, m1, m5, m3, m2)
			res, _ := service.GetTextDataByName(tc.cName, tc.uID)
			assert.Equal(t, tc.want.data, res)
		})
	}
}

func TestGetBinByName(t *testing.T) {
	type want struct {
		data models.BinaryDataModel
		err  error
	}
	type test struct {
		name  string
		cName string
		uID   int64
		want  want
	}
	tests := []test{
		{
			name: "Test GetBins function #1; Default call",
			uID:  1,
			want: want{
				data: models.BinaryDataModel{
					Name: "Card125",
					Data: []byte("fdfdf"),
				},
				err: nil,
			},
		},
		{
			name: "Test GetBins function #2; Error call",
			uID:  4,
			want: want{
				data: models.BinaryDataModel{},
				err:  errors.New("Conflict"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m1 := NewMockCardStorage(ctrl)
			m2 := NewMockAuthStorage(ctrl)
			m3 := NewMockBinStorage(ctrl)
			m4 := NewMockStorage(ctrl)
			m5 := NewMockTextStorage(ctrl)
			m6 := NewMockUserStorage(ctrl)
			m3.EXPECT().GetBinByName(context.Background(), tc.cName, tc.uID).Return(tc.want.data, tc.want.err)
			service := New(&client.KeeperClient{}, m4, m6, m1, m5, m3, m2)
			res, _ := service.GetBinByName(tc.cName, tc.uID)
			assert.Equal(t, tc.want.data, res)
		})
	}
}

// func TestSyncBD(t *testing.T) {
// 	type want struct {
// 		data models.SyncModel
// 	}
// 	type test struct {
// 		name  string
// 		card  models.CardModel
// 		sCard models.SyncCardModel
// 		uID   int64
// 		want  want
// 	}
// 	tests := []test{
// 		{
// 			name: "TestSyncBD 1",
// 			card: models.CardModel{
// 				Name:    "Test card",
// 				Number:  "8899223344556677",
// 				Date:    "08/29",
// 				CVVCode: 654,
// 			},
// 			sCard: models.SyncCardModel{
// 				UserID:  1,
// 				Name:    "Test card",
// 				Number:  "8899223344556677",
// 				Date:    "08/29",
// 				CVVCode: 654,
// 				Deleted: false,
// 				Updated: time.Now().Format(time.RFC3339),
// 			},
// 			uID: 1,
// 			want: want{
// 				data: models.SyncModel{
// 					Cards: []models.SyncCardModel{},
// 					Texts: []models.SyncTextDataModel{},
// 					Bins:  []models.SyncBinaryDataModel{},
// 					Auth:  []models.SyncLoginModel{},
// 				},
// 			},
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			err := MigrDown()
// 			assert.NoError(t, err)
// 			sModel := models.SyncModel{
// 				Cards: []models.SyncCardModel{},
// 				Texts: []models.SyncTextDataModel{},
// 				Bins:  []models.SyncBinaryDataModel{},
// 				Auth:  []models.SyncLoginModel{},
// 			}
// 			sModel.Cards = append(sModel.Cards, tc.sCard)
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()
// 			err = MigrUp()
// 			assert.NoError(t, err)
// 			storage, err := storage.New("testdb.db")
// 			assert.NoError(t, err)
// 			m := NewMockClient(ctrl)
// 			m.EXPECT().Sync(context.Background(), sModel, tc.uID).Return(models.SyncModel{
// 				Cards: []models.SyncCardModel{},
// 				Texts: []models.SyncTextDataModel{},
// 				Bins:  []models.SyncBinaryDataModel{},
// 				Auth:  []models.SyncLoginModel{},
// 			}, nil)
// 			service := New(m, storage, storage, storage, storage, storage, storage)
// 			_, err = service.SaveCard(tc.card, tc.uID)
// 			assert.NoError(t, err)
// 			err = service.SyncBD(context.Background(), tc.uID)
// 			assert.NoError(t, err)
// 		})
// 	}
// }

// func MigrUp() error {
// 	sPath := "testdb.db"
// 	mPath := "migrations/"
// 	migratePath := "file://" + mPath
// 	storPath := fmt.Sprintf("sqlite3://%s", sPath)
// 	m, err := migrate.New(
// 		migratePath,
// 		storPath,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	if err := m.Up(); err != nil {
// 		if errors.Is(err, migrate.ErrNoChange) {
// 			return nil
// 		}
// 		return err
// 	}
// 	return nil
// }

// func MigrDown() error {
// 	sPath := "testdb.db"
// 	mPath := "migrations/"
// 	migratePath := "file://" + mPath
// 	storPath := fmt.Sprintf("sqlite3://%s", sPath)
// 	m, err := migrate.New(
// 		migratePath,
// 		storPath,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	if err := m.Down(); err != nil {
// 		if errors.Is(err, migrate.ErrNoChange) {
// 			return nil
// 		}
// 		return err
// 	}
// 	return nil
// }
