package handler

import (
	"airbnb/features/user"
	"errors"
)

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func ToResponse(data user.Core) UserResponse {
	return UserResponse{
		ID:      data.ID,
		Name:    data.Name,
		Email:   data.Email,
		Phone:   data.Phone,
		Address: data.Address,
	}
}

type ProfileResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Address        string `json:"address"`
	Role           string `json:"role"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
}

func ToProfileResponse(data user.Core) ProfileResponse {
	return ProfileResponse{
		ID:             data.ID,
		Name:           data.Name,
		Email:          data.Email,
		Address:        data.Address,
		Role:           data.Role,
		Phone:          data.Phone,
		ProfilePicture: data.ProfilePicture,
	}
}

func ConvertUpdateResponse(input user.Core) (interface{}, error) {
	ResponseFilter := user.Core{}
	ResponseFilter = input
	result := make(map[string]interface{})
	if ResponseFilter.ID != 0 {
		result["id"] = ResponseFilter.ID
	}
	if ResponseFilter.Name != "" {
		result["name"] = ResponseFilter.Name
	}
	if ResponseFilter.Email != "" {
		result["email"] = ResponseFilter.Email
	}
	if ResponseFilter.Address != "" {
		result["address"] = ResponseFilter.Address
	}
	if ResponseFilter.Phone != "" {
		result["role"] = ResponseFilter.Role
	}
	if ResponseFilter.Phone != "" {
		result["phone"] = ResponseFilter.Phone
	}
	if ResponseFilter.Password != "" {
		result["password"] = ResponseFilter.Password
	}
	if ResponseFilter.ProfilePicture != "" {
		result["profile_picture"] = ResponseFilter.ProfilePicture
	}

	if len(result) <= 1 {
		return user.Core{}, errors.New("no data was change")
	}
	return result, nil
}
