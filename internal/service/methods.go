package service

import (
	"Hezzl_testing/internal/service/dto"
	"log"
)

func (s *ProjectGoodService) Create(projectId int, req dto.CreateGoodRequest) (*dto.CreateGoodResponse, error) {
	good, err := s.repo.Create(projectId, req.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return good, nil
}

func (s *ProjectGoodService) Update(req dto.UpdateGoodRequest, id, projectId int) (*dto.UpdateGoodResponse, error) {
	good, err := s.repo.Update(id, projectId, req.Name, req.Description)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return good, nil
}

func (s *ProjectGoodService) Delete(id, projectId int) (*dto.DeleteGoodResponse, error) {
	deleted, err := s.repo.Delete(id, projectId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return deleted, nil
}

func (s *ProjectGoodService) Get(limit, offset int) (*dto.GetGoodsResponse, error) {
	goods, err := s.repo.Get(limit, offset)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return goods, nil
}

func (s *ProjectGoodService) PatchPriority(req dto.ReprioritizeRequest, id, projectId int) (*dto.ReprioritizeResponse, *dto.UpdateGoodsResponse, error) {
	priority, updated, err := s.repo.PatchPriority(id, projectId, req.NewPriority)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	return priority, updated, nil
}
