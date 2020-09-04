// +build !longTests

package server

import "testing"

func Test_Execute_Language_Long(t *testing.T) {
    t.Run("POST /execute/python timeout error", func(t *testing.T) {
        post("/execute/python", `{"Code": "while True:\n    continue"}`).
        expectCode(t, 200).
        expectErrorResponse(t, "Time Limit Exceeded", "Time Limit Exceeded 15s")
    })
}
