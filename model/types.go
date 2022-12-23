package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 메뉴를 추가할 때 body에 해당하는 구조체
type AddMenuType struct {
	NewOrder OrderedList `json:"neworder"`
	NewItem primitive.ObjectID `json:"newitem"`
	OrderId primitive.ObjectID `json:"orderid"`
}


// 메뉴를 변경할 때 body에 해당하는 구조체
type ChangeMenuType struct {
	OrderId primitive.ObjectID `json:"orderid"`
	LegacyFoodId primitive.ObjectID `json:"legacyfoodid"`
	NewFoodId primitive.ObjectID `json:"newfoodid"`
}

type SuggestionType struct {
	Ids []primitive.ObjectID `json:"ids"`
}