package api

import "github.com/alexbsec/ragconverter/types"

type DivineRequesterInterface interface {
	GetItem(itemRequest types.ItemRequest) (types.ItemResponse, error)
}
