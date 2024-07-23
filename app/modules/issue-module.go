package modules

import (
	"github.com/Savitar465/task-manager/app/controller"
	"github.com/Savitar465/task-manager/app/repository"
	"github.com/Savitar465/task-manager/app/service"
	"github.com/google/wire"
)

var IssueServiceSet = wire.NewSet(service.IssueServiceInit,
	wire.Bind(new(service.IssueService), new(*service.IssueServiceImpl)),
)

var IssueRepoSet = wire.NewSet(repository.IssueRepositoryInit,
	wire.Bind(new(repository.IssueRepository), new(*repository.IssueRepositoryImpl)),
)

var IssueCtrlSet = wire.NewSet(controller.IssueControllerrInit,
	wire.Bind(new(controller.IssueController), new(*controller.IssueControllerImpl)),
)
