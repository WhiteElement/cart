package shoppinglist

import (
	"cartv2/cart/internal/item/item"
	"time"
)

type List struct {
	Id       int
	Name     string
	Items    []item.Item
	Archived bool
	Created  time.Time
	Updated  time.Time
}
