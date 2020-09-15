package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/lkelly93/scheduler/internal/executable"
)

/////////////
// Utility //
/////////////

func TestMain(m *testing.M) {
	executable.Init()
	os.Exit(m.Run())
}

type responseWrapper struct {
	res *http.Response

	pattern string
}

func (rw responseWrapper) expectCode(t *testing.T, code int) responseWrapper {
	t.Helper()
	if rw.res.StatusCode != code {
		t.Errorf("Expected HTTP %d, got HTTP %d", code, rw.res.StatusCode)
	}
	return rw
}

func (rw responseWrapper) expectJSON(t *testing.T, expectedJSON string) responseWrapper {
	t.Helper()
	var expected interface{}
	if json.Unmarshal([]byte(expectedJSON), &expected) != nil {
		t.Fatalf("Could not unmarshal expectedJSON.")
	}
	body, _ := ioutil.ReadAll(rw.res.Body)
	var actual interface{}
	if json.Unmarshal(body, &actual) != nil {
		t.Fatalf("Response body is not valid JSON.")
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	return rw
}

func (rw responseWrapper) expectErrorResponse(t *testing.T, errorType string, reason string) responseWrapper {
	t.Helper()
	expected := errorResponse{
		ErrorType: errorType,
		Error:     reason,
	}
	expectedJSON, _ := json.Marshal(expected)
	return rw.expectJSON(t, string(expectedJSON))
}

func getsm(sm *http.ServeMux, target string) responseWrapper {
	req := httptest.NewRequest("GET", target, nil)
	handler, pattern := sm.Handler(req)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return responseWrapper{res: rr.Result(), pattern: pattern}
}

func get(target string) responseWrapper {
	sm := newServeMux()
	return getsm(sm, target)
}

func postsm(sm *http.ServeMux, target string, body string) responseWrapper {
	req := httptest.NewRequest("POST", target, strings.NewReader(body))
	handler, pattern := sm.Handler(req)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return responseWrapper{res: rr.Result(), pattern: pattern}
}

func post(target string, body string) responseWrapper {
	sm := newServeMux()
	return postsm(sm, target, body)
}

///////////
// Tests //
///////////

func Test_Execute_Language(t *testing.T) {
	t.Run("GET /execute/python returns HTTP 400", func(t *testing.T) {
		get("/execute/python").
			expectCode(t, 400)
		t.Helper()
	})

	t.Run("GET /execute/unsupportedLanguage", func(t *testing.T) {
		get("/execute/unsupportedLanguage").
			expectCode(t, 400)
	})

	t.Run("POST /execute/unsupportedLanguage", func(t *testing.T) {
		post("/execute/unsupportedLanguage", `{ "Code": "print(\"Hello World\")" }`).
			expectJSON(t, `{"ErrorType":"unsupported_language",
	                    "Error": "unsupportedLanguage is not a supported language."}`)
	})

	t.Run("POST /execute/differentUnsupportedLanguage", func(t *testing.T) {
		post("/execute/differentUnsupportedLanguage", `{ "Code": "print(\"Hello World\")" }`).
			expectJSON(t, `{"ErrorType":"unsupported_language",
	                    "Error": "differentUnsupportedLanguage is not a supported language."}`)
	})

	t.Run("POST /execute/python with correct JSON", func(t *testing.T) {
		post("/execute/python", `{ "Code": "print(\"Hello World\")" }`).
			expectCode(t, 200).
			expectJSON(t, `{"Stdout": "Hello World\n"}`)
	})

	t.Run("POST /execute/python with valid but incorrect JSON", func(t *testing.T) {
		post("/execute/python", `{}`).
			expectCode(t, 400)
	})

	t.Run("POST /execute/python with invalid JSON", func(t *testing.T) {
		post("/execute/python", `}`).
			expectCode(t, 400)
	})

	t.Run("POST /execute/python/extraneous", func(t *testing.T) {
		post("/execute/python/extraneous", `{}`).
			expectCode(t, 404)
	})

	t.Run("POST /execute/", func(t *testing.T) {
		post("/execute/", `{}`).
			expectCode(t, 400).
			expectErrorResponse(t, "Invalid Request", "/execute/<language> requires language.")
	})

	t.Run("POST /execute/python non-string Code", func(t *testing.T) {
		post("/execute/python", `{"Code": 123}`).
			expectCode(t, 400).
			expectErrorResponse(t, "Invalid Request", "Code must be a string.")
	})

	t.Run("POST /execute/python code reaches character limit", func(t *testing.T) {
		before := "print(\\\""
		after := "\\\")"
		var b strings.Builder
		// add 2 since escape characters get dropped
		for i := 0; i < codeCharacterLimit-len(before)-len(after)+2; i++ {
			fmt.Fprintf(&b, "-")
		}
		fmt.Println(codeCharacterLimit)
		code := fmt.Sprint(before, b.String(), after)
		post("/execute/python", fmt.Sprintf(`{"Code": "%s"}`, code)).
			expectCode(t, 200).
			expectJSON(t, fmt.Sprintf(`{"Stdout": "%s\n"}`, b.String()))
	})

	t.Run("POST /execute/python exceeding code character limit", func(t *testing.T) {
		var b strings.Builder
		before := "print(\\\""
		after := "\\\")"
		fmt.Fprintf(&b, before)
		// add 2 since escape characters get dropped
		for i := 0; i < codeCharacterLimit-len(before)-len(after)+3; i++ {
			fmt.Fprintf(&b, "-")
		}
		fmt.Fprintf(&b, after)
		code := b.String()
		//fmt.Printf(`{"Code": "%s"}`, code)
		post("/execute/python", fmt.Sprintf(`{"Code": "%s"}`, code)).
			expectCode(t, 400).
			expectErrorResponse(t, "Invalid Request",
				fmt.Sprintf("Code character limit of %d chars exceeded.", codeCharacterLimit))
	})

	t.Run("POST /execute/python runtime error", func(t *testing.T) {
		post("/execute/python", `{"Code": "fail" }`).
			expectCode(t, 200).
			expectErrorResponse(t, "Runtime Error",
				"Traceback (most recent call last):\n"+
					"  File \"PythonRunner.py\", line 2, in <module>\n"+
					"    fail\n"+
					"NameError: name 'fail' is not defined\n")
	})

}
