package service

import (
	"github.com/ra192/card-api-go/model"
	"log"
)

func СreateCard(card *model.Card)  {
	model.DB.Create(card)
	log.Printf("Customer was created with id: %d", card.ID)
}
