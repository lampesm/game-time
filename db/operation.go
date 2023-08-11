package db

import (
	"github.com/lampesm/game-time/config"
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

func GetGames() ([]forms.GamePayload, error) {
	var games []forms.GamePayload

	if err := DB.Table(config.GAMES_TBL).
		Select("name, description, image, cover, category").
		Scan(&games).
		Error; err != nil {
		return nil, err
	}

	return games, nil
}
