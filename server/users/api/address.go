package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shubhamjagdhane/myBooksShelf/server/users/context"
	"github.com/shubhamjagdhane/myBooksShelf/server/users/token"
)

// getAddressResponse structure to response for get address
type getAddressResponse struct {
	Username string `json:"username"`
	Address  context.Addresses
}

// setAddressRequest structure for to access the address
type setAddressRequest struct {
	City      string `json:"city"`
	Building  string `json:"building" binding:"required"`
	FlatNo    string `json:"flat_no" binding:"required"`
	Pincode   uint64 `json:"pincode" binding:"required"`
	IsDefault string `json:"is_default" binding:"required"`
}

// getAddress gives us address of user by specifying username
func (server *Server) getAddress(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload == nil {
		err := errors.New("Cannot authorize user to update address")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	address, err := server.store.GetAddress(authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var rsp getAddressResponse
	rsp.Address = address
	rsp.Username = authPayload.Username

	ctx.JSON(http.StatusOK, address)
}

// setAddress insert the address for the user
func (server *Server) setAddress(ctx *gin.Context) {
	var req setAddressRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload == nil {
		err := errors.New("Cannot authorize user to update address")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	address := context.Addresses{
		City:     req.City,
		Building: req.Building,
		FlatNo:   req.FlatNo,
		Pincode:  req.Pincode,
	}

	newAddress, err := server.store.SetAddress(authPayload.Username, address, req.IsDefault)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var rsp getAddressResponse
	rsp.Address = newAddress
	rsp.Username = authPayload.Username

	ctx.JSON(http.StatusOK, rsp)
}

// setDefaultAddressReq structure for accepting address id to set default
type setDefaultAddressReq struct {
	AddressID uint64 `json:"address_id" binding:"required"`
}

func (server *Server) setDefaultAddress(ctx *gin.Context) {
	var req setDefaultAddressReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload == nil {
		err := errors.New("Cannot authorize user to update address")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	message, err := server.store.SetDefaultAddress(req.AddressID, authPayload.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
