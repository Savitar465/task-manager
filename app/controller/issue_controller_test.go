package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Savitar465/task-manager/app/constant"
	"github.com/Savitar465/task-manager/app/controller"
	"github.com/Savitar465/task-manager/app/dto"
	"github.com/Savitar465/task-manager/app/models"
	"github.com/Savitar465/task-manager/app/service"
	"github.com/Savitar465/task-manager/app/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockIssueService is a mock implementation of IssueService
type MockIssueService struct {
	GetAllIssuesFunc func(c *gin.Context) ([]models.Issue, error)
	CreateIssueFunc  func(c *gin.Context) models.Issue
	DeleteIssueFunc  func(c *gin.Context, id uint)
}

func (m *MockIssueService) GetAllIssues(c *gin.Context) ([]models.Issue, error) {
	if m.GetAllIssuesFunc != nil {
		return m.GetAllIssuesFunc(c)
	}
	return nil, nil
}

func (m *MockIssueService) CreateIssue(c *gin.Context) models.Issue {
	if m.CreateIssueFunc != nil {
		return m.CreateIssueFunc(c)
	}
	return models.Issue{}
}

func (m *MockIssueService) DeleteIssue(c *gin.Context, id uint) {
	if m.DeleteIssueFunc != nil {
		m.DeleteIssueFunc(c, id)
	}
}

var _ service.IssueService = &MockIssueService{}

func TestIssueControllerImpl_GetAll_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockIssueService{}
	issueController := controller.IssueControllerrInit(mockService)

	expectedIssuesModels := []models.Issue{
		{ID: 1, TypeId: "bug", Title: "Test Bug 1", Description: "A test bug", StageId: "open", BoardId: "proj1", Assignee: "dev1", BaseModel: models.BaseModel{CreatedAt: time.Now().String(), UpdatedAt: time.Now().String(), CreatedBy: "Admin", UpdatedBy: "Admin"}},
		{ID: 2, TypeId: "feature", Title: "Test Feature 1", Description: "A test feature", StageId: "inprogress", BoardId: "proj1", Assignee: "dev2", BaseModel: models.BaseModel{CreatedAt: time.Now().String(), UpdatedAt: time.Now().String(), CreatedBy: "Admin", UpdatedBy: "Admin"}},
	}

	mockService.GetAllIssuesFunc = func(c *gin.Context) ([]models.Issue, error) {
		return expectedIssuesModels, nil
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request, _ = http.NewRequest(http.MethodGet, "/issues", nil)

	issueController.GetAll(c)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response util.ApiResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, constant.Success, response.Status)
	assert.NotEmpty(t, response.Data)

	// Convert response.Data (which is []interface{}) to []dto.IssueResponse for comparison
	var actualIssueResponses []dto.IssueResponse
	dataBytes, _ := json.Marshal(response.Data)
	err = json.Unmarshal(dataBytes, &actualIssueResponses)
	assert.NoError(t, err)

	expectedIssueResponses := []dto.IssueResponse{}
	for _, issueModel := range expectedIssuesModels {
		expectedIssueResponses = append(expectedIssueResponses, dto.ModelToResponse(issueModel))
	}

	assert.Equal(t, len(expectedIssueResponses), len(actualIssueResponses))
	// For a more thorough check, you might iterate and compare individual fields,
	// especially if timestamps or other generated fields are involved.
	// For simplicity here, we're comparing the marshalled DTOs.
	// This requires dto.IssueResponse to have json tags that match the output.
	assert.ElementsMatch(t, expectedIssueResponses, actualIssueResponses)
}

func TestIssueControllerImpl_GetAll_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockIssueService{}
	issueController := controller.IssueControllerrInit(mockService)

	mockService.GetAllIssuesFunc = func(c *gin.Context) ([]models.Issue, error) {
		// Service layer is expected to panic, but if it returned an error, controller handles it.
		// Based on current controller logic, it logs the error and c.JSON is not called.
		return nil, errors.New("service error")
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request, _ = http.NewRequest(http.MethodGet, "/issues", nil)

	issueController.GetAll(c)

	// As per current controller implementation, if service.GetAllIssues returns an error,
	// the error is logged, and no JSON response is written.
	// Gin's default response code when nothing is written is http.StatusOK unless changed.
	assert.Equal(t, http.StatusOK, recorder.Code) // or http.StatusInternalServerError if controller was modified to set it
	assert.Empty(t, recorder.Body.String(), "Response body should be empty when service returns an error and controller doesn't write JSON")

	// If the service layer panics (as it currently does), the PanicHandler in the service
	// would likely write an error response. This test case assumes the service *returns* an error
	// to the controller, and the controller's `if err != nil` block is hit.
	// If GetAllIssues *panics* and PanicHandler handles it by writing a JSON error response,
	// then the assertions here would need to change to expect that JSON error response.
	// The current controller code's `else` means if `err != nil`, `c.JSON` is skipped.
}

func TestIssueControllerImpl_Create_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockIssueService{}
	issueController := controller.IssueControllerrInit(mockService)

	expectedIssueModel := models.Issue{
		ID:          1,
		TypeId:      "story",
		Title:       "New Story",
		Description: "A detailed story.",
		StartDate:   time.Now().Format("2006-01-02"),
		DueDate:     time.Now().AddDate(0, 0, 10).Format("2006-01-02"),
		StageId:     "backlog",
		BoardId:     "proj-gamma",
		Assignee:    "userX",
		BaseModel: models.BaseModel{
			CreatedBy: "Admin",
			UpdatedBy: "Admin",
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		},
	}

	mockService.CreateIssueFunc = func(c *gin.Context) models.Issue {
		// Here you could add assertions to check the input DTO if needed,
		// by binding c.Request.Body to dto.IssueRequest
		return expectedIssueModel
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	requestPayload := dto.IssueRequest{
		TypeId:      expectedIssueModel.TypeId,
		Title:       expectedIssueModel.Title,
		Description: expectedIssueModel.Description,
		StartDate:   expectedIssueModel.StartDate,
		DueDate:     expectedIssueModel.DueDate,
		StageId:     expectedIssueModel.StageId,
		BoardId:     expectedIssueModel.BoardId,
		Assignee:    expectedIssueModel.Assignee,
	}
	jsonPayload, _ := json.Marshal(requestPayload)
	httpRequest, _ := http.NewRequest(http.MethodPost, "/issues", bytes.NewBuffer(jsonPayload))
	httpRequest.Header.Set("Content-Type", "application/json")
	c.Request = httpRequest

	issueController.Create(c)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response util.ApiResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, constant.Success, response.Status)

	var actualIssueResponse dto.IssueResponse
	dataBytes, _ := json.Marshal(response.Data)
	err = json.Unmarshal(dataBytes, &actualIssueResponse)
	assert.NoError(t, err)

	expectedResponseDto := dto.ModelToResponse(expectedIssueModel)
	assert.Equal(t, expectedResponseDto.ID, actualIssueResponse.ID)
	assert.Equal(t, expectedResponseDto.Title, actualIssueResponse.Title)
	assert.Equal(t, expectedResponseDto.Description, actualIssueResponse.Description)
	// Compare other relevant fields as necessary
}

func TestIssueControllerImpl_Create_ServicePanics_InvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockIssueService{}
	issueController := controller.IssueControllerrInit(mockService)

	mockService.CreateIssueFunc = func(c *gin.Context) models.Issue {
		// This simulates the service layer panicking, e.g., due to a binding error it detected.
		// The PanicHandler in the service layer is expected to catch this.
		panic(constant.InvalidRequest)
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	// A valid body is provided, but the service will panic.
	requestPayload := dto.IssueRequest{TypeId: "t", Title: "T", Description: "D", StartDate: "S", DueDate: "D", StageId: "S", BoardId: "B", Assignee: "A"}
	jsonPayload, _ := json.Marshal(requestPayload)
	httpRequest, _ := http.NewRequest(http.MethodPost, "/issues", bytes.NewBuffer(jsonPayload))
	httpRequest.Header.Set("Content-Type", "application/json")
	c.Request = httpRequest

	// The PanicHandler in the service should catch the panic and write an error JSON response.
	// We need to ensure the PanicHandler is actually invoked.
	// Since the controller calls the service method directly, and the service method
	// has `defer util.PanicHandler(c)`, this should work.
	issueController.Create(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	var errorResponse util.ApiResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, constant.InvalidRequest, errorResponse.Status)
	// Optionally, assert the error message if util.BuildErrorResponse includes specific messages
	// For example: assert.Contains(t, errorResponse.Message, "Invalid request data")
}

func TestIssueControllerImpl_Create_ServicePanics_UnknownError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockIssueService{}
	issueController := controller.IssueControllerrInit(mockService)

	mockService.CreateIssueFunc = func(c *gin.Context) models.Issue {
		// Simulates a panic like a database error from the repository within the service.
		panic(constant.UnknownError)
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	requestPayload := dto.IssueRequest{TypeId: "t", Title: "T", Description: "D", StartDate: "S", DueDate: "D", StageId: "S", BoardId: "B", Assignee: "A"}
	jsonPayload, _ := json.Marshal(requestPayload)
	httpRequest, _ := http.NewRequest(http.MethodPost, "/issues", bytes.NewBuffer(jsonPayload))
	httpRequest.Header.Set("Content-Type", "application/json")
	c.Request = httpRequest

	issueController.Create(c)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	var errorResponse util.ApiResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, constant.UnknownError, errorResponse.Status)
}

func TestIssueControllerImpl_Delete_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := &MockIssueService{}
	issueController := controller.IssueControllerrInit(mockService)

	testID := "123"
	var calledID uint
	serviceCalled := false

	mockService.DeleteIssueFunc = func(c *gin.Context, id uint) {
		serviceCalled = true
		calledID = id
		// No panic means success. PanicHandler in service won't write a response.
		// Controller should then write the success response.
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "id", Value: testID}}
	c.Request, _ = http.NewRequest(http.MethodDelete, "/issues/"+testID, nil)

	issueController.Delete(c)

	assert.True(t, serviceCalled, "Service DeleteIssue should be called")
	assert.Equal(t, uint(123), calledID, "Service DeleteIssue called with wrong ID")
	assert.Equal(t, http.StatusOK, recorder.Code)

	var response util.ApiResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, constant.Success, response.Status)
	assert.Nil(t, response.Data, "Data should be nil for successful deletion")
}

func TestIssueControllerImpl_Delete_ServicePanics_UnknownError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := &MockIssueService{}
	issueController := controller.IssueControllerrInit(mockService)

	testID := "456"
	serviceCalled := false

	mockService.DeleteIssueFunc = func(c *gin.Context, id uint) {
		serviceCalled = true
		assert.Equal(t, uint(456), id, "Service DeleteIssue called with wrong ID before panic")
		panic(constant.UnknownError) // Simulate service error (e.g., issue not found, DB error)
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "id", Value: testID}}
	c.Request, _ = http.NewRequest(http.MethodDelete, "/issues/"+testID, nil)

	issueController.Delete(c)

	assert.True(t, serviceCalled, "Service DeleteIssue should be called")
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	var errorResponse util.ApiResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, constant.UnknownError, errorResponse.Status)
}

func TestIssueControllerImpl_Delete_InvalidIdFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := &MockIssueService{} // Service should not be called
	issueController := controller.IssueControllerrInit(mockService)
	serviceCalled := false

	mockService.DeleteIssueFunc = func(c *gin.Context, id uint) {
		serviceCalled = true // This should not happen
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	invalidTestID := "abc"
	c.Params = gin.Params{gin.Param{Key: "id", Value: invalidTestID}}
	c.Request, _ = http.NewRequest(http.MethodDelete, "/issues/"+invalidTestID, nil)

	issueController.Delete(c)

	assert.False(t, serviceCalled, "Service DeleteIssue should not be called with invalid ID format")
	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	var errorResponse util.ApiResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	// Check for a more specific error message if your BuildErrorResponse sets one for this case
	assert.Contains(t, errorResponse.Message, "Invalid ID format", "Error message should indicate invalid ID format")
	// The Status field in BuildErrorResponse might be generic or specific.
	// For this case, it's likely a general "Bad Request" or similar if not customized.
	// The controller code uses util.BuildErrorResponse("Invalid ID format", err.Error(), nil)
	// so the Status field in ApiResponse (which is string) will be "Invalid ID format".
	assert.Equal(t, "Invalid ID format", errorResponse.Status)
}
