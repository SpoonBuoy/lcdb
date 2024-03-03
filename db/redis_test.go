package db

import (
	"testing"
)

func TestGetRedisClient(t *testing.T) {
	client1 := GetRedisClient()
	client2 := GetRedisClient()

	if client1 != client2 {
		t.Fail()
	} else {
		//pass
	}
}
