package deleting

import (
	"github.com/google/uuid"
)

type DeleteArt struct {
	ArticleID uuid.UUID `json:"id"`
}
