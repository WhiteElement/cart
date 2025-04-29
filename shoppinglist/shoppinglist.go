package shoppinglist

import (
	"cartv2/cart/item/item"
	"time"
)

type Shoppinglist struct {
	Id      int
	Name    string
	Items   []item.Item
	Created time.Time
	Updated time.Time
}
