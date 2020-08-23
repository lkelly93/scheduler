package program_test

import (
	"os"
	"testing"

	"github.com/lkelly93/code-runner/program"
)

func TestCreateProgram(t *testing.T) {
	prog, _ := program.NewProgram("python", "print('Hello World')")
	var expectedLang = "python"
	var expectedCode = "print('Hello World')"
	if prog.Lang != expectedLang {
		t.Errorf("Expected %s but got %s", expectedLang, prog.Lang)
	}
	if prog.Code != expectedCode {
		t.Errorf("Expected %s but got %s", expectedCode, prog.Code)
	}
}

func TestUnsupportedLanguage(t *testing.T) {
	_, actual := program.NewProgram("NotALang", "NotRealCode")
	expected := "NotALang is not a supported language"

	assertEquals(expected, actual.Error(), t)

}

func TestPythonCreateFile(t *testing.T) {
	prog, _ := program.NewProgram("python", "print('Hello World')")
	runnerFileFunctor := program.GetFunctor(prog.Lang)

	sysCommand, fileLocation := runnerFileFunctor(prog)

	actual := sysCommand + " " + fileLocation
	expected := "python3 ../runnerFiles/PythonRunner.py"

	assertEquals(expected, actual, t)

	os.Remove(fileLocation)
}

func TestRunPythonCode(t *testing.T) {
	prog, _ := program.NewProgram("python", "print('Hello World')")
	actual, err := program.Run(prog)
	expected := "Hello World\n"

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	//TODO:Check if the file was properly deleted
	assertEquals(expected, actual, t)
}

func TestRunBadPythonCode(t *testing.T) {
	prog, _ := program.NewProgram("python", "print('Hi")
	actual, err := program.Run(prog)
	expected := "  File \"../runnerFiles/PythonRunner.py\", line 1\n" +
		"    print('Hi\n" +
		"            ^\n" +
		"SyntaxError: EOL while scanning string literal\n"

	if err == nil {
		t.Errorf("This should of failed and did not")
	}

	assertEquals(expected, actual, t)
}

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
