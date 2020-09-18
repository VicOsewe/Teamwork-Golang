package updating

import (
	"github.com/google/uuid"
)

type UpdateAtricle struct {
	ArticleID uuid.UUID `json:"id"`
	Message   string    `json:"message"`
	Title     string    `json:"title"`
}
