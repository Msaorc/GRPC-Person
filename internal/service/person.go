package service

import (
	"context"

	"github.com/Msaorc/GRPC-Person/internal/database"
	"github.com/Msaorc/GRPC-Person/internal/pb"
)

type PersonService struct {
	pb.UnimplementedPersonServiceServer
	PersonDB database.Person
}

func NewPersonService(personDB database.Person) *PersonService {
	return &PersonService{
		PersonDB: personDB,
	}
}

func (p *PersonService) CreatePerson(ctx context.Context, in *pb.CreatePersonRequest) (*pb.Person, error) {
	person, err := p.PersonDB.Create(in.Name, in.Year)
	if err != nil {
		return nil, err
	}

	return &pb.Person{
		Id:   person.ID,
		Name: person.Name,
		Year: person.Year,
	}, nil
}
