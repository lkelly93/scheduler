package program_test

import (
	"testing"

	"github.com/lkelly93/scheduler/internal/program"
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

/****** Python Tests******/

func TestRunPythonCode(t *testing.T) {
	prog, _ := program.NewProgram("python", "print('Hello World')")
	expected := "Hello World\n"
	genericRunCode(prog, expected, t)
}

func TestRunBadPythonCode(t *testing.T) {
	prog, _ := program.NewProgram("python", "print('Hi")
	expected := "  File \"../runner_files/PythonRunner.py\", line 1\n" +
		"    print('Hi\n" +
		"            ^\n" +
		"SyntaxError: EOL while scanning string literal\n"

	genericRunBadCode(prog, expected, t)
}

/****** Java Tests******/

func TestRunJavaCode(t *testing.T) {
	prog, _ := program.NewProgram("java", "public static void main(String[] args){System.out.println(\"Hello World\");}")
	expected := "Hello World\n"
	genericRunCode(prog, expected, t)
}

func TestRunBadJavaCode(t *testing.T) {
	prog, _ := program.NewProgram("java", "public static void main(String[] args){System.out.println(\"Hello World\")}")
	expected := "../runner_files/JavaRunner.java:1: error: ';' expected\n" +
		"import java.util.*;public class JavaRunner{public static void main(String[] args){System.out.println(\"Hello World\")}}\n" +
		"                                                                                                                   ^\n" +
		"1 error\n" +
		"error: compilation failed\n"

	genericRunBadCode(prog, expected, t)

}

func genericRunCode(prog *program.Program, expected string, t *testing.T) {
	actual, err := prog.Run()

	if err != nil {
		t.Fatal(err.Error())
	}

	//TODO:Check if the file was properly deleted
	assertEquals(expected, actual, t)
}

func genericRunBadCode(prog *program.Program, expected string, t *testing.T) {
	actual, err := prog.Run()

	if err == nil {
		t.Fatal("This should of failed and did not")
	}

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
