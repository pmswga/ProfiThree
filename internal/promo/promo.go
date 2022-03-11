package promo

import (
	"src/internal/participant"
	"src/internal/prize"
)

type Promo struct {
	Id           int                       `json:"id"`
	Name         string                    `json:"name"`
	Description  string                    `json:"description"`
	Prizes       []prize.Prize             `json:"prizes"`
	Participants []participant.Participant `json:"participants"`
}
