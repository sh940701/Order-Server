package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuModel struct {
	Client *mongo.Client
	Menucollection *mongo.Collection
}

type Review struct {
	Score int `bson:"score" json:"score"`
	Review string `bson:"review" json:"review"`
}

type Menu struct {
	Id primitive.ObjectID `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
	Orderable bool `bson:"orderable" json:"orderable"`
	Limit int `bson:"limit" json:"limit"`
	Price int `bson:"price" json:"price"`
	From string `bson:"from" json:"from"`
	Orderedcount int `bson:"orderedcount" json:"orderedcount"`
	Reviews []Review `bson:"reviews" json:"reviews"`
}

func GetMenuModel(db, host, model string) (*MenuModel, error) {
	m := &MenuModel{}
	var err error

	if m.Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(host)); err != nil {
		return nil, err
	} else if err := m.Client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	} else {
		m.Menucollection = m.Client.Database(db).Collection(model)
	}
	return m, nil
}