package user

import (
	"fmt"
	"spotssync/internal/auth"
	"spotssync/internal/domain/user/dto"
)

var ErrInvalidCredentials = fmt.Errorf("invalid email or password")

type service struct {
	repo       Repository
	jwtService auth.JWTService
}

func NewService(repo Repository, jwtService auth.JWTService) *service {
	return &service{repo, jwtService}
}

func (s *service) CreateUser(req dto.CreateUserRequest) (*dto.UserResponse, error) {

	user := User{
		Name:  req.Name,
		Email: req.Email,
		Role:  Role(req.Role),
	}

	// hash password and set to user.Password
	err := user.hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt.String(),
	}

	return &response, nil

}

func (s *service) LoginUser(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrInvalidCredentials // User not found
	}

	// check password
	err = user.checkPassword(req.Password)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// generate token
	token, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Name, string(user.Role[0]))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	response := dto.LoginResponse{
		User: dto.UserResponse{
			Token:     token,
			ID:        user.ID,
			Name:      user.Name,
			Role:      string(user.Role),
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
		},
	}
	return &response, nil
}

func (s *service) getUserByEmail(email string) (*dto.UserResponse, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	response := dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Role:      string(user.Role),
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
	}
	return &response, nil
}
