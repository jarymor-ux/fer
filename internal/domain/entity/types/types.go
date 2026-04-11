package types

import (
	"github.com/google/uuid"
	"github.com/hashicorp/go-set/v3"
)

type ChatID uuid.UUID
type HashSet[t any] set.HashSet[t, string]
