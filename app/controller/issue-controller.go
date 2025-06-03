package controller

import (
	"github.com/Savitar465/task-manager/app/constant"
	"github.com/Savitar465/task-manager/app/dto"
	issueService "github.com/Savitar465/task-manager/app/service"
	"github.com/Savitar465/task-manager/app/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
// @Success      200 {object} util.ApiResponse "Issue deleted successfully"
// @Failure      400 {object} util.ApiResponse "Invalid ID format"
// @Failure      500 {object} util.ApiResponse "Internal server error"
// @Router       /issues/{id} [delete]
func (h *IssueControllerImpl) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Assuming ID fits in uint32
	if err != nil {
		log.Error("Invalid ID format: ", err)
		c.JSON(http.StatusBadRequest, util.BuildErrorResponse("Invalid ID format", err.Error(), nil))
		return
	}

	// The service method DeleteIssue has its own PanicHandler for service/repository errors.
	// If DeleteIssue panics, its PanicHandler will write the response.
	// If it completes without panic, then this controller code will execute.
	h.service.DeleteIssue(c, uint(id))

	// If service.DeleteIssue panics, PanicHandler in service writes response, and this line might not be reached
	// or Gin's recovery might have already finalized the response.
	// However, if DeleteIssue is successful (no panic), we should confirm the response.
	// Check if response has already been written by PanicHandler.
	// If service.DeleteIssue calls PanicHandler which calls c.JSON, then c.Writer.Written() will be true.
	if !c.Writer.Written() {
		c.JSON(http.StatusOK, util.BuildResponse(constant.Success, nil))
	}
}

func IssueControllerrInit(issueSrv issueService.IssueService) *IssueControllerImpl {
	return &IssueControllerImpl{
		service: issueSrv,
	}
}
