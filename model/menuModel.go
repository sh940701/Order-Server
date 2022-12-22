package model

import (
	"go.mongodb.org/mongo-driver/bson"
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
	ID primitive.ObjectID `bson:"_id,omitempty"`
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
func (m *MenuModel) Add(data []byte) (primitive.ObjectID, error) {
	newMenu := &Menu{}
	json.Unmarshal(data, newMenu)

	fmt.Println("Unmarshar: ", newMenu)

	result, err := m.Menucollection.InsertOne(context.TODO(), newMenu)
	if err != nil {
		return result.InsertedID.(primitive.ObjectID), err
	} else {
		return result.InsertedID.(primitive.ObjectID), nil
	}
}

// DB 메뉴 data를 업데이트하는 메서드
func (m *MenuModel) Update(data []byte) (interface{}, error) {
	id, key, value := util.GetJsonIdKeyValue(data)
	
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: key, Value: value}}}}

	result, err := m.Menucollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return result.MatchedCount, err
	} else {
		return result.MatchedCount, nil
	}
}

// DB 메뉴 data를 삭제하는 메서드
func (m *MenuModel) Delete(data []byte) (interface{}, error) {
	id, _, _ := util.GetJsonIdKeyValue(data)
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := m.Menucollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, err
	} else {
		return result.DeletedCount, nil
	}
}

// 메뉴 리스트를 조회하는 메서드
func (m *MenuModel) GetList(category string) []Menu {
	menus := []Menu{}
	filter := bson.D{}

	opt := options.Find().SetSort(bson.D{{Key: category, Value: -1}})
	cursor, err := m.Menucollection.Find(context.TODO(), filter, opt)
	util.PanicHandler(err)
	err = cursor.All(context.TODO(), &menus)
	util.PanicHandler(err)

	return menus
}

// 메뉴별 평점/리뷰 데이터 조회하는 메서드
func (m *MenuModel) GetReview() {

}

// 메뉴별 평점/리뷰 데이터 추가하는 메서드
func (m *MenuModel) AddReview() {
	
}

// 메뉴의 limit을 1 줄여주고 count를 1 올려주는 함수
func (m *MenuModel) LimitAndCountUpdate(id primitive.ObjectID, limit, count int) {
	var update bson.D
	filter := bson.D{{Key: "_id", Value: id}}
	if limit != 1 {
		update = bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "limit", Value : limit - 1}, 
				{Key: "orderedcount", Value: count + 1}}}}
	} else {
		update = bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "limit", Value : limit - 1}, 
				{Key: "orderedcount", Value: count + 1}, 
				{Key: "orderable", Value: false}}}}
	}
	_, err := m.Menucollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		util.PanicHandler(err)
	}
}


// 메뉴를 하나만 찾아주는 함수
func (m *MenuModel) GetOneMenu(id primitive.ObjectID) (*Menu, error) {
	menu := &Menu{}
	filter := bson.D{{Key: "_id", Value: id}}
	err := m.Menucollection.FindOne(context.TODO(), filter).Decode(menu)
	if err != nil {
		return nil, err
	} else {
		return menu, nil
	}
}