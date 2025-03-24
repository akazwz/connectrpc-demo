package auth

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	authv1 "opendrive/gen/auth/v1"
	"opendrive/gen/auth/v1/authv1connect"
	"opendrive/models"
	"opendrive/shared/database"
	"opendrive/utils"
	"time"
)

func NewAuthService() authv1connect.PublicAuthServiceHandler {
	return &service{}
}

type service struct{}

func (s service) Login(ctx context.Context, req *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.LoginResponse], error) {
	db := database.GetDataBase()
	var user models.User
	if err := db.Model(models.User{}).Where("email = ?", req.Msg.Email).First(&user).Error; err != nil {
		return nil, err
	}
	if !utils.ComparePassword(user.Password, req.Msg.Password) {
		return nil, nil
	}
	accessToken := utils.SignToken(user.ID, time.Now().Add(time.Hour*24))
	var refreshToken = models.RefreshToken{
		UserID: user.ID,
	}
	if err := db.Model(models.RefreshToken{}).Create(&refreshToken).Error; err != nil {
		return nil, err
	}
	return connect.NewResponse[authv1.LoginResponse](
		&authv1.LoginResponse{
			User: &authv1.User{
				Id:        user.ID,
				Email:     user.Email,
				Username:  user.Username,
				Avatar:    user.Avatar,
				CreatedAt: user.CreatedAt.String(),
				UpdatedAt: user.UpdatedAt.String(),
			},
			AccessToken:  accessToken,
			RefreshToken: refreshToken.Token,
		}), nil
}

func (s service) Register(ctx context.Context, req *connect.Request[authv1.RegisterRequest]) (*connect.Response[authv1.RegisterResponse], error) {
	db := database.GetDataBase()
	user := models.User{
		Email:    req.Msg.Email,
		Username: req.Msg.Username,
		Password: req.Msg.Password,
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	accessToken := utils.SignToken(user.ID, time.Now().Add(time.Hour*24))
	var refreshToken = models.RefreshToken{
		UserID: user.ID,
	}
	if err := db.Model(models.RefreshToken{}).Create(&refreshToken).Error; err != nil {
		return nil, err
	}
	return connect.NewResponse[authv1.RegisterResponse](
		&authv1.RegisterResponse{
			User: &authv1.User{
				Id:        user.ID,
				Email:     user.Email,
				Username:  user.Username,
				Avatar:    user.Avatar,
				CreatedAt: user.CreatedAt.String(),
				UpdatedAt: user.UpdatedAt.String(),
			},
			AccessToken:  accessToken,
			RefreshToken: refreshToken.Token,
		}), nil
}

func (s service) RefreshToken(ctx context.Context, req *connect.Request[authv1.RefreshTokenRequest]) (*connect.Response[authv1.RefreshTokenResponse], error) {
	db := database.GetDataBase()
	var refreshToken models.RefreshToken
	if err := db.Model(models.RefreshToken{}).Where("token = ?", req.Msg.RefreshToken).Find(&refreshToken).Error; err != nil {
		return nil, err
	}
	if refreshToken.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}
	accessToken := utils.SignToken(refreshToken.UserID, time.Now().Add(time.Hour*24))
	return connect.NewResponse[authv1.RefreshTokenResponse](&authv1.RefreshTokenResponse{
		AccessToken: accessToken,
	}), nil
}

func (s service) Whoami(ctx context.Context, req *connect.Request[authv1.WhoamiRequest]) (*connect.Response[authv1.WhoamiResponse], error) {
	userID := ctx.Value("user_id").(string)
	db := database.GetDataBase()
	var user models.User
	if err := db.Model(models.User{}).Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return connect.NewResponse[authv1.WhoamiResponse](
		&authv1.WhoamiResponse{
			User: &authv1.User{
				Id:        user.ID,
				Email:     user.Email,
				Username:  user.Username,
				Avatar:    user.Avatar,
				CreatedAt: user.CreatedAt.String(),
				UpdatedAt: user.UpdatedAt.String(),
			},
		}), nil
}
