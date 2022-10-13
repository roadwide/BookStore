package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetUUID() string {
	rand.Seed(time.Now().UnixNano())
	uuid := time.Now().Format("20060102150405") + fmt.Sprintf("%d", rand.Int31n(1000000))
	return uuid
}
