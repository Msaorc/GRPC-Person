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

func (pf *ProfessionService) CreateProfession(ctx context.Context, in *pb.CreateProfessionRequest) (*pb.ProfessionResponse, error) {
	profession, err := pf.ProfessionDB.Create(in.Description)
	if err != nil {
		return nil, err
	}

	professionResponse := &pb.Profession{
		Id:          profession.ID,
		Description: profession.Description,
	}

	return &pb.ProfessionResponse{Profession: professionResponse}, nil
}
