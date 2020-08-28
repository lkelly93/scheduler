package executable

import (
	"os"
	"testing"
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

	createFileFunction := getFileCreationFunction(lang)
	_, fileLocation := createFileFunction(code, nil)
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
	actual := removeFilePath(message, mockFilePath)

	assertEquals(expected, actual, t)
}

// /****** Supporting Methods ******/
func genericCreateFile(lang string, code string, expected string, t *testing.T) {
	createFileFunction := getFileCreationFunction(lang)

	sysCommand, fileLocation := createFileFunction(code, nil)
	defer os.Remove(fileLocation)

	actual := sysCommand + " " + fileLocation

	assertEquals(expected, actual, t)
}
