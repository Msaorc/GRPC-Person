package service

import (
	"context"

	"github.com/Msaorc/GRPC-Person/internal/database"
	"github.com/Msaorc/GRPC-Person/internal/pb"
)

type ProfessionService struct {
	pb.UnimplementedProfessionServiceServer
	ProfessionDB database.Profession
}

func NewProfessionService(professionDB database.Profession) *ProfessionService {
	return &ProfessionService{
		ProfessionDB: professionDB,
	}
}

func (pf *ProfessionService) CreateProfession(ctx context.Context, in *pb.CreateProfessionRequest) (*pb.Profession, error) {
	profession, err := pf.ProfessionDB.Create(in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.Profession{
		Id:          profession.ID,
		Description: profession.Description,
	}, nil
}

func (pf *ProfessionService) ListProfession(ctx context.Context, in *pb.Blank) (*pb.ProfessionList, error) {
	professionList, err := pf.ProfessionDB.FindAll()
	if err != nil {
		return nil, err
	}

	var professionResponse []*pb.Profession
	for _, profession := range professionList {
		p := &pb.Profession{
			Id:          profession.ID,
			Description: profession.Description,
		}
		professionResponse = append(professionResponse, p)
	}
	return &pb.ProfessionList{Professions: professionResponse}, nil
}

func (pf *ProfessionService) GetProfession(ctx context.Context, in *pb.ProfessionGetRequest) (*pb.Profession, error) {
	profession, err := pf.ProfessionDB.Find(in.Id)
	if err != nil {
		return nil, err
	}

	professionResponse := &pb.Profession{
		Id:          profession.ID,
		Description: profession.Description,
	}

	return professionResponse, nil
}
