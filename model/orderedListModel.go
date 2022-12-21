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

func GetOrderedListModel(db, host, model string) (*OrderedListModel, error) {
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

// 주문내역 전체 리스트 가져오는 메서드
func (o *OrderedListModel) GetAll() {

}

// 주문내역 하나만 가져오는 메서드
func (o *OrderedListModel) GetOne() {

}

// 요청받은 주문의 진행상태 업데이트 하는 메서드
func (o *OrderedListModel) UpdateStatus() {

}

// 주문내역의 리뷰 작성 여부 업데이트하는 메서드
func (o *OrderedListModel) UpdateReviewable() {

}

// 주문을 추가하는 메서드
func (o *OrderedListModel) Add() {

}

// 하루치 주문 번호 업데이트 하는 메서드
func (o *OrderedListModel) UpdateDayOrderCount() {

}

// 주문 메뉴 변경하기
func (o *OrderedListModel) ChangeOrder() {

}

// 주문 메뉴 추가하기
func (o *OrderedListModel) AddOrder() {

}