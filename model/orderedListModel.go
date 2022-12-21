package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderedListModel struct {
	Client *mongo.Client
	Collection *mongo.Collection
}

type BuyerInfo struct {
	Pnum int `bson:"pnum" json:"pnum"`
	Address string `bson:"address" json:"address"`
}

type OrderedList struct {
	Id primitive.ObjectID `bson:"id" json:"id"`
	IsReviewed bool `bson:"isreviewed" json:"isreviewed"`
	Status string `bson:"status" json:"status"`
	BuyerInfo BuyerInfo `bson:"buyerinfo" json:"buyerinfo"`
	Orderedmenus []primitive.ObjectID `bson:"orderedmenus"`
}

func GetOrderedListModel(host, db, model string) (*OrderedListModel, error) {
	m := &OrderedListModel{}
	var err error
	
	if m.Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(host)); err != nil {
		return nil, err
	} else if err := m.Client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	} else {
		m.Collection = m.Client.Database(db).Collection(model)
	}
	return m, nil
}