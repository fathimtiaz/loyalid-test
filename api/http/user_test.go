package http

import (
	"context"
	"loyalid-test/application/user"
	"loyalid-test/application/user/mocks"
	"loyalid-test/domain"
	"loyalid-test/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_CurrentUser(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name       string
		h          *UserHandler
		args       args
		initArgs   func(args *args)
		initMocks  func(h *UserHandler)
		wantStatus int
	}{
		{
			h: &UserHandler{},
			initArgs: func(args *args) {
				w := httptest.NewRecorder()
				args.c, _ = gin.CreateTestContext(w)

				request := httptest.NewRequest(http.MethodGet, "/user/current", nil)
				args.c.Request = request

				args.c.Request = args.c.Request.WithContext(context.WithValue(args.c, jwtmiddleware.ContextKey{}, &validator.ValidatedClaims{RegisteredClaims: validator.RegisteredClaims{Subject: "auth0|id"}}))
			},
			initMocks: func(h *UserHandler) {
				userRepo := mocks.NewRepository(t)
				h.UserService = user.New(userRepo)

				userRepo.On("GetUserByAuth0Id", mock.Anything, mock.Anything).Return(domain.User{}, repository.ErrNotFound)
			},
			wantStatus: http.StatusNotFound,
		},
		{
			h: &UserHandler{},
			initArgs: func(args *args) {
				w := httptest.NewRecorder()
				args.c, _ = gin.CreateTestContext(w)

				request := httptest.NewRequest(http.MethodGet, "/user/current", nil)
				args.c.Request = request

				args.c.Request = args.c.Request.WithContext(context.WithValue(args.c, jwtmiddleware.ContextKey{}, &validator.ValidatedClaims{RegisteredClaims: validator.RegisteredClaims{Subject: "auth0|id"}}))
			},
			initMocks: func(h *UserHandler) {
				userRepo := mocks.NewRepository(t)
				h.UserService = user.New(userRepo)

				userRepo.On("GetUserByAuth0Id", mock.Anything, mock.Anything).Return(domain.User{Username: "user"}, nil)
			},
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		tt.initArgs(&tt.args)
		tt.initMocks(tt.h)

		t.Run(tt.name, func(t *testing.T) {
			tt.h.CurrentUser(tt.args.c)
		})

		assert.Equal(t, tt.args.c.Writer.Status(), tt.wantStatus)
	}
}
