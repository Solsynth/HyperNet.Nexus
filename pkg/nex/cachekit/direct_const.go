package cachekit

import "fmt"

// Those constants are used to directly get the cached data from redis
// Formatted like {prefix}#{key}
const (
	DAAttachment = "attachment"
	DAUser       = "account"
)

func FKey(prefix string, key any) string {
	return fmt.Sprintf("%s#%v", prefix, key)
}
