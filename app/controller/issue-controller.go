package controller

import (
	"github.com/Savitar465/task-manager/app/constant"
	"github.com/Savitar465/task-manager/app/dto"
	issueService "github.com/Savitar465/task-manager/app/service"
	"github.com/Savitar465/task-manager/app/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type IssueController interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type IssueControllerImpl struct {
	service issueService.IssueService
}

// GetAll ShowIssues godoc
//
//	@Summary		Get all issues
//	@Description	get all issues
//	@Tags			issues
//	@Produce		json
//	@Success		200	{array} dto.IssueResponse
//	@Router			/issues [get]
func (h *IssueControllerImpl) GetAll(c *gin.Context) {
	issues, err := h.service.GetAllIssues(c)
	if err != nil {
		log.Error("Error when get data. Error: ", err)
	} else {
		var issuesResponse []dto.IssueResponse
		for _, issue := range issues {
			issuesResponse = append(issuesResponse, dto.ModelToResponse(issue))
		}
		c.JSON(http.StatusOK, util.BuildResponse(constant.Success, issuesResponse))
	}
}

// Create Post ShowIssues godoc
//
// @Summary      Create an Issue
// @Description  Create a new issue using the provided IssueRequest
// @Tags         issues
// @Accept       json
// @Produce      json
// @Param        issue  body  dto.IssueRequest  true  "Issue Request"
// @Success      200    {object}  dto.IssueResponse  "Created issue"
// @Router       /issues [post]
func (h *IssueControllerImpl) Create(c *gin.Context) {
	issue := h.service.CreateIssue(c)
	c.JSON(http.StatusOK, util.BuildResponse(constant.Success, dto.ModelToResponse(issue)))
}

// Delete removes an issue.
//
// @Summary      Delete an Issue
// @Description  Delete an issue by its ID
// @Tags         issues
// @Produce      json
// @Param        id  path  int  true  "Issue ID"
// @Success      200 "Issue deleted"
// @Router       /issues/{id} [delete]
func (h *IssueControllerImpl) Delete(c *gin.Context) {

	c.Status(http.StatusOK)
}

func IssueControllerrInit(issueSrv issueService.IssueService) *IssueControllerImpl {
	return &IssueControllerImpl{
		service: issueSrv,
	}
}
