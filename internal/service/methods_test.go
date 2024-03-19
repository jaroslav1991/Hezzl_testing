package service

import (
	"Hezzl_testing/internal/service/dto"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestProjectGoodService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := dto.CreateGoodRequest{Name: "test1"}

	response := dto.CreateGoodResponse{
		ProjectId:   1,
		Name:        "test1",
		Description: "",
		Priority:    1,
		Removed:     false,
		CreatedAt:   time.Now(),
	}

	repo := NewMockRepository(ctrl)
	repo.EXPECT().Create(1, request.Name).Return(&response, nil)

	service := NewService(repo)

	actualData, err := service.Create(1, request)
	assert.NoError(t, err)
	assert.Equal(t, &response, actualData)
}

func TestProjectGoodService_Create_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := dto.CreateGoodRequest{Name: "test1"}
	repo := NewMockRepository(ctrl)
	repo.EXPECT().Create(1, request.Name).Return(nil, errors.New("some error"))

	service := NewService(repo)

	_, err := service.Create(1, request)
	assert.Error(t, err)
}

func TestProjectGoodService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	response := dto.DeleteGoodResponse{
		Id:        1,
		ProjectId: 1,
		Removed:   true,
	}
	repo := NewMockRepository(ctrl)
	repo.EXPECT().Delete(1, 1).Return(&response, nil)

	service := NewService(repo)
	actual, err := service.Delete(1, 1)
	assert.NoError(t, err)
	assert.Equal(t, &response, actual)
}

func TestProjectGoodService_Delete_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().Delete(1, 1).Return(nil, errors.New("some error"))

	service := NewService(repo)
	_, err := service.Delete(1, 1)
	assert.Error(t, err)
}
