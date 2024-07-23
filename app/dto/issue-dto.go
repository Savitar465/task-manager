package dto

import (
	"github.com/Savitar465/task-manager/app/models"
	"time"
)

type IssueResponse struct {
	ID          int       `json:"id"`
	TypeId      string    `json:"typeId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	DueDate     time.Time `json:"dueDate"`
	Assignee    string    `json:"assignee"`
	models.BaseModel
}

func ModelToResponse(issue models.Issue) IssueResponse {
	return IssueResponse{
		ID:          issue.ID,
		TypeId:      issue.TypeId,
		Title:       issue.Title,
		Description: issue.Description,
		StartDate:   issue.StartDate,
		DueDate:     issue.DueDate,
		Assignee:    issue.Assignee,
		BaseModel:   issue.BaseModel,
	}
}
