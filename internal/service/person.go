package service

import (
	"context"
	"io"

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

func (p *PersonService) GetPerson(ctx context.Context, in *pb.PersonGetRequest) (*pb.Person, error) {
	person, err := p.PersonDB.Find(in.Id)
	if err != nil {
		return nil, err
	}

	personResponse := &pb.Person{
		Id:   person.ID,
		Name: person.Name,
		Year: person.Year,
	}

	return personResponse, nil
}

func (p *PersonService) CreatePersonStream(stream pb.PersonService_CreatePersonStreamServer) error {
	people := &pb.PersonList{}

	for {
		person, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(people)
		}
		if err != nil {
			return err
		}

		personResult, err := p.PersonDB.Create(person.Name, person.Year)
		if err != nil {
			return err
		}

		people.People = append(people.People, &pb.Person{
			Id:   personResult.ID,
			Name: personResult.Name,
			Year: person.Year,
		})
	}
}
