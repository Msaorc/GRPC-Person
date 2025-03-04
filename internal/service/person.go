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

func (p *PersonService) ListPerson(ctx context.Context, in *pb.Blank) (*pb.PersonList, error) {
	peopleList, err := p.PersonDB.FindAll()
	if err != nil {
		return nil, err
	}

	var peopleResponse []*pb.Person
	for _, person := range peopleList {
		people := &pb.Person{
			Id:   person.ID,
			Name: person.Name,
			Year: person.Year,
		}
		peopleResponse = append(peopleResponse, people)
	}
	return &pb.PersonList{People: peopleResponse}, nil
}
