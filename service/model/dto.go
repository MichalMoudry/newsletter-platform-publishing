package model

import "github.com/google/uuid"

type NewsletterUpdateModel struct {
	NewsletterId        uuid.UUID
	Name                string
	Description         string
	OldConcurrencyStamp uuid.UUID
}