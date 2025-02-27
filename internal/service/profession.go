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
