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

func TestProjectGoodService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	response := dto.GetGoodsResponse{
		Meta: dto.Meta{
			Total:   2,
			Removed: 0,
			Limit:   2,
			Offset:  1,
		},
		Goods: []dto.Good{{
			Id:          1,
			ProjectId:   1,
			Name:        "t1",
			Description: "",
			Priority:    1,
			Removed:     false,
			CreatedAt:   now,
		}, {
			Id:          2,
			ProjectId:   1,
			Name:        "t2",
			Description: "",
			Priority:    2,
			Removed:     false,
			CreatedAt:   now,
		}},
	}

	repo := NewMockRepository(ctrl)
	repo.EXPECT().Get(2, 1).Return(&response, nil)

	service := NewService(repo)
	actual, err := service.Get(2, 1)
	assert.NoError(t, err)
	assert.Equal(t, &response, actual)
}

func TestProjectGoodService_Get_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().Get(2, 1).Return(nil, errors.New("some error"))

	service := NewService(repo)
	_, err := service.Get(2, 1)
	assert.Error(t, err)
}

func TestProjectGoodService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := dto.UpdateGoodRequest{
		Name:        "t2",
		Description: "desc",
	}

	response := dto.UpdateGoodResponse{
		Id:          1,
		ProjectId:   1,
		Name:        "t2",
		Description: "desc",
		Priority:    1,
		Removed:     false,
		CreatedAt:   time.Time{},
	}

	repo := NewMockRepository(ctrl)
	repo.EXPECT().Update(1, 1, "t2", "desc").Return(&response, nil)

	service := NewService(repo)
	actual, err := service.Update(request, 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, &response, actual)
}

func TestProjectGoodService_Update_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := dto.UpdateGoodRequest{
		Name:        "t2",
		Description: "desc",
	}

	repo := NewMockRepository(ctrl)
	repo.EXPECT().Update(1, 1, "t2", "desc").Return(nil, errors.New("some error"))

	service := NewService(repo)
	_, err := service.Update(request, 1, 1)
	assert.Error(t, err)
}

func TestProjectGoodService_PatchPriority(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := dto.ReprioritizeRequest{NewPriority: 5}

	response := dto.ReprioritizeResponse{
		Id:       1,
		Priority: 5,
	}

	updatedList := dto.UpdateGoodsResponse{Goods: []dto.Good{{
		Id:          1,
		ProjectId:   1,
		Name:        "t1",
		Description: "",
		Priority:    1,
		Removed:     false,
		CreatedAt:   time.Time{},
	}, {
		Id:          2,
		ProjectId:   01,
		Name:        "t2",
		Description: "",
		Priority:    2,
		Removed:     false,
		CreatedAt:   time.Time{},
	}}}

	repo := NewMockRepository(ctrl)
	repo.EXPECT().PatchPriority(1, 1, request.NewPriority).Return(&response, &updatedList, nil)

	service := NewService(repo)
	actual, _, err := service.PatchPriority(request, 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, &response, actual)
}

func TestProjectGoodService_PatchPriority_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := dto.ReprioritizeRequest{NewPriority: 5}

	repo := NewMockRepository(ctrl)
	repo.EXPECT().PatchPriority(1, 1, request.NewPriority).Return(nil, nil, errors.New("some error"))

	service := NewService(repo)
	_, _, err := service.PatchPriority(request, 1, 1)
	assert.Error(t, err)
}
