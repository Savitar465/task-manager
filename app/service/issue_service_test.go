package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Savitar465/task-manager/app/constant"
	"github.com/Savitar465/task-manager/app/dto"
	"github.com/Savitar465/task-manager/app/models"
	"github.com/Savitar465/task-manager/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIssueServiceImpl_CreateIssue_Success(t *testing.T) {
	// Initialize Gin for testing
	gin.SetMode(gin.TestMode)

	// Setup Mock Repository
	mockRepo := &repository.MockIssueRepository{}
	expectedIssue := models.Issue{
		ID:          1,
		TypeId:      "task",
		Title:       "Test Issue",
		Description: "This is a test issue",
		StartDate:   time.Now().String(),
		DueDate:     time.Now().AddDate(0, 0, 7).String(),
		StageId:     "todo",
		BoardId:     "project-alpha",
		Assignee:    "user1",
		BaseModel: models.BaseModel{
			CreatedBy: "Admin",
			UpdatedBy: "Admin",
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		},
	}

	mockRepo.SaveFunc = func(issue *models.Issue) (models.Issue, error) {
		// We are assigning the ID and timestamps here to simulate the DB behavior
		// In a real scenario, the DB would generate the ID and timestamps.
		// For the purpose of this test, we ensure the returned issue from the mock
		// matches what we expect CreateIssue to return.
		// The important part is that the input `issue` to `SaveFunc` has the correct data from the request.
		assert.Equal(t, expectedIssue.TypeId, issue.TypeId)
		assert.Equal(t, expectedIssue.Title, issue.Title)
		assert.Equal(t, expectedIssue.Description, issue.Description)
		// ... other assertions for fields passed to Save ...
		return expectedIssue, nil
	}

	// Initialize Service with Mock Repo
	issueService := IssueServiceInit(mockRepo)

	// Create Test HTTP Context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Prepare Request Body
	requestBody := dto.IssueRequest{
		TypeId:      expectedIssue.TypeId,
		Title:       expectedIssue.Title,
		Description: expectedIssue.Description,
		StartDate:   expectedIssue.StartDate,
		DueDate:     expectedIssue.DueDate,
		StageId:     expectedIssue.StageId,
		BoardId:     expectedIssue.BoardId,
		Assignee:    expectedIssue.Assignee,
	}
	jsonBody, _ := json.Marshal(requestBody)
	c.Request, _ = http.NewRequest(http.MethodPost, "/issues", bytes.NewBuffer(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the method
	var result models.Issue
	assert.NotPanics(t, func() {
		result = issueService.CreateIssue(c)
	})

	// Assertions
	assert.Equal(t, expectedIssue.ID, result.ID)
	assert.Equal(t, expectedIssue.Title, result.Title)
	// Add more assertions if necessary for other fields
}

func TestIssueServiceImpl_CreateIssue_BindingError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := &repository.MockIssueRepository{} // Save should not be called
	issueService := IssueServiceInit(mockRepo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Prepare Invalid Request Body (e.g., missing required fields or malformed JSON)
	// For this test, let's send an empty JSON object, which should fail "required" bindings
	jsonBody, _ := json.Marshal(map[string]string{})
	c.Request, _ = http.NewRequest(http.MethodPost, "/issues", bytes.NewBuffer(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	recoveredPanic := ""
	func() {
		defer func() {
			if r := recover(); r != nil {
				recoveredPanic = r.(string) // Assuming PanicException passes a string
			}
		}()
		issueService.CreateIssue(c)
	}()

	assert.Equal(t, constant.InvalidRequest, recoveredPanic)
}

func TestIssueServiceImpl_CreateIssue_RepositorySaveError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := &repository.MockIssueRepository{}
	expectedError := errors.New("database error")

	mockRepo.SaveFunc = func(issue *models.Issue) (models.Issue, error) {
		return models.Issue{}, expectedError
	}

	issueService := IssueServiceInit(mockRepo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	requestBody := dto.IssueRequest{
		TypeId:      "task",
		Title:       "Test Issue",
		Description: "This is a test issue",
		StartDate:   time.Now().String(),
		DueDate:     time.Now().AddDate(0, 0, 7).String(),
		StageId:     "todo",
		BoardId:     "project-alpha",
		Assignee:    "user1",
	}
	jsonBody, _ := json.Marshal(requestBody)
	c.Request, _ = http.NewRequest(http.MethodPost, "/issues", bytes.NewBuffer(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	recoveredPanic := ""
	func() {
		defer func() {
			if r := recover(); r != nil {
				recoveredPanic = r.(string) // Assuming PanicException passes a string
			}
		}()
		issueService.CreateIssue(c)
	}()

	assert.Equal(t, constant.UnknownError, recoveredPanic)
}

func TestIssueServiceImpl_GetAllIssues_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := &repository.MockIssueRepository{}
	expectedIssues := []models.Issue{
		{ID: 1, Title: "Issue 1", Description: "First issue"},
		{ID: 2, Title: "Issue 2", Description: "Second issue"},
	}

	mockRepo.FindAllIssuesFunc = func(ctx context.Context) ([]models.Issue, error) {
		return expectedIssues, nil
	}

	issueService := IssueServiceInit(mockRepo)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/issues", nil) // Context for FindAllIssues

	var result []models.Issue
	var err error
	assert.NotPanics(t, func() {
		result, err = issueService.GetAllIssues(c)
	})

	assert.NoError(t, err)
	assert.Equal(t, expectedIssues, result)
}

func TestIssueServiceImpl_GetAllIssues_EmptyList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := &repository.MockIssueRepository{}
	expectedIssues := []models.Issue{} // Empty slice

	mockRepo.FindAllIssuesFunc = func(ctx context.Context) ([]models.Issue, error) {
		return expectedIssues, nil
	}

	issueService := IssueServiceInit(mockRepo)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/issues", nil)

	var result []models.Issue
	var err error
	assert.NotPanics(t, func() {
		result, err = issueService.GetAllIssues(c)
	})

	assert.NoError(t, err)
	assert.Empty(t, result)
	assert.Equal(t, expectedIssues, result) // Ensure it's an empty slice not nil
}

func TestIssueServiceImpl_GetAllIssues_RepositoryError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := &repository.MockIssueRepository{}
	expectedError := errors.New("database query failed")

	mockRepo.FindAllIssuesFunc = func(ctx context.Context) ([]models.Issue, error) {
		return nil, expectedError
	}

	issueService := IssueServiceInit(mockRepo)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/issues", nil)

	recoveredPanic := ""
	assert.PanicsWithValue(t, constant.UnknownError, func() {
		_, _ = issueService.GetAllIssues(c) // Error is returned, but panic is expected by PanicHandler
	}, "Expected GetAllIssues to panic with UnknownError")

	// Verify the panic value if needed, though PanicsWithValue already does this.
	// This defer func is more for illustration or if specific checks on the panic value are needed beyond equality.
	func() {
		defer func() {
			if r := recover(); r != nil {
				recoveredPanic = r.(string)
			}
		}()
		_, _ = issueService.GetAllIssues(c)
	}()
	// If the function did not panic and instead returned an error, this would be an alternative check:
	// assert.Error(t, err)
	// assert.Nil(t, result)
	// For panic:
	if recoveredPanic != "" { // This check is only useful if PanicsWithValue was not used.
		assert.Equal(t, constant.UnknownError, recoveredPanic)
	}
}

func TestIssueServiceImpl_DeleteIssue_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := &repository.MockIssueRepository{}
	issueService := IssueServiceInit(mockRepo)

	testID := uint(1)
	var calledID uint

	mockRepo.DeleteByIdFunc = func(id uint) error {
		calledID = id
		return nil // Success
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	assert.NotPanics(t, func() {
		issueService.DeleteIssue(c, testID)
	})

	assert.Equal(t, testID, calledID, "DeleteById should be called with the correct ID")
}

func TestIssueServiceImpl_DeleteIssue_RepositoryError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := &repository.MockIssueRepository{}
	issueService := IssueServiceInit(mockRepo)

	testID := uint(2)
	expectedError := errors.New("deletion failed")

	mockRepo.DeleteByIdFunc = func(id uint) error {
		return expectedError
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	assert.PanicsWithValue(t, constant.UnknownError, func() {
		issueService.DeleteIssue(c, testID)
	}, "Expected DeleteIssue to panic with UnknownError on repository error")
}
