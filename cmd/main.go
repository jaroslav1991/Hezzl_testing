package main

import (
	"Hezzl_testing/internal/config"
	"Hezzl_testing/internal/connection"
	storage "Hezzl_testing/internal/db"
	"Hezzl_testing/internal/handlers"
	"Hezzl_testing/internal/model"
	"Hezzl_testing/internal/service"
	"Hezzl_testing/internal/service/repository"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	dbConf := config.GetDBConfig()

	db, err := connection.NewPostgresDB(dbConf)
	if err != nil {
		log.Println(err)
		return
	}

	st := storage.NewStorage(db)
	repo := repository.NewRepository(db)
	srv := service.NewService(repo)
	hd := handlers.NewHandler(srv)

	if err := model.InitSchema(st); err != nil {
		log.Println(err)
		return
	}

	r := gin.Default()

	r.POST("/good/create/", hd.CreateHandler())
	r.PATCH("/good/update/", hd.UpdateHandler())
	r.PATCH("/good/reprioritize/", hd.PriorityHandler())
	r.DELETE("/good/remove/", hd.DeleteHandler())
	r.GET("/goods/list/", hd.GetHandler())

	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}

}
