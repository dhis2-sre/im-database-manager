package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestRequest struct {
	OneOf string `binding:"required,oneOf=one of"`
}

func TestRegisterValidation(t *testing.T) {
	err := RegisterValidation()
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	dummyRequest, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	ctx.Request = dummyRequest

	err = ctx.ShouldBind(&TestRequest{OneOf: "one"})
	assert.NoError(t, err)

	err = ctx.ShouldBind(&TestRequest{OneOf: "of"})
	assert.NoError(t, err)

	err = ctx.ShouldBind(&TestRequest{OneOf: "oh no"})
	assert.Error(t, err)
	assert.Equal(t, "Key: 'TestRequest.OneOf' Error:Field validation for 'OneOf' failed on the 'oneOf' tag", err.Error())
}
