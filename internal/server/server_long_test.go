// +build !longTests

package server

import (
    "testing"
    "time"
    "net/http"
)

func TestStartHTTPServer(t *testing.T) {
    srv := NewHTTPServer()
    go srv.Start(8765)
    defer srv.Stop()
    time.Sleep(1 * time.Second)
    _, err := http.Get("http://localhost:8765")
    if (err != nil) {
        t.Errorf("Expected nil err from http.Get but got:  %s", err)
    }
}

func TestStopHTTPServer(t *testing.T) {
    srv := NewHTTPServer()
    go srv.Start(8765)
    time.Sleep(1 * time.Second)
    srv.Stop()
    _, err := http.Get("http://localhost:8765")
    if (err == nil) {
        t.Errorf("Expected connection refused error.")
    }
}
