package main

import (
	"context"
	"log"
	"time"

	"github.com/AndrewMislyuk/go-shop-logger/internal/config"
	"github.com/AndrewMislyuk/go-shop-logger/internal/repository"
	"github.com/AndrewMislyuk/go-shop-logger/internal/server"
	"github.com/AndrewMislyuk/go-shop-logger/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(cfg.DB.URI)
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.DB.Database)

	auditRepo := repository.NewAudit(db)
	auditService := service.NewAudit(auditRepo)
	auditSrv := server.NewAuditServer(auditService)
	srv := server.New(auditSrv)

	log.Println("SERVER STARTED: ", time.Now())

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
