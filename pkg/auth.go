package auth

import (
	"context"

	"github.com/killerekt/grpc-prac/pb"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Email    string `gorm:"type:text;not null;primary key"`
	Password string `gorm:"type:text;not null"`
}

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	DB *gorm.DB
}

func NewAuthService(DB *gorm.DB) AuthService {
	return AuthService{
		DB: DB,
	}
}

func (s *AuthService) SignUp(context context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	err = s.DB.Create(User{Email: req.Email, Password: string(hashedPassword)}).Error
	if err != nil {
		return nil, err
	}

	res := &pb.AuthResponse{Message: "Successfully SignUp the User", AccessToken: "Currently not implemented"}

	return res, nil
}

func (s *AuthService) Login(context context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	var user User
	err := s.DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	res := &pb.AuthResponse{Message: "Successfully logged in", AccessToken: "Currently not implemented"}
	return res, nil
}
