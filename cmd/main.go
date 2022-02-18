package main

import (
	"log"

	"github.com/zhansul19/qltIntern/pkg/handlers"
	"github.com/zhansul19/qltIntern/pkg/repository"
	"github.com/zhansul19/qltIntern/pkg/service"

	"github.com/zhansul19/qltIntern"
	_ "github.com/lib/pq"
)
func main(){
	db,err:=repository.NewPostgresDB(repository.Config{
		Host:      "localhost",
		Port:      "5433",
		Username: 	"root",
		Password:  "123456",
		DbName:    "finance_db",
		SSLmode:   "disable",	
	})
	if err != nil {
		log.Fatalf("Error initializing db: %s",err.Error())
	}

	repository:=repository.NewRepository(db)
	service:=service.NewService(repository)
	handlers:=handlers.NewHandler(service)
	server:=new(finance.Server)
	if err:=server.Run("8080",handlers.InitRoutes());err!=nil {
		log.Fatalf("Error occured while running server: %s",err.Error())
	}
}