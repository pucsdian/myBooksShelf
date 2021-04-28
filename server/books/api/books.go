package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shubhamjagdhane/myBooksShelf/server/books/context"
	"github.com/shubhamjagdhane/myBooksShelf/server/books/token"
)

// getBooksResponse structure to response for get address
type getBooksResponse struct {
	SearchResult []context.Books `json:"search_result"`
	TotalRecords uint64          `json:"total_records"`
}

// booksSearchRequest structure to access login information for user
type booksSearchRequest struct {
	SearchString string `json:"search_string"`
}

// getDefaultBooks gives us books with matching search_string with book title or author
func (server *Server) getDefaultBooks(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Allow-Headers", "*")
	ctx.Header("Content-Type", "application/json")
	var req booksSearchRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if req.SearchString == "" {
		rsp := getBooksResponse{
			TotalRecords: 0,
			SearchResult: make([]context.Books, 0),
		}
		ctx.JSON(http.StatusOK, rsp)
		return
	}

	books, err := server.store.GetDefaultBooks(req.SearchString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := getBooksResponse{
		TotalRecords: uint64(len(books)),
		SearchResult: books,
	}
	ctx.JSON(http.StatusOK, rsp)
}

// booksSearchOffsetRequest structure to access login information for user
type booksSearchOffsetRequest struct {
	SearchString string `json:"search_string"`
	Offset       int    `json:"offset"`
}

// getBooksByOffset gives us books with matching search_string with book title or author having mentioned offset
func (server *Server) getBooksByOffset(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Allow-Headers", "*")
	ctx.Header("Content-Type", "application/json")
	var req booksSearchOffsetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if req.SearchString == "" {
		rsp := getBooksResponse{
			TotalRecords: 0,
			SearchResult: make([]context.Books, 0),
		}
		ctx.JSON(http.StatusOK, rsp)
		return
	}

	books, err := server.store.GetBooksByOffset(req.SearchString, req.Offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := getBooksResponse{
		SearchResult: books,
		TotalRecords: uint64(len(books)),
	}
	ctx.JSON(http.StatusOK, rsp)
}

// getRandomBooks gives as random 50 books
func (server *Server) getRandomBooks(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Allow-Headers", "*")
	ctx.Header("Content-Type", "application/json")
	books, err := server.store.GetRandomBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := getBooksResponse{
		TotalRecords: uint64(len(books)),
		SearchResult: books,
	}
	ctx.JSON(http.StatusOK, rsp)
}

type addToCartRequest struct {
	BookID []int `json:"books_id"`
}

func (server *Server) addToCart(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Allow-Headers", "*")
	ctx.Header("Content-Type", "application/json")

	var req addToCartRequest
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

	err := server.store.AddToCart(req.BookID, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (server *Server) getCartItems(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Content-Type", "application/x-www-form-urlencoded")
	// ctx.Header("Access-Control-Allow-Methods", "*")
	// ctx.Header("Access-Control-Allow-Headers", "*")
	// ctx.Header("Content-Type", "application/json")
	// ctx.Header("Access-Control-Allow-Credentials", "*")

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload == nil {
		err := errors.New("Cannot authorize user to update address")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	books, err := server.store.GetCartItems(authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := getBooksResponse{
		TotalRecords: uint64(len(books)),
		SearchResult: books,
	}
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) preFlight(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "https://localhost:8081")
	fmt.Println("Authorization using preFlight: ", ctx.GetHeader("Authorization"))
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
