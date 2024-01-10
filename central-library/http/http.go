package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/rbojan2000/central-library/model"
	"github.com/rbojan2000/central-library/repository"
)

type Server struct {
	repository repository.Repository
}

func NewServer(repository repository.Repository) *Server {
	return &Server{repository: repository}
}

func (s Server) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	user, err := s.repository.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) CreateUser(ctx *gin.Context) {
	var newUser model.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if _, err := s.repository.GetUser(ctx, newUser.ID); err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user exists"})
		return
	}

	newUser.NumOfRentedBooks = 0
	newUser, err := s.repository.CreateUser(ctx, newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": newUser})
}

func (s Server) UpdateUserNumOfBooksRented(ctx *gin.Context) {
	membership := ctx.Param("membership")
	numStr := ctx.Param("num")

	if membership == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument membership"})
		return
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument num"})
		return
	}

	user, err := s.repository.GetUserByMembership(ctx, membership)

	if user.NumOfRentedBooks == 3 && num == 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user has already rented 3 books"})
		return
	}

	user.NumOfRentedBooks = user.NumOfRentedBooks + num

	user, err = s.repository.UpdateUser(ctx, user)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("ID")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	if err := s.repository.DeleteUser(ctx, id); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
