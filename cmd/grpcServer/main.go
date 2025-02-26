package main

import (
	"database/sql"
	"net"

	"github.com/Msaorc/GRPC-Person/internal/database"
	"github.com/Msaorc/GRPC-Person/internal/pb"
	"github.com/Msaorc/GRPC-Person/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	personDB := database.NewPerson(db)
	personService := service.NewPersonService(*personDB)
	professionDB := database.NewProfession(db)
	professionService := service.NewProfessionService(*professionDB)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPersonServiceServer(grpcServer, personService)
	pb.RegisterProfessionServiceServer(grpcServer, professionService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
