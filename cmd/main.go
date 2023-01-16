package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"

	"github.com/Levap123/task-manager-tasks-service/internal/repository"
	"github.com/Levap123/task-manager-tasks-service/internal/repository/mongodb"
	"github.com/Levap123/task-manager-tasks-service/internal/service"
	"github.com/Levap123/task-manager-tasks-service/internal/transport/rpc"
	"github.com/Levap123/task-manager-tasks-service/proto"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()
	db, err := mongodb.InitDb()
	if err != nil {
		log.Fatalln(err)
	}
	tasksCollection := mongodb.CreateTasksCollection(db)
	repo := repository.NewRepoMongo(tasksCollection)
	service := service.NewService(repo)
	handlerRPC := rpc.NewHandler(service)
	fmt.Println(os.Getenv("TASKS_ADDRESS"), 123)
	listener, err := net.Listen("tcp", os.Getenv("TASKS_ADDRESS"))
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	proto.RegisterTaskManagerServer(s, handlerRPC)
	if err := s.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
