package grpc

import (
	"context"

	"github.com/MuhammedAshifVnr/user_service/internal/models"
	"github.com/MuhammedAshifVnr/user_service/internal/service"
	pb "github.com/MuhammedAshifVnr/user_service/proto"
)

// UserHandler handles gRPC requests for user operations
type UserHandler struct {
	userService service.UserService
	pb.UnimplementedUserServiceServer
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUserByID fetches a user by ID
func (h *UserHandler) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.User, error) {
	user, err := h.userService.GetUserByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:      uint64(user.ID),
		Fname:   user.FName,
		City:    user.City,
		Phone:   user.Phone,
		Height:  user.Height,
		Married: user.Married,
	}, nil
}

// GetUsersByIDs fetches multiple users by their IDs
func (h *UserHandler) GetUsersByIDs(ctx context.Context, req *pb.GetUsersByIDsRequest) (*pb.GetUsersByIDsResponse, error) {
	ids := make([]uint, len(req.Ids))
	for i, id := range req.Ids {
		ids[i] = uint(id)
	}

	users, err := h.userService.GetUsersByIDs(ids)
	if err != nil {
		return nil, err
	}

	response := &pb.GetUsersByIDsResponse{}
	for _, user := range users {
		response.Users = append(response.Users, &pb.User{
			Id:      uint64(user.ID),
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}
	return response, nil
}

// SearchUsers searches for users based on criteria
func (h *UserHandler) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	users, err := h.userService.SearchUsers(req.City, req.Phone, req.Query, req.Married, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	response := &pb.SearchUsersResponse{}
	for _, user := range users {
		response.Users = append(response.Users, &pb.User{
			Id:      uint64(user.ID),
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}
	return response, nil
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.Empty, error) {
	user := &models.User{
		FName:   req.Fname,
		City:    req.City,
		Phone:   req.Phone,
		Height:  req.Height,
		Married: req.Married,
	}

	if err := h.userService.CreateUser(user); err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
