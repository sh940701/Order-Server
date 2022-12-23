package model

import (
	"github.com/codestates/WBABEProject-08/commits/main/util"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"fmt"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderedListModel struct {
	Client *mongo.Client
	Collection *mongo.Collection
}

type BuyerInfo struct {
	Pnum string `bson:"pnum" json:"pnum"`
	Address string `bson:"address" json:"address"`
}

type OrderedList struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	IsReviewed bool `bson:"isreviewed" json:"isreviewed"`
	Status string `bson:"status" json:"status"`
	BuyerInfo BuyerInfo `bson:"buyerinfo" json:"buyerinfo"`
	Orderedmenus []primitive.ObjectID `bson:"orderedmenus"`
}

type DayCounter struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Count int `bson:"daycount"`
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
func (o *OrderedListModel) GetAll(exceptId string) *[]OrderedList {
	// 주문내역 리스트에는 daycount document도 있으므로, 이를 제외하고 가져온다.
	exId, err := primitive.ObjectIDFromHex(exceptId)
	util.PanicHandler(err)

	list := []OrderedList{}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: exId}}}}
	cursor, err := o.Collection.Find(context.TODO(), filter)
	util.PanicHandler(err)
	cursor.All(context.TODO(), &list)

	return &list
}


// 주문내역 하나만 가져오는 메서드
func (o *OrderedListModel) GetOne(id primitive.ObjectID) *OrderedList {
	list := &OrderedList{}
	
	filter := bson.D{{Key: "_id", Value: id}}
	o.Collection.FindOne(context.TODO(), filter).Decode(list)

	return list
}


// 주문의 진행상태 업데이트 하는 메서드
func (o *OrderedListModel) UpdateStatus(order *OrderedList) string {
	// 배달 완료된 상태이면 return
	if order.Status == "배달완료" {
		return "이미 완료된 주문입니다."
	}

	// 현재 상태에 따른 업데이트 진행
	var status string
	if order.Status == "주문접수" {
		status = "조리중"
	} else if order.Status == "조리중" {
		status = "배달중"
	} else if order.Status == "배달중" {
		status = "배달완료"
	} 
	filter := bson.D{{Key: "_id", Value: order.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: status}}}}
	_, err := o.Collection.UpdateOne(context.TODO(), filter, update)
	util.PanicHandler(err)

	s := fmt.Sprintf("상태가 %s으로 변경되었습니다.", status)
	return s
}


// 주문내역의 리뷰 작성 여부 업데이트하는 메서드
func (o *OrderedListModel) UpdateReviewable(id primitive.ObjectID) {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "isreviewed", Value: true}}}}
	_, err := o.Collection.UpdateOne(context.TODO(), filter, update)
	util.PanicHandler(err)
}


// 주문을 추가하는 메서드
func (o *OrderedListModel) Add(list *OrderedList) primitive.ObjectID {
	result, err := o.Collection.InsertOne(context.TODO(), list)
	if err != nil {
		util.PanicHandler(err)
	}
	return result.InsertedID.(primitive.ObjectID)
}


// 주문 메뉴 변경하기
func (o *OrderedListModel) ChangeOrder(order *OrderedList, change *ChangeMenuType) error {
	isChanged := false
	// 먼저 변경하고싶은 메뉴가 orderlist에 있는지 확인한다.
	for idx, value := range order.Orderedmenus {
		if value == change.LegacyFoodId {
			order.Orderedmenus[idx] = change.NewFoodId
			isChanged = true
		}
	}
	if !isChanged {
		return errors.New("주문 내역에 해당 메뉴가 존재하지 않습니다")
	}
	filter := bson.D{{Key: "_id", Value: change.OrderId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "orderedmenus", Value: order.Orderedmenus}}}}
	_, err := o.Collection.UpdateOne(context.TODO(), filter, update)
	util.PanicHandler(err)

	return nil
}


// 주문 메뉴 추가하기
func (o *OrderedListModel) AddOrder(addStruct *AddMenuType, legacyOrder *OrderedList) primitive.ObjectID {
	// 추가하고자 하는 음식의 아이디를 이전 주문의 음식 배열에 넣어주는 로직
	filter := bson.D{{Key: "_id", Value: addStruct.OrderId}}
	legacyOrder.Orderedmenus = append(legacyOrder.Orderedmenus, addStruct.NewItem)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "orderedmenus", Value: legacyOrder.Orderedmenus}}}}
	_, err := o.Collection.UpdateOne(context.TODO(), filter, update)
	util.PanicHandler(err)

	return addStruct.OrderId
}


// 일별 주문 count +1 해주는 함수
func (o *OrderedListModel) DayOrderCount(sid string) int {
	// string을 objectId 타입으로 바꿔줌
	id, err := primitive.ObjectIDFromHex(sid)
	util.PanicHandler(err)
	filter := bson.D{{Key: "_id", Value: id}}
	counter := &DayCounter{}

	// 오늘 주문량을 가져오고
	o.Collection.FindOne(context.TODO(), filter).Decode(counter)
	// 변수에 담고
	num := counter.Count
	// +1로 업데이트한다.
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "daycount", Value: num + 1}}}}

	_, err = o.Collection.UpdateOne(context.TODO(), filter, update)
	util.PanicHandler(err)
	
	// 오늘 주문량 반환
	return num
}