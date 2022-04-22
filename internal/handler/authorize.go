package handler

import (
	"sort"

	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
)

const AdministratorGroupName = "administrators"

func CanAccess(user *models.User, database *model.Database) bool {
	return isAdministrator(user) ||
		IsGroupAdministrator(database.GroupID, user.AdminGroups) ||
		isMemberOf(database.GroupID, user.Groups)
}

func isAdministrator(user *models.User) bool {
	return contains(AdministratorGroupName, user.Groups)
}

func contains(groupName string, groups []*models.Group) bool {
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Name <= groups[j].Name
	})

	index := sort.Search(len(groups), func(i int) bool {
		return groups[i].Name >= groupName
	})

	return index < len(groups) && groups[index].Name == groupName
}

func IsGroupAdministrator(groupId uint, groups []*models.Group) bool {
	return isMemberOf(groupId, groups)
}

func isMemberOf(groupId uint, groups []*models.Group) bool {
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].ID <= groups[j].ID
	})

	index := sort.Search(len(groups), func(i int) bool {
		return uint(groups[i].ID) >= groupId
	})

	return index < len(groups) && uint(groups[index].ID) == groupId
}
