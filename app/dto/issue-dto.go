package dto

import (
	"github.com/Savitar465/task-manager/app/models"
)

type IssueResponse struct {
	ID          int    `json:"id"`
	TypeId      string `json:"typeId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	DueDate     string `json:"dueDate"`
	StageId     string `json:"stageId"`
	BoardId     string `json:"boardId"`
	Assignee    string `json:"assignee"`
	models.BaseModel
}

type IssueRequest struct {
	TypeId      string `json:"typeId" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	StartDate   string `json:"startDate" binding:"required"`
	DueDate     string `json:"dueDate" binding:"required"`
	StageId     string `json:"stageId" binding:"required"`
	BoardId     string `json:"boardId" binding:"required"`
	Assignee    string `json:"assignee" binding:"required"`
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
