package model

import (
	"fmt"
	"github.com/codestates/WBABEProject-08/commits/main/util"
	"encoding/json"
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
	Score int `bson:"score,omitemepty" json:"score"`
	Review string `bson:"review,omitemepty" json:"review"`
}

type Menu struct {
	_Id primitive.ObjectID `bson:"id"`
	Name string `bson:"name" json:"name"`
	Orderable bool `bson:"orderable" json:"orderable"`
	Limit int `bson:"limit" json:"limit"`
	Price int `bson:"price" json:"price"`
	From string `bson:"from" json:"from"`
	Orderedcount int `bson:"orderedcount" json:"orderedcount"`
	Reviews []Review `bson:"reviews,omitemepty" json:"reviews"`
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


// DB에 메뉴 data를 추가하는 메서드
func (m *MenuModel) Add(data []byte) interface{} {
	newMenu := &Menu{}
	json.Unmarshal(data, newMenu)

	fmt.Println("Unmarshar: ", newMenu)

	result, err := m.Menucollection.InsertOne(context.TODO(), newMenu)
	util.PanicHandler(err)

	return result.InsertedID
}

// DB 메뉴 data를 업데이트하는 메서드
func (m *MenuModel) Update() {

}

// DB 메뉴 data를 삭제하는 메서드
func (m *MenuModel) Delete() {

}

// 메뉴 리스트를 조회하는 메서드
func (m *MenuModel) GetList() {

}

// 메뉴별 평점/리뷰 데이터 조회하는 메서드
func (m *MenuModel) GetReview() {

}

// 메뉴별 평점/리뷰 데이터 추가하는 메서드
func (m *MenuModel) AddReview() {
	
}