package repository

import (
	"context"
	"github.com/Savitar465/task-manager/app/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IssueRepository interface {
	Save(issue *models.Issue) (models.Issue, error)
	FindIssueById(issueId string) models.Issue
	FindAllIssues(ctx context.Context) ([]models.Issue, error)
	DeleteUserById() (string, error)
}

type IssueRepositoryImpl struct {
	db *gorm.DB
}

func IssueRepositoryInit(db *gorm.DB) *IssueRepositoryImpl {
	return &IssueRepositoryImpl{
		db: db,
	}
}

func (i IssueRepositoryImpl) FindAllIssues(ctx context.Context) ([]models.Issue, error) {
	var issues []models.Issue
	var err = i.db.WithContext(ctx).Find(&issues).Error
	if err != nil {
		return nil, err
	}
	return issues, nil
}

func (i IssueRepositoryImpl) Save(issue *models.Issue) (models.Issue, error) {
	var err = i.db.Save(issue).Error
	if err != nil {
		log.Error("Error when save data. Error: ", err)
		return models.Issue{}, err
	}
	return *issue, nil
}

func (i IssueRepositoryImpl) FindIssueById(issueId string) models.Issue {
	log.Info("start to execute program get data user {}", issueId)
	//var err = i.db.Find(&models.Issue{}, idIssue)
	//if err != nil {
	//	log.Error("Error when get data. Error: ", err)
	//
	//}
	return models.Issue{}
}

func (i IssueRepositoryImpl) DeleteUserById() (string, error) {
	var err = i.db.Delete(&models.Issue{}).Error
	if err != nil {
		log.Error("Error when delete data. Error: ", err)
		return "Error when delete data", err
	}
	return "Issue Deleted", nil
}
