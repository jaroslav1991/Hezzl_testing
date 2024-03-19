package dto

import (
	"time"
)

type CreateGoodRequest struct {
	Name string `json:"name"`
}

type CreateGoodResponse struct {
	Id          int       `json:"id"`
	ProjectId   int       `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Removed     bool      `json:"removed"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateGoodRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateGoodResponse struct {
	Id          int       `json:"id"`
	ProjectId   int       `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Removed     bool      `json:"removed"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateGoodsResponse struct {
	Goods []Good `json:"goods"`
}

type DeleteGoodResponse struct {
	Id        int  `json:"id"`
	ProjectId int  `json:"project_id"`
	Removed   bool `json:"removed"`
}

type ReprioritizeRequest struct {
	NewPriority int `json:"newPriority"`
}

type ReprioritizeResponse struct {
	Id       int `json:"id"`
	Priority int `json:"priority"`
}

type Priorities struct {
	Id       int `json:"id"`
	Priority int `json:"priority"`
}

type GetGoodsResponse struct {
	Meta  Meta   `json:"meta"`
	Goods []Good `json:"goods"`
}

type Meta struct {
	Total   int `json:"total"`
	Removed int `json:"removed"`
	Limit   int `json:"limit"`
	Offset  int `json:"offset"`
}

type Good struct {
	Id          int       `json:"id"`
	ProjectId   int       `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Removed     bool      `json:"removed"`
	CreatedAt   time.Time `json:"created_at"`
}
