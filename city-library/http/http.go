package http

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/rbojan2000/city/config"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rbojan2000/city/model"
	"github.com/rbojan2000/city/repository"
)

type Server struct {
	repository repository.Repository
}

func NewServer(repository repository.Repository) *Server {
	return &Server{repository: repository}
}

func (s Server) GetBorrow(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	borrow, err := s.repository.GetBorrow(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrBorrowNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"borrow": borrow})
}

func (s Server) evidentBorrow(membershipID string, num string) (bool, error) {
	config := config.NewConfig()
	url := fmt.Sprintf("http://%s:%s/users/%s/%s", config.CentralLibraryHost, config.CentralLibraryPort, membershipID, num)

	data := []byte(`{}`)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return false, fmt.Errorf("Central-library not avaliable!")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("User with membership %s already has 3 rented books.", membershipID)
	}

	return true, nil
}

func (s Server) CreateBorrow(ctx *gin.Context) {

	var borrow model.Borrow
	if err := ctx.ShouldBindJSON(&borrow); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	response, _ := s.evidentBorrow(borrow.Membership, "1")
	if response != true {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User already has 3 rented books."})
		return
	}

	println(response)

	borrow, err := s.repository.CreateBorrow(ctx, borrow)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"borrow": borrow})
}

func (s Server) DeleteBorrow(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	borrow, _ := s.repository.GetBorrow(ctx, id)
	s.evidentBorrow(borrow.Membership, "-1")

	if err := s.repository.DeleteBorrow(ctx, id); err != nil {
		if errors.Is(err, repository.ErrBorrowNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
