package handler

import (
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
)

const AdministratorGroupName = "administrators"

func CanAccess(user *models.User, database *model.Database) bool {
	return isAdministrator(user) ||
		IsGroupAdministrator(database.GroupName, user.AdminGroups) ||
		isMemberOf(database.GroupName, user.Groups)
}

func isAdministrator(user *models.User) bool {
	return isMemberOf(AdministratorGroupName, user.Groups)
}

func IsGroupAdministrator(groupName string, groups []*models.Group) bool {
	return isMemberOf(groupName, groups)
}

func isMemberOf(groupName string, groups []*models.Group) bool {
	for _, group := range groups {
		if groupName == group.Name {
			return true
		}
	}
	return false
}
