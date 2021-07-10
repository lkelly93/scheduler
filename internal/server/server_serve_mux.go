package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/lkelly93/executable/pkg/executable"
)

const codeCharacterLimit = 80000

type executeHandler struct {
}

type errorResponse struct {
	ErrorType string
	Error     string
}

type executionOutputResponse struct {
	Stdout string
}

func writeJSON(rw http.ResponseWriter, obj interface{}) error {
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = rw.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func writeErrorResponse(rw http.ResponseWriter, code int, errorType string, reason string) {
	rw.WriteHeader(code)
	writeJSON(rw, errorResponse{
		ErrorType: errorType,
		Error:     reason,
	})
}

func invalidRequest(rw http.ResponseWriter, reason string) {
	writeErrorResponse(rw, 400, "Invalid Request", reason)
}

func notFound(rw http.ResponseWriter, path string) {
	writeErrorResponse(rw, 404, "Not Found", path+" not found.")
}

func (executeHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		invalidRequest(rw, "/execute/<language> requires POST.")
		return
	}

	path := strings.Split(req.URL.Path, "/")
	if len(path) > 3 {
		notFound(rw, req.URL.Path)
		return
	}

	if path[2] == "" {
		invalidRequest(rw, "/execute/<language> requires language.")
		return
	}

	language := path[2]
	codeBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("[ERROR] Unable to read HTTP request.")
		rw.WriteHeader(500)
		return
	}
	var msg map[string]interface{}
	err = json.Unmarshal(codeBytes, &msg)
	if err != nil {
		invalidRequest(rw, "Invalid JSON (Could not parse JSON).")
		return
	}
	_, ok := msg["Code"]
	if !ok {
		invalidRequest(rw, "Code is a required field.")
		return
	}

	code, ok := msg["Code"].(string)
	if !ok {
		invalidRequest(rw, "Code must be a string.")
		return
	}
	if len(code) > codeCharacterLimit {
		invalidRequest(rw, fmt.Sprintf("Code character limit of %d chars exceeded.", codeCharacterLimit))
		return
	}

	//exec, err := executable.NewExecutable(language, code, nil)
	exec, err := executable.NewExecutable(language, code, randSeq(16))

	if err != nil {
		writeErrorResponse(rw, 400, "unsupported_language",
			fmt.Sprintf("%s is not a supported language.", language))
		return
	}

	//TODO add exec to queue here...
	result, err := exec.Run()
	if err != nil {
		if _, ok = err.(*executable.RuntimeError); ok {
			writeErrorResponse(rw, 200, "Runtime Error", err.Error())
		} else if _, ok = err.(*executable.CompilationError); ok {
			writeErrorResponse(rw, 200, "Compilation Error", err.Error())
		} else if _, ok = err.(*executable.TimeLimitExceededError); ok {
			writeErrorResponse(rw, 200, "Time Limit Exceeded", err.Error())
		} else {
			writeErrorResponse(rw, 500, "Internal Server Error", "Internal Server Error")
		}
		return
	}
	rw.WriteHeader(200)
	writeJSON(rw, executionOutputResponse{
		Stdout: result.Output,
	})
}

func newServeMux() *http.ServeMux {
	sm := http.NewServeMux()
	sm.Handle("/execute/", executeHandler{})
	return sm
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
