package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"loyalid-test/application/product"
	"loyalid-test/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestProductHandler_CreateProduct(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name       string
		h          *ProductHandler
		args       args
		initArgs   func(args *args)
		initStubs  func(h *ProductHandler)
		wantStatus int
	}{
		{
			name: "nil request body",
			h:    &ProductHandler{},
			initArgs: func(args *args) {
				w := httptest.NewRecorder()
				args.c, _ = gin.CreateTestContext(w)

				request := httptest.NewRequest(http.MethodPost, "/product", nil)

				args.c.Request = request
			},
			initStubs:  func(h *ProductHandler) {},
			wantStatus: http.StatusNotAcceptable,
		},
		{
			name: "arbitrary sql error",
			h:    &ProductHandler{},
			initArgs: func(args *args) {
				w := httptest.NewRecorder()
				args.c, _ = gin.CreateTestContext(w)

				body, _ := json.Marshal(map[string]interface{}{"name": "A", "price": 10})
				request := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(body))

				args.c.Request = request
			},
			initStubs: func(h *ProductHandler) {
				productRepoStub := product.RepoStub{}
				productRepoStub.Err = errors.New("sql error")

				h.ProductService = product.New(productRepoStub)
			},
			wantStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt.initArgs(&tt.args)
		tt.initStubs(tt.h)

		t.Run(tt.name, func(t *testing.T) {
			tt.h.CreateProduct(tt.args.c)
		})

		assert.Equal(t, tt.args.c.Writer.Status(), tt.wantStatus)
	}
}

func TestProductHandler_ListProduct(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name       string
		h          *ProductHandler
		args       args
		initArgs   func(args *args)
		initStubs  func(h *ProductHandler)
		wantStatus int
	}{
		{
			name: "arbitrary sql error",
			h:    &ProductHandler{},
			initArgs: func(args *args) {
				w := httptest.NewRecorder()
				args.c, _ = gin.CreateTestContext(w)

				request := httptest.NewRequest(http.MethodGet, "/product?limit=10&page=1", nil)

				args.c.Request = request
			},
			initStubs: func(h *ProductHandler) {
				productRepoStub := product.RepoStub{}
				productRepoStub.Err = errors.New("sql error")

				h.ProductService = product.New(productRepoStub)
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "repository err not found",
			h:    &ProductHandler{},
			initArgs: func(args *args) {
				w := httptest.NewRecorder()
				args.c, _ = gin.CreateTestContext(w)

				request := httptest.NewRequest(http.MethodGet, "/product?limit=10&page=1", nil)

				args.c.Request = request
			},
			initStubs: func(h *ProductHandler) {
				productRepoStub := product.RepoStub{}
				productRepoStub.Err = repository.ErrNotFound

				h.ProductService = product.New(productRepoStub)
			},
			wantStatus: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		tt.initArgs(&tt.args)
		tt.initStubs(tt.h)

		t.Run(tt.name, func(t *testing.T) {
			tt.h.ListProduct(tt.args.c)
		})

		assert.Equal(t, tt.args.c.Writer.Status(), tt.wantStatus)
	}
}
