package handler

import (
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"sort"
)

const AdministratorGroupName = "administrators"

func CanAccess(user *models.User, database *model.Database) bool {
	return isAdministrator(user) || IsGroupAdministrator(user.AdminGroups, database.GroupID) || isMemberOfById(user.Groups, database.GroupID)
}

func isMemberOfById(groups []*models.Group, groupId uint) bool {
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].ID <= groups[j].ID
	})

	index := sort.Search(len(groups), func(i int) bool {
		return uint(groups[i].ID) >= groupId
	})

	return index < len(groups) && uint(groups[index].ID) == groupId
}

func isAdministrator(user *models.User) bool {
	return isMemberOf(user, AdministratorGroupName)
}

func isMemberOf(user *models.User, groupName string) bool {
	return contains(groupName, user.Groups)
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

func IsGroupAdministrator(groups []*models.Group, groupId uint) bool {
	return isMemberOfById(groups, groupId)
}
