package executable_test

import (
	"testing"

	"github.com/lkelly93/scheduler/internal/executable"
)

func TestNewExecutable(t *testing.T) {
	_, err := executable.NewExecutable("python", "print('Hello World')")
	if err != nil {
		t.Error(err)
	}
}

func TestNewExecutableFail(t *testing.T) {
	_, err := executable.NewExecutable("Not a Language", "Not Code")
	if err == nil {
		t.Errorf("This test should of failed but it didn't")
	}
}

/****** Python Tests******/

func TestRunPythonCode(t *testing.T) {
	prog, _ := executable.NewExecutable("python", "print('Hello World')")
	expected := "Hello World\n"
	genericRunCode(prog, expected, t)
}

func TestRunBadPythonCode(t *testing.T) {
	prog, _ := executable.NewExecutable("python", "print('Hi")
	expected := "  File \"PythonRunner.py\", line 1\n" +
		"    print('Hi\n" +
		"            ^\n" +
		"SyntaxError: EOL while scanning string literal\n"

	genericRunBadCode(prog, expected, t)
}

/****** Java Tests******/

func TestRunJavaCode(t *testing.T) {
	prog, _ := executable.NewExecutable("java", "public static void main(String[] args){System.out.println(\"Hello World\");}")
	expected := "Hello World\n"
	genericRunCode(prog, expected, t)
}

func TestRunBadJavaCode(t *testing.T) {
	prog, _ := executable.NewExecutable("java", "public static void main(String[] args){System.out.println(\"Hello World\")")
	expected := "JavaRunner.java:3: error: ';' expected\n" +
		"public static void main(String[] args){System.out.println(\"Hello World\")\n" +
		"                                                                        ^\n" +
		"JavaRunner.java:5: error: reached end of file while parsing\n" +
		"}\n" +
		" ^\n" +
		"2 errors\n" +
		"error: compilation failed\n"

	genericRunBadCode(prog, expected, t)

}

func genericRunCode(prog executable.Executable, expected string, t *testing.T) {
	actual := prog.Run()

	//TODO:Check if the file was properly deleted
	assertEquals(expected, actual, t)
}

func genericRunBadCode(prog executable.Executable, expected string, t *testing.T) {
	actual := prog.Run()
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
