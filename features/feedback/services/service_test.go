package services

import (
	"airbnb/features/feedback"
	"airbnb/helper"
	"airbnb/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestFeedback(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)
	inputData := feedback.FeedbackCore{ID: 1, Rating: 5, Note: "bagus banget"}
	resData := feedback.FeedbackCore{ID: 1, Rating: 5, Note: "bagus banget"}

	t.Run("success add feedback", func(t *testing.T) {
		repo.On("AddFeedback", uint(1), inputData).Return(nil).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.AddFeedback(pToken, uint(1), inputData)
		assert.Nil(t, err)
		assert.NotEqual(t, resData.ID, err)
		repo.AssertExpectations(t)
	})

	t.Run("invalid jwt", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateToken(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.AddFeedback(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})

	// t.Run("feedback not found", func(t *testing.T) {
	// 	repo.On("AddFeedback", uint(1), uint(1), inputData).Return(errors.New("feedback not found")).Once()

	// 	srv := New(repo)
	// 	_, token := helper.GenerateToken(1)
	// 	pToken := token.(*jwt.Token)
	// 	pToken.Valid = true
	// 	res := feedback.FeedbackCore{}
	// 	err := srv.AddFeedback(pToken, uint(1), inputData)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, uint(0), res.ID)
	// 	assert.ErrorContains(t, err, "not found")
	// 	repo.AssertExpectations(t)
	// })

	// t.Run("server problem", func(t *testing.T) {
	// 	repo.On("AddFeedback", uint(1), uint(1), inputData).Return(errors.New("server problem")).Once()

	// 	srv := New(repo)
	// 	_, token := helper.GenerateToken(1)
	// 	pToken := token.(*jwt.Token)
	// 	pToken.Valid = true
	// 	res := feedback.FeedbackCore{}
	// 	err := srv.AddFeedback(pToken, uint(1), inputData)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, uint(0), res.ID)
	// 	assert.ErrorContains(t, err, "server")
	// 	repo.AssertExpectations(t)
	// })
}

func TestListFeedback(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)

	resData := []feedback.FeedbackCore{{
		ID:     1,
		Rating: 5,
		Note:   "bagus sekali",
	}}

	t.Run("success get user feedback", func(t *testing.T) {
		repo.On("ListFeedback", uint(1), uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.ListFeedback(pToken, uint(1))
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("feedback not found", func(t *testing.T) {
		repo.On("ListFeedback", uint(1), uint(1)).Return([]feedback.FeedbackCore{}, errors.New("feedback not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.ListFeedback(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("there is a problem with the server", func(t *testing.T) {
		repo.On("ListFeedback", uint(1), uint(1)).Return([]feedback.FeedbackCore{}, errors.New("there is a problem with the server")).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.ListFeedback(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})
}
