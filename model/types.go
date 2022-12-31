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

type IdType struct {
	Id primitive.ObjectID `json:"id"`
}

type AddMenuDataType struct {
	Name string
	Orderable bool
	Limit int
	Price int
	From string
}

type UpdateMenuDataType struct {
	Id primitive.ObjectID
	Key string
	Value string
}

type AddOrderType struct {
	Buyerinfo BuyerInfo
	OrderedMenus []OrderedMenu
}

type AddMenusType struct {
	NewItem OrderedMenu
	OrderId primitive.ObjectID
	Neworder AddOrderType
}