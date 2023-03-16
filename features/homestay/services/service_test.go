package services

import (
	"airbnb/features/homestay"
	"airbnb/helper"
	"airbnb/mocks"
	"errors"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewHomestayData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	inputData := homestay.Core{
		ID:       0,
		Name:     "villa 1",
		Address:  "Jogja",
		Phone:    "081234567",
		Price:    500000,
		Facility: "5 Kamar Tidur, 2 Kamar Mandi dalam, AC, Private Pool, dan Rooftop",
	}
	resData := homestay.Core{
		ID:       1,
		Name:     "villa 1",
		Address:  "Jogja",
		Phone:    "081234567",
		Price:    500000,
		Facility: "5 Kamar Tidur, 2 Kamar Mandi dalam, AC, Private Pool, dan Rooftop",
	}

	t.Run("success post homestay", func(t *testing.T) {
		repo.On("Add", uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, *imageTrueCnv, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("invalid jwt", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateToken(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		_, err := srv.Add(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})

	t.Run("cannot post homestay", func(t *testing.T) {
		repo.On("Add", uint(1), mock.Anything).Return(homestay.Core{}, errors.New("server error"))
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}
