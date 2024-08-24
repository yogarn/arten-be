package rest

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yogarn/arten/entity"
	"github.com/yogarn/arten/model"
	"github.com/yogarn/arten/pkg/response"
)

func (r *Rest) GetLoginUser(ctx *gin.Context) {
	user, _ := r.jwt.GetLoginUser(ctx)
	response.Success(ctx, http.StatusOK, "success", user)
}

func (r *Rest) Register(ctx *gin.Context) {
	userReq := &model.UserRegister{}
	if err := ctx.ShouldBindJSON(userReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}
	user, err := r.service.UserService.Register(userReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to register user", err)
		return
	}
	response.Success(ctx, http.StatusCreated, "success", user)
}

func (r *Rest) Login(ctx *gin.Context) {
	userReq := &model.UserLogin{}
	if err := ctx.ShouldBindJSON(userReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}
	user, err := r.service.UserService.Login(userReq)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("username doesn't exists")
		}
		response.Error(ctx, http.StatusInternalServerError, "failed to login", err)
		return
	}

	userDetails, err := r.service.UserService.GetUserByUsername(userReq.Username)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get user id", err)
		return
	}

	session := &entity.Session{
		Id:           uuid.New(),
		UserId:       userDetails.ID,
		RefreshToken: user.RefreshToken,
		IpAddress:    r.service.UserService.GetIpAddress(ctx),
		DeviceInfo:   r.service.UserService.GetDeviceInfo(ctx),
		CreatedAt:    time.Now(),
	}

	err = r.service.SessionService.CreateSession(session)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to create session", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success", user)
}

func (r *Rest) RefreshToken(ctx *gin.Context) {
	tokenReq := &model.RefreshToken{}
	if err := ctx.ShouldBindJSON(tokenReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	newToken, err := r.service.UserService.RefreshToken(tokenReq.Token)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to refresh token", err)
		return
	}

	err = r.service.SessionService.CheckSession(tokenReq.Token)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "invalid token", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success", newToken)
}

func (r *Rest) Logout(ctx *gin.Context) {
	tokenReq := &model.RefreshToken{}
	if err := ctx.ShouldBindJSON(tokenReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	userId, err := r.jwt.GetLoginUser(ctx)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "unauthorized", err)
		return
	}

	err = r.service.SessionService.DeleteSession(userId, tokenReq.Token)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to delete session", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success", "session deleted successfully")
}

func (r *Rest) UpdateProfile(ctx *gin.Context) {
	var userReq model.UpdateUser

	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	user, err := r.service.UserService.UpdateUser(ctx, &userReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to update profile", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success", user)
}

func (r *Rest) SendOtp(ctx *gin.Context) {
	var jsonData map[string]string

	if err := ctx.ShouldBindJSON(&jsonData); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	username, exists := jsonData["username"]
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "username not provided", nil)
		return
	}

	err := r.service.UserService.SendOtp(username)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to send otp", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success", "otp sent successfully")
}

func (r *Rest) ActivateUser(ctx *gin.Context) {
	var otpRequest model.OtpRequest

	if err := ctx.ShouldBindJSON(&otpRequest); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	err := r.service.UserService.ActivateUser(&otpRequest)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed to activate user", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success", "account activated successfully")
}
