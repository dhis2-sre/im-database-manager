package handler

import (
	"testing"

	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"github.com/stretchr/testify/assert"
)

func TestCanAccess_AccessDenied(t *testing.T) {
	var group = "123"
	user := &models.User{
		Groups: []*models.Group{
			{
				Name: group,
			},
		},
	}

	database := &model.Database{GroupName: group}

	canAccess := CanAccess(user, database)

	assert.True(t, canAccess)
}

func TestCanAccess_IsMemberOf(t *testing.T) {
	var group = "123"
	user := &models.User{
		Groups: []*models.Group{
			{
				Name: group,
			},
		},
	}

	database := &model.Database{GroupName: group}

	canAccess := CanAccess(user, database)

	assert.True(t, canAccess)
}

func TestCanAccess_IsGroupAdministrator(t *testing.T) {
	var group = "123"
	user := &models.User{
		AdminGroups: []*models.Group{
			{
				Name: group,
			},
		},
	}

	database := &model.Database{GroupName: group}

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
	var group = "123"
	adminGroups := []*models.Group{
		{
			Name: group,
		},
	}

	isAdmin := IsGroupAdministrator(group, adminGroups)

	assert.True(t, isAdmin)
}
