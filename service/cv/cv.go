package cv

import (
	"context"
	"log"

	"github.com/napakornsk/cv-backend/database"
	"github.com/napakornsk/cv-backend/models"
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

	cv := req.GetCv()
	tx := database.PostgresDb.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Omit("id").Create(cv).Error; err != nil {
		return nil, err
	}

	cvModel := new(models.CV)
	if err := tx.Last(cvModel).Error; err != nil {
		return nil, err
	}

	return &pb.CreateCVRes{
		Id: cvModel.ID,
	}, nil
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
				Id:  1,
				Pos: "TEST",
			},
		},
	}, nil
}
