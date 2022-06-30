package app

import (
	"promotion/configs"
	"promotion/pkg/databases"
	"promotion/pkg/logger"
	"promotion/pkg/pubsub"
)

type infrastructure struct {
	log *logger.Logger
	db  databases.Databases
	ps  *pubsub.PubSub
}

func initInfra(cfg *configs.Config) *infrastructure {
	log := logger.New(cfg)

	db, err := databases.New(cfg, log)
	if err != nil {
		log.Fatalf("Failed to initialize DB")
	}

	ps, err := pubsub.NewPubSub(cfg, log)
	if err != nil {
		log.Fatalf("Failed to initialize PubSub")
	}
	log.Info("PubSub initialized successfully")

	return &infrastructure{log, db, ps}
}
