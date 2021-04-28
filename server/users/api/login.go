package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shubhamjagdhane/myBooksShelf/server/users/context"
	"github.com/shubhamjagdhane/myBooksShelf/server/users/util"
)

// loginUserRequest structure to access login information for user
type loginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// loginUserResponse structure to response to user
type loginUserResponse struct {
	AccessToken string        `json:"access_token"`
	User        context.Users `json:"user"`
}

// loginUser identify user and response with the token and user information
func (server *Server) loginUser(ctx *gin.Context) {
	fmt.Println("The loginUser method is invoked...")
	ctx.Writer.Header().Set("Content-type", "application/json; charset=UTF-8")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = util.VerifyPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		ENV.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        user,
	}
	ctx.JSON(http.StatusOK, rsp)

}
