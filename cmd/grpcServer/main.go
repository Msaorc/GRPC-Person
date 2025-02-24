package main

import (
	"database/sql"

	"github.com/Msaorc/GRPC-Person/internal/database"
	"github.com/Msaorc/GRPC-Person/internal/pb"
	"github.com/Msaorc/GRPC-Person/internal/service"
	"google.golang.org/grpc"
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
	pb.RegisterPersonServiceServer(grpcServer, personService)
	pb.RegisterPersonServiceServer(grpcServer, professionService)
}
