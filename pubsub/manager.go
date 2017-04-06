package pubsub

import (
	"math/rand"
	"time"
)

//use randomMessage
const (
	letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// RandomInterval is return duration time for asyncMode
func RandomInterval(max int) time.Duration {
	var td time.Duration
	if maxInterval > 0 {
		interval := rand.Intn(maxInterval)
		td = time.Duration(interval) * time.Millisecond
	}
	return td
}

// RandomMessage is
func RandomMessage(strlen int) string {
	message := make([]byte, strlen)
	cache, remain := rand.Int63(), letterIdxMax
	for i := strlen - 1; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		idx := int(cache & letterIdxMask)
		if idx < len(letters) {
			message[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(message)

}
