package cv

import (
	"context"
	"log"

	pb "github.com/napakornsk/cv-backend/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type CVServer struct {
	MongoClient    *mongo.Client
	PostgresClient *gorm.DB
	pb.UnimplementedCVServiceServer
}

func (s *CVServer) CreateCV(ctx context.Context, req *pb.CreateCVReq) (
	*pb.CreateCVRes,
	error,
) {
	log.Println("Received CreateCV request")

	// Example logic to create a CV
	// cv := req.GetCv() // Extract the CV object from the request
	return nil, nil
}

func (s *CVServer) GetCV(ctx context.Context, req *pb.GetCVReq) (
	*pb.GetCVRes,
	error,
) {
	log.Printf("GetCV was invoked with ID: %d", req.GetId())

	return &pb.GetCVRes{
		Cv: &pb.CV{
			Id: 1,
			Intro: &pb.Intro{
				Id: 1,
				Pos: "TEST",
			},
		},
	}, nil
}
