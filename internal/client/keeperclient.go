package client

import (
	"context"
	"strconv"

	"github.com/Dorrrke/GophKeeper-client/internal/domain/models"
	gophkeeperv1 "github.com/Dorrrke/goph-keeper-proto/gen/go/gophkeeper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type KeeperClient struct {
	client gophkeeperv1.GophKeeperClient
	conn   *grpc.ClientConn
	token  string
}

func New(ctx context.Context, addr string) (*KeeperClient, error) {
	conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := gophkeeperv1.NewGophKeeperClient(conn)
	return &KeeperClient{
		client: client,
		conn:   conn,
	}, nil
}

func (c *KeeperClient) Register(ctx context.Context, login, password string) error {
	var header metadata.MD
	_, err := c.client.SignUp(ctx, &gophkeeperv1.SignUpRequest{
		Login:    login,
		Password: password,
	}, grpc.Header(&header))
	if err != nil {
		return err
	}
	tokens := header.Get("Authorization")
	c.token = tokens[0]
	return nil
}

func (c *KeeperClient) Login(ctx context.Context, login, password string) error {
	var header metadata.MD
	_, err := c.client.SignIn(ctx, &gophkeeperv1.SingInRequest{
		Login:    login,
		Password: password,
	}, grpc.Header(&header))
	if err != nil {
		return err
	}
	c.token = header.Get("Authorization")[0]
	return nil
}

func (c *KeeperClient) Sync(ctx context.Context, model models.SyncModel, uID int64) (models.SyncModel, error) {
	protoModel := modelToProtoModel(model)
	md := metadata.Pairs("Authorization", c.token)
	mCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := c.client.SyncDB(mCtx, &gophkeeperv1.SyncDBRequest{
		Auth:  protoModel.Auth,
		Bins:  protoModel.Bins,
		Cards: protoModel.Cards,
		Texts: protoModel.Texts,
	})
	if err != nil {
		return models.SyncModel{}, err
	}

	resModel, err := protoModelToModel(models.ProtoSyncModel{
		Cards: res.Cards,
		Auth:  res.Auth,
		Texts: res.Texts,
		Bins:  res.Bins,
	}, uID)
	if err != nil {
		return models.SyncModel{}, err
	}
	return resModel, nil
}

func modelToProtoModel(model models.SyncModel) models.ProtoSyncModel {
	var pModel models.ProtoSyncModel
	for _, data := range model.Bins {
		bin := &gophkeeperv1.SyncBinData{
			Name:    data.Name,
			Data:    data.Data,
			Deleted: data.Deleted,
			Updated: data.Updated,
		}
		pModel.Bins = append(pModel.Bins, bin)
	}
	for _, data := range model.Auth {
		auth := &gophkeeperv1.SyncAuth{
			Name:     data.Name,
			Login:    data.Login,
			Password: data.Password,
			Deleted:  data.Deleted,
			Updated:  data.Updated,
		}
		pModel.Auth = append(pModel.Auth, auth)
	}
	for _, data := range model.Cards {
		card := &gophkeeperv1.SyncCard{
			Name:    data.Name,
			Number:  data.Number,
			Date:    data.Date,
			Cvv:     strconv.Itoa(data.CVVCode),
			Deleted: data.Deleted,
			Updated: data.Updated,
		}
		pModel.Cards = append(pModel.Cards, card)
	}
	for _, data := range model.Texts {
		text := &gophkeeperv1.SyncText{
			Name:    data.Name,
			Data:    data.Data,
			Deleted: data.Deleted,
			Updated: data.Updated,
		}
		pModel.Texts = append(pModel.Texts, text)
	}

	return pModel
}

func protoModelToModel(model models.ProtoSyncModel, uID int64) (models.SyncModel, error) {
	var sModel models.SyncModel
	for _, data := range model.Bins {
		bin := models.SyncBinaryDataModel{
			UserID:  uID,
			Name:    data.Name,
			Data:    data.Data,
			Deleted: data.Deleted,
			Updated: data.Updated,
		}
		sModel.Bins = append(sModel.Bins, bin)
	}
	for _, data := range model.Auth {
		auth := models.SyncLoginModel{
			UserID:   uID,
			Name:     data.Name,
			Login:    data.Login,
			Password: data.Password,
			Deleted:  data.Deleted,
			Updated:  data.Updated,
		}
		sModel.Auth = append(sModel.Auth, auth)
	}
	for _, data := range model.Cards {
		cvv, err := strconv.Atoi(data.Cvv)
		if err != nil {
			return models.SyncModel{}, err
		}
		card := models.SyncCardModel{
			UserID:  uID,
			Name:    data.Name,
			Number:  data.Number,
			Date:    data.Date,
			CVVCode: cvv,
			Deleted: data.Deleted,
			Updated: data.Updated,
		}
		sModel.Cards = append(sModel.Cards, card)
	}
	for _, data := range model.Texts {
		text := models.SyncTextDataModel{
			UserID:  uID,
			Name:    data.Name,
			Data:    data.Data,
			Deleted: data.Deleted,
			Updated: data.Updated,
		}
		sModel.Texts = append(sModel.Texts, text)
	}

	return sModel, nil
}
