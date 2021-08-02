package server

import (
	"net/http/httptest"
	"testing"
)

func TestNewServer(t *testing.T) {
	srv := httptest.NewServer(NewServer().httpServer.Handler)
	defer srv.Close()

}
