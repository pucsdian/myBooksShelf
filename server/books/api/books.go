package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shubhamjagdhane/myBooksShelf/server/books/context"
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
	ctx.Writer.Header().Set("Content-type", "application/json; charset=UTF-8")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
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
	ctx.Writer.Header().Set("Content-type", "application/json; charset=UTF-8")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
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
	ctx.Writer.Header().Set("Content-type", "application/json; charset=UTF-8")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
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
