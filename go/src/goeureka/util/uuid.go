package util

import "twinj/uuid"

func GetUUID() string {
        return uuid.NewV4().String()
}
