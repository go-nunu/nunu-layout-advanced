package uuid

import uuid "github.com/satori/go.uuid"

func GenUUID() string {
	return uuid.NewV4().String()
}
