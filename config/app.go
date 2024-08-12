package config

import (
	"github.com/Savitar465/task-manager/app/controller"
	"github.com/Savitar465/task-manager/app/repository"
	"github.com/Savitar465/task-manager/app/service"
)

type Initialization struct {
	IssueRepo repository.IssueRepository
	IssueSvc  service.IssueService
	IssueCtrl controller.IssueController
}

func NewInitialization(userRepo repository.IssueRepository,
	userService service.IssueService,
	userCtrl controller.IssueController,
) *Initialization {
	return &Initialization{
		IssueRepo: userRepo,
		IssueSvc:  userService,
		IssueCtrl: userCtrl,
	}
}
