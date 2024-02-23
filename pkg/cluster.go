package pkg

import "log"

//cluster will contain all the DBControllers and can be used for monitoring

type DBCluster struct {
	Cluster []DBController
}

func NewCluster() *DBCluster {
	return &DBCluster{}
}

func (c *DBCluster) PingAll() {
	log.Printf("Testing the connection to all dbs")

	for _, db := range c.Cluster {
		db.MutexDb.TryLock()
		sqlDb, err := db.DB.DB()
		if err != nil {
			log.Printf("error testing ping for db %s:%s", db.Config.Host, db.Config.Port)
			continue
		}
		err = sqlDb.Ping()
		if err != nil {
			log.Printf("the %s:%s db is down", db.Config.Host, db.Config.Port)
			continue
		}
		db.MutexDb.Unlock()
		log.Printf("the %s:%s db is up", db.Config.Host, db.Config.Port)
	}
}
