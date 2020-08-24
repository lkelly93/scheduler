package runners_test

import (
	"os"
	"testing"

	"github.com/lkelly93/code-runner/runners"
)

func TestPythonCreateFile(t *testing.T) {
	runnerFileFunctor := runners.GetFunctor("python")

	sysCommand, fileLocation := runnerFileFunctor("print('Hello World')")

	actual := sysCommand + " " + fileLocation
	expected := "python3 ../bin/runner_files/PythonRunner.py"

	assertEquals(expected, actual, t)

	os.Remove(fileLocation)
}

func TestJavaCrateFile(t *testing.T) {
	runnerFileFunctor := runners.GetFunctor("java")

	sysCommand, fileLocation := runnerFileFunctor("public static void main(String[] args){System.out.println(\"Hello World\");}")

	actual := sysCommand + " " + fileLocation
	expected := "java ../bin/runner_files/JavaRunner.java"

	assertEquals(expected, actual, t)

	os.Remove(fileLocation)
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
