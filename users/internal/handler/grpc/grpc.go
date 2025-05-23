package grpc

import (
	"context"

	"github.com/Binit-Dhakal/e-fit/users/gen"
	"github.com/Binit-Dhakal/e-fit/users/internal/model"
	"github.com/Binit-Dhakal/e-fit/users/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedUserServiceServer
	ctrl ports.UserService
}

func NewHandler(ctrl ports.UserService) *Handler {
	return &Handler{
		ctrl: ctrl,
	}
}

func (h *Handler) GetUserByID(ctx context.Context, req *gen.GetUserByIDRequest) (*gen.GetUserByIDResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil or empty id")
	}

	user, err := h.ctrl.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.GetUserByIDResponse{User: model.UserToProto(user)}, nil
}

func (h *Handler) GetUserByEmail(ctx context.Context, req *gen.GetUserByEmailRequest) (*gen.GetUserByEmailResponse, error) {
	if req == nil || req.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil or empty email")
	}

	user, err := h.ctrl.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.GetUserByEmailResponse{User: model.UserToProto(user)}, nil
}

func (h *Handler) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	if req == nil || req.User == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil or empty user")
	}

	user := model.ProtoToUser(req.User)
	err := h.ctrl.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.CreateUserResponse{}, nil
}
