package main

import (
	"context"
	"errors"
	"log"

	"github.com/micro/go-micro"

	pb "github.com/finalsatan/shiiip-vessel/proto/vessel"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type Repository struct {
	vessels []*pb.Vessel
}

func (repo *Repository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("can not find vessel by that spec")
}

type service struct {
	repo repository
}

func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	vessel, err := s.repo.FindAvailable(spec)
	if err != nil {
		return err
	}

	resp.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Murphy Zhang", MaxWeight: 200000, Capacity: 500},
	}

	repo := &Repository{vessels}

	srv := micro.NewService(
		micro.Name("shiiip.vessel"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to run shiiip vessel service server: %v", err)
	}
}
