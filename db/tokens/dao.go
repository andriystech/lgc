package tokens

import (
	"fmt"

	"github.com/andriystech/lgc/models"
	"github.com/google/uuid"
)

const template = "ws://fancy-chat.io/ws&token=%s"

func NewToken() *models.Token {
	return &models.Token{
		Url: fmt.Sprintf(template, uuid.New().String()),
	}
}