package uuid

import uuid "github.com/google/uuid"

func GenUUID() string {
	return uuid.NewString()
}
