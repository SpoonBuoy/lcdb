package services

import "github.com/spoonbuoy/lcdb/interfaces/icache"

type Services struct {
	cache icache.ICacheService
	db    icache.ICacheRepo
}

func InitServices() Services {
	return Services{}
}
