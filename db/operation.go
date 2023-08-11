package db

import (
	"github.com/lampesm/game-time/forms"
	"github.com/lampesm/game-time/models"
)

func InsertGame(payload forms.GamePayload) error {
	if err := DB.Create(&models.Game{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       &payload.Image,
		Cover:       &payload.Cover,
		Category:    payload.Category,
	}).Error; err != nil {
		return err
	}

	return nil
}
