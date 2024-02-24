package clients

import "github.com/spoonbuoy/lcdb/clients/kafka"

func Init() {
	kafka.Init()
	kafka.InitProducer()
	kafka.InitConsumer()
}
