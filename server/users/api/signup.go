package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shubhamjagdhane/myBooksShelf/server/users/context"
	"github.com/shubhamjagdhane/myBooksShelf/server/users/util"
)

// CreateAccountRequest structure to create user
type CreateAccountRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// createAccount takes user information and create account for the users
func (server *Server) createAccount(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-type", "application/json; charset=UTF-8")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := context.Users{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword,
	}

	newUser, err := server.store.CreateUser(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newUser)
}
