package service

import (
	"context"
	"errors"
	"fmt"
	"server/model"
	"server/repo"
	"server/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	secretKey = "secretKey"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrTokenGeneration = errors.New("error generating token")
	ErrInternalServer  = errors.New("internal server error")
	ErrTimeoutExceeded = errors.New("timeout exceeded")
)

type UserService struct {
	userRepo repo.UserRepo
}

func NewService(userRepo repo.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(2)*time.Second)
	defer cancel()
	hashedPassword, err := utils.HashString(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.userRepo.CreateUser(ctx, user)
}

type JwtClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *UserService) Login(ctx context.Context, user *model.User) (*model.LoginUserResp, error) {
	// Set an arbitrary timeout duration
	ctx, cancel := context.WithTimeout(ctx, time.Duration(2)*time.Second)
	defer cancel()

	// Context error handling
	select {
	case <-ctx.Done():
		return nil, ErrTimeoutExceeded
	default:
		// Continue processing
	}

	// Get user by username
	dbUser, err := s.GetUserByUsername(ctx, user.Username)
	if errors.Is(err, ErrUserNotFound) {
		return nil, fmt.Errorf("bad request: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("internal server error: %w", err)
	}

	// Check hashed password
	err = utils.CheckHashedString(user.Password, dbUser.Password)
	if err != nil {
		return nil, fmt.Errorf("bad request: %w", ErrInvalidPassword)
	} else if err != nil {
		return nil, fmt.Errorf("internal server error: %w", err)
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		ID:       strconv.Itoa(dbUser.ID),
		Username: dbUser.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(dbUser.ID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, fmt.Errorf("internal server error: %w", ErrTokenGeneration)
	}

	return &model.LoginUserResp{
		AccessToken: ss,
		Username:    dbUser.Username,
		ID:          dbUser.ID,
	}, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(2)*time.Second)
	defer cancel()
	return s.userRepo.GetUser(ctx, username, false)
}

func (s *UserService) GetUserById(ctx context.Context, id int) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(2)*time.Second)
	defer cancel()
	return s.userRepo.GetUser(ctx, id, true)
}
