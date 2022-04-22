package handler

import (
	"testing"

	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"github.com/stretchr/testify/assert"
)

func TestCanAccess_AccessDenied(t *testing.T) {
	var groupId uint64 = 123
	user := &models.User{
		Groups: []*models.Group{
			{
				ID: groupId,
			},
		},
	}

	database := &model.Database{GroupID: uint(groupId)}

	canAccess := CanAccess(user, database)

	assert.True(t, canAccess)
}

func TestCanAccess_IsMemberOf(t *testing.T) {
	var groupId uint64 = 123
	user := &models.User{
		Groups: []*models.Group{
			{
				ID: groupId,
			},
		},
	}

	database := &model.Database{GroupID: uint(groupId)}

	canAccess := CanAccess(user, database)

	assert.True(t, canAccess)
}

func TestCanAccess_IsGroupAdministrator(t *testing.T) {
	var groupId uint64 = 123
	user := &models.User{
		AdminGroups: []*models.Group{
			{
				ID: groupId,
			},
		},
	}

	database := &model.Database{GroupID: uint(groupId)}

	canAccess := CanAccess(user, database)

	assert.True(t, canAccess)
}

func TestCanAccess_IsAdministrator(t *testing.T) {
	user := &models.User{
		Groups: []*models.Group{
			{
				Name: AdministratorGroupName,
			},
		},
	}

	database := &model.Database{}

	canAccess := CanAccess(user, database)

	assert.True(t, canAccess)
}

func TestIsGroupAdministrator(t *testing.T) {
	var groupId uint64 = 123
	adminGroups := []*models.Group{
		{
			ID: groupId,
		},
	}

	isAdmin := IsGroupAdministrator(uint(groupId), adminGroups)

	assert.True(t, isAdmin)
}
