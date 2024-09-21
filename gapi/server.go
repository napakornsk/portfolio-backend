package gapi

import (
	"github.com/napakornsk/cv-backend/config"
	"github.com/napakornsk/cv-backend/db"
	"github.com/napakornsk/cv-backend/pb"
	"github.com/napakornsk/cv-backend/token"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     config.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer()
