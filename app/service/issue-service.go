package service

import (
	"fmt"
	"github.com/Savitar465/task-manager/app/constant"
	"github.com/Savitar465/task-manager/app/dto"
	"github.com/Savitar465/task-manager/app/models"
	"github.com/Savitar465/task-manager/app/repository"
	"github.com/Savitar465/task-manager/app/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

type IssueService interface {
	CreateIssue(c *gin.Context) models.Issue
	GetAllIssues(c *gin.Context) ([]models.Issue, error)
	DeleteIssue()
}

type IssueServiceImpl struct {
	issueRepo repository.IssueRepository
}

func IssueServiceInit(userRepository repository.IssueRepository) *IssueServiceImpl {
	return &IssueServiceImpl{
		issueRepo: userRepository,
	}
}

func (i IssueServiceImpl) CreateIssue(c *gin.Context) models.Issue {
	defer util.PanicHandler(c)
	log.Info("start to execute program add data user")
	var request dto.IssueRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Error when binding data. Error: ", err)
		util.PanicException(constant.InvalidRequest)
	}
	var issue models.Issue
	issue.TypeId = request.TypeId
	issue.Title = request.Title
	issue.Description = request.Description
	issue.StartDate = request.StartDate
	issue.DueDate = request.DueDate
	issue.StageId = request.StageId
	issue.BoardId = request.BoardId
	issue.Assignee = request.Assignee
	issue.CreatedBy = "Admin"
	issue.UpdatedBy = "Admin"
	issue.CreatedAt = time.Now().String()
	issue.UpdatedAt = time.Now().String()
	data, err := i.issueRepo.Save(&issue)
	if err != nil {
		log.Error("Happened error when save data to database. Error: ", err)
		util.PanicException(constant.UnknownError)
	}
	return data
}

func (i IssueServiceImpl) GetAllIssues(c *gin.Context) ([]models.Issue, error) {
	defer util.PanicHandler(c)
	data, err := i.issueRepo.FindAllIssues(c.Request.Context())
	if err != nil {
		log.Error("Happened error when get data from database. Error: ", err)
		util.PanicException(constant.UnknownError)
		return nil, err
	}
	return data, nil
}

func (i IssueServiceImpl) DeleteIssue() {
	fmt.Println("Hi Delete")
}
