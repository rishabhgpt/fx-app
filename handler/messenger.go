package handler

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/rishabhgpt/fx-app/proto"
)

type Handler = pb.MessageServiceServer

type handler struct {
	pb.UnimplementedMessageServiceServer
}

func New() (Handler, error) {
	return &handler{}, nil

}

func (h *handler) SendMessage(ctx context.Context, r *pb.Message) (*pb.Reply, error) {
	return &pb.Reply{Status: r.Body, Id: uuid.NewString()}, nil
}
