//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=interfaces_mock.go

package service

import (
	"Hezzl_testing/internal/service/dto"
)

type ProjectGoodService struct {
	repo Repository
}

func NewService(repo Repository) *ProjectGoodService {
	return &ProjectGoodService{repo: repo}
}

type Repository interface {
	Create(projectId int, name string) (*dto.CreateGoodResponse, error)
	Update(id, projectId int, name, description string) (*dto.UpdateGoodResponse, error)
	Delete(id, projectId int) (*dto.DeleteGoodResponse, error)
	Get(limit, page int) (*dto.GetGoodsResponse, error)
	PatchPriority(id, projectId, newPriority int) (*dto.ReprioritizeResponse, *dto.UpdateGoodsResponse, error)
}
