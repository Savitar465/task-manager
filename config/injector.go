//go:build wireinject
// +build wireinject

package config

import (
	"github.com/Savitar465/task-manager/app/modules"
	"github.com/google/wire"
)

var dbsql = InitDB
var db = wire.NewSet(dbsql)

func InitApp() *Initialization {
	wire.Build(NewInitialization, db, modules.IssueCtrlSet, modules.IssueServiceSet, modules.IssueRepoSet)
	return nil
}
