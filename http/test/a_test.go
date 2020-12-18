package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestHandlerxxx(t *testing.T) {
	handler := initHandler()

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	t.Run("ping", func(t *testing.T) {
		{
			e.Builder(func(req *httpexpect.Request) {
				req.WithHeader("X-Auth-Id", "test")
			}).GET("/ping").
				Expect().
				Status(http.StatusOK).ContentType("text/plain").Text().Equal("pong")
		}
		{
			e.GET("/ping").
				Expect().
				Status(http.StatusUnauthorized)
		}
		{
			e.POST("/ping").
				Expect().
				Status(http.StatusMethodNotAllowed)
		}
	})
	t.Run("user", func(t *testing.T) {
		{
			e.GET("/user").
				Expect().
				Status(http.StatusUnauthorized).Body().NotEmpty().Equal("\n")
		}
		{
			e.Builder(func(req *httpexpect.Request) {
				req.WithHeader("X-Auth-Id", "test")
			}).GET("/user").
				Expect().
				Status(http.StatusNotFound)
		}
		{
			e.Builder(func(req *httpexpect.Request) {
				req.WithHeader("X-Auth-Id", "test")
			}).GET("/user/111").
				Expect().
				Status(http.StatusNotFound)
		}
		{
			raw := e.Builder(func(req *httpexpect.Request) {
				req.WithHeader("X-Auth-Id", "test")
			}).GET("/user/aaa").
				Expect().
				Status(http.StatusOK).ContentType("application/json").JSON().Object()
			raw.ContainsMap(map[string]interface{}{
				"name": "aaa",
			})
		}
	})
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func setup() {
}

func teardown() {
}
