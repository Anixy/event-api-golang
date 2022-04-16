package controllers

import (
	"testing"

	"github.com/gin-gonic/gin"
)


func TestUserControllerImpl_Register(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name           string
		userController UserController
		args           args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.userController.Register(tt.args.c)
		})
	}
}
