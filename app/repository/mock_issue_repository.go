package repository

import (
	"context"
	"github.com/Savitar465/task-manager/app/models"
)

type MockIssueRepository struct {
	SaveFunc        func(issue *models.Issue) (models.Issue, error)
	FindAllIssuesFunc func(ctx context.Context) ([]models.Issue, error)
	// Add other methods here as needed
	FindIssueByIdFunc func(issueId string) models.Issue
	DeleteByIdFunc    func(id uint) error
	// DeleteUserByIdFunc func() (string, error) // Removed as per new requirement
}

func (m *MockIssueRepository) Save(issue *models.Issue) (models.Issue, error) {
	if m.SaveFunc != nil {
		return m.SaveFunc(issue)
	}
	// Default behavior or return an error if SaveFunc is not set
	return models.Issue{}, nil // Or return an error like: errors.New("SaveFunc not implemented")
}

func (m *MockIssueRepository) FindAllIssues(ctx context.Context) ([]models.Issue, error) {
	if m.FindAllIssuesFunc != nil {
		return m.FindAllIssuesFunc(ctx)
	}
	// Default behavior or return an error if FindAllIssuesFunc is not set
	return nil, nil // Or return an error like: errors.New("FindAllIssuesFunc not implemented")
}

// Implement other methods of the IssueRepository interface here as needed
func (m *MockIssueRepository) FindIssueById(issueId string) models.Issue {
	if m.FindIssueByIdFunc != nil {
		return m.FindIssueByIdFunc(issueId)
	}
	// Default behavior or return an error if FindIssueByIdFunc is not set
	return models.Issue{} // Or return an error like: errors.New("FindIssueByIdFunc not implemented")
}

// func (m *MockIssueRepository) DeleteUserById() (string, error) { // Removed as per new requirement
// 	if m.DeleteUserByIdFunc != nil {
// 		return m.DeleteUserByIdFunc()
// 	}
// 	// Default behavior or return an error if DeleteUserByIdFunc is not set
// 	return "", nil // Or return an error like: errors.New("DeleteUserByIdFunc not implemented")
// }

func (m *MockIssueRepository) DeleteById(id uint) error {
	if m.DeleteByIdFunc != nil {
		return m.DeleteByIdFunc(id)
	}
	// Default behavior or return an error if DeleteByIdFunc is not set
	return nil // Or return an error like: errors.New("DeleteByIdFunc not implemented")
}
