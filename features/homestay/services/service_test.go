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

func TestUpdate(t *testing.T) {
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
		ID:       1,
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

	t.Run("success update homestay", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, uint(1), *imageTrueCnv, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputData).Return(homestay.Core{}, errors.New("data not found")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("invalid jwt", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateToken(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		_, err := srv.Update(pToken, uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputData).Return(homestay.Core{}, errors.New("server problem")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.NotEqual(t, 0, res.ID)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewHomestayData(t)

	t.Run("success delete homestay", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(1)).Return(nil).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Delete", uint(2), uint(2)).Return(errors.New("data not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(2)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken, 2)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})
}

func TestShowDetail(t *testing.T) {
	repo := mocks.NewHomestayData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	resData := homestay.Core{
		ID:       1,
		Name:     "villa 1",
		Address:  "Jogja",
		Phone:    "081234567",
		Price:    500000,
		Image:    imageTrueCnv.Filename,
		Facility: "5 Kamar Tidur, 2 Kamar Mandi dalam, AC, Private Pool, dan Rooftop",
	}

	t.Run("success get homestay detail", func(t *testing.T) {
		repo.On("ShowDetail", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.ShowDetail(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("ShowDetail", uint(1)).Return(homestay.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		res, err := srv.ShowDetail(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.NotEqual(t, 0, res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("server problem", func(t *testing.T) {
		repo.On("ShowDetail", uint(1)).Return(homestay.Core{}, errors.New("server problem")).Once()
		srv := New(repo)
		res, err := srv.ShowDetail(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.NotEqual(t, 0, res.ID)
		repo.AssertExpectations(t)
	})
}
