package executable

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestPythonCreateFile(t *testing.T) {
	lang := "python"
	code := "print('Hello World')"
	expected := "python3 ../runner_files/PythonRunner.py"

	genericCreateFile(lang, code, expected, t)
}

func TestPythonCreateFileCustomFileSettings(t *testing.T) {
	lang := "python"
	code := "print('Hello World')"
	fileSettings := FileSettings{
		Imports:        "",
		ClassName:      "SillyPythonName",
		TrailingCode:   "",
		FileNamePrefix: "PREFIX",
	}
	expectedRunnerFile := "../runner_files/PREFIXSillyPythonName.py"

	function := getFileCreationFunction(lang)
	_, fileLocation := function(code, &fileSettings)
	os.Remove(fileLocation)

	assertEquals(expectedRunnerFile, fileLocation, t)
}
func TestJavaCreateFileCustomFileSettings(t *testing.T) {
	lang := "java"
	code := "public static void main(String[] args){System.out.println(\"Hello World\");}"
	fileSettings := FileSettings{
		Imports:        "",
		ClassName:      "SillyJavaName",
		TrailingCode:   "",
		FileNamePrefix: "PREFIX",
	}
	expectedRunnerFile := "../runner_files/PREFIXSillyJavaName.java"

	function := getFileCreationFunction(lang)
	_, fileLocation := function(code, &fileSettings)
	os.Remove(fileLocation)

	assertEquals(expectedRunnerFile, fileLocation, t)
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

func TestAllSupportedDefaultSettingsJava(t *testing.T) {
	actual := fileSettingsDefaults["java"]
	expected := FileSettings{
		Imports:        "import java.util.*;",
		ClassName:      "JavaRunner",
		TrailingCode:   "",
		FileNamePrefix: "",
	}

	assertEquals(expected.Imports, actual.Imports, t)
	assertEquals(expected.ClassName, actual.ClassName, t)
	assertEquals(expected.TrailingCode, actual.TrailingCode, t)
	assertEquals(expected.FileNamePrefix, actual.FileNamePrefix, t)
}

func TestAllSupportedDefaultSettingsPython(t *testing.T) {
	actual := fileSettingsDefaults["python"]
	expected := FileSettings{
		Imports:        "import numpy as np",
		ClassName:      "PythonRunner",
		TrailingCode:   "",
		FileNamePrefix: "",
	}

	assertEquals(expected.Imports, actual.Imports, t)
	assertEquals(expected.ClassName, actual.ClassName, t)
	assertEquals(expected.TrailingCode, actual.TrailingCode, t)
	assertEquals(expected.FileNamePrefix, actual.FileNamePrefix, t)
}

func TestCreateFileAndAddCode(t *testing.T) {
	outFile := "Test.txt"
	code := "This is a test\n"

	err := createFileAndAddCode(outFile, code)
	defer os.Remove(outFile)

	if err != nil {
		t.Errorf("Could not create a runner file")
	}

	file, err := ioutil.ReadFile(outFile)
	if err != nil {
		t.Errorf("Couldn't open runner file after it was created.")
	}

	fileText := string(file)

	if fileText != code {
		t.Errorf("Runner file text was not correct.")
	}
}

func TestGetRunnerFileLocation(t *testing.T) {
	suffix := "Test.txt"
	expected := "../runner_files/Test.txt"

	actual := getRunnerFileLocation(suffix)

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
