package services

import (
	"context"
	"errors"
	"testing"

	"github.com/Dorrrke/GophKeeper-client/internal/client"
	"github.com/Dorrrke/GophKeeper-client/internal/domain/models"
	"github.com/golang/mock/gomock"
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
			m := NewMockStorage(ctrl)
			m.EXPECT().SaveCard(context.Background(), tc.card, tc.uID).Return(tc.want.uID, tc.want.err)
			service := New(&client.KeeperClient{}, m)
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
			m := NewMockStorage(ctrl)
			m.EXPECT().SaveLogin(context.Background(), tc.login, tc.uID).Return(tc.want.uID, tc.want.err)
			service := New(&client.KeeperClient{}, m)
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
			m := NewMockStorage(ctrl)
			m.EXPECT().SaveText(context.Background(), tc.text, tc.uID).Return(tc.want.uID, tc.want.err)
			service := New(&client.KeeperClient{}, m)
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
			m := NewMockStorage(ctrl)
			m.EXPECT().SaveBin(context.Background(), tc.text, tc.uID).Return(tc.want.uID, tc.want.err)
			service := New(&client.KeeperClient{}, m)
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
			m := NewMockStorage(ctrl)
			m.EXPECT().GetAllBin(context.Background(), tc.uID).Return(tc.want.bins, tc.want.err)
			service := New(&client.KeeperClient{}, m)
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
			m := NewMockStorage(ctrl)
			m.EXPECT().GetAllCards(context.Background(), tc.uID).Return(tc.want.cards, tc.want.err)
			service := New(&client.KeeperClient{}, m)
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
			m := NewMockStorage(ctrl)
			m.EXPECT().GetAllLogins(context.Background(), tc.uID).Return(tc.want.logins, tc.want.err)
			service := New(&client.KeeperClient{}, m)
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
			m := NewMockStorage(ctrl)
			m.EXPECT().GetAllTextData(context.Background(), tc.uID).Return(tc.want.text, tc.want.err)
			service := New(&client.KeeperClient{}, m)
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
			m := NewMockStorage(ctrl)
			m.EXPECT().GetCardByName(context.Background(), tc.cName, tc.uID).Return(tc.want.data, tc.want.err)
			service := New(&client.KeeperClient{}, m)
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
			m := NewMockStorage(ctrl)
			m.EXPECT().GetLoginByName(context.Background(), tc.cName, tc.uID).Return(tc.want.data, tc.want.err)
			service := New(&client.KeeperClient{}, m)
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
			m := NewMockStorage(ctrl)
			m.EXPECT().GetTextDataByName(context.Background(), tc.cName, tc.uID).Return(tc.want.data, tc.want.err)
			service := New(&client.KeeperClient{}, m)
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
			m := NewMockStorage(ctrl)
			m.EXPECT().GetBinByName(context.Background(), tc.cName, tc.uID).Return(tc.want.data, tc.want.err)
			service := New(&client.KeeperClient{}, m)
			res, _ := service.GetBinByName(tc.cName, tc.uID)
			assert.Equal(t, tc.want.data, res)
		})
	}
}
