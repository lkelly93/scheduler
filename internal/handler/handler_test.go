package handler_test

import (
	"os"
	"testing"

	"github.com/lkelly93/scheduler/internal/handler"
)

func TestPythonCreateFile(t *testing.T) {
	lang := "python"
	code := "print('Hello World')"
	expected := "python3 ../runner_files/PythonRunner.py"

	genericCreateFile(lang, code, expected, t)
}

func TestJavaCreateFile(t *testing.T) {
	lang := "java"
	code := "public static void main(String[] args){System.out.println(\"Hello World\");}"
	expected := "java ../runner_files/JavaRunner.java"

	genericCreateFile(lang, code, expected, t)
}

func TestCreateRunnerFile(t *testing.T) {
	lang := "python"
	code := "print('Hello World')"

	createFileFunction := handler.GetFileHandler(lang, nil)
	_, fileLocation := createFileFunction.CreateRunnerFile(code)
	defer os.Remove(fileLocation)

	_, err := os.Stat(fileLocation)
	if err != nil {
		t.Errorf("%s file was not created", fileLocation)
	}
}

func TestRemoveFilePath(t *testing.T) {
	message := "/path/to/runner/file/PythonRunner.py had an error Python();<_aRunner.py"
	expected := "PythonRunner.py had an error Python();<_aRunner.py"
	mockFilePath := "/path/to/runner/file/PythonRunner.py"
	actual := handler.RemoveFilePath(message, mockFilePath)

	assertEquals(expected, actual, t)
}

func genericCreateFile(lang string, code string, expected string, t *testing.T) {
	createFileFunction := handler.GetFileHandler(lang, nil)

	sysCommand, fileLocation := createFileFunction.CreateRunnerFile(code)
	defer os.Remove(fileLocation)

	actual := sysCommand + " " + fileLocation

	assertEquals(expected, actual, t)
}

/****** Supporting Methods ******/
func assertEquals(expected string, actual string, t *testing.T) {
	if actual != expected {
		i := 0
		var expectedChar byte
		var actualChar byte
		for i < len(expected) && i < len(actual) {
			if expected[i] != actual[i] {
				expectedChar = expected[i]
				actualChar = actual[i]
				break
			}
			i++
		}
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
		t.Errorf("Error at index %d, expected %c but was %c", i, expectedChar, actualChar)
	}
}
