package executable

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestNewExecutable(t *testing.T) {
	lang := "python"
	code := "print('Hello World')"
	exe, err := NewExecutable(lang, code, nil)
	if err != nil {
		t.Error(err)
	}

	//Cast Executable interface to state struct
	state := exe.(*executableState)

	assertEquals(code, state.code, t)
}

func TestNewExecutableFail(t *testing.T) {
	lang := "Not a Language"
	_, err := NewExecutable(lang, "Not Code", nil)
	if err == nil {
		t.Errorf("\"%s\" was accepted as a language and should not of been.", lang)
	}
}

// /***** Test Good Runs *****/
func TestRunPythonCode(t *testing.T) {
	prog, _ := NewExecutable("python", "print('Hello World')", nil)
	expected := "Hello World\n"
	genericRunCode(prog, expected, t)
}

func TestRunPythonCodeLonger(t *testing.T) {
	fileLocation := "test_data/longPythonCode.py"
	longCodeFile, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		t.Errorf("Could not read in %s", fileLocation)
	}
	code := string(longCodeFile)
	prog, _ := NewExecutable("python", code, nil)
	expected := "Male\n"
	genericRunCode(prog, expected, t)
}

func TestRunJavaCode(t *testing.T) {
	code := "public static void main(String[] args){System.out.println(\"Hello World\");}"
	prog, _ := NewExecutable("java", code, nil)
	expected := "Hello World\n"
	genericRunCode(prog, expected, t)
}

func TestRunJavaCodeLonger(t *testing.T) {
	fileLocation := "test_data/longJavaCode.java"
	longCode, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		t.Errorf("Could not read in %s", fileLocation)
	}
	code := string(longCode)
	prog, _ := NewExecutable("java", code, nil)
	var expected strings.Builder
	expected.WriteString("NonRecursive\n")
	expected.WriteString("[0, 1, 0, 0, 1, 0, 1, 0]\n")
	expected.WriteString("[0, 0, 0, 0, 0, 1, 1, 0]\n")
	genericRunCode(prog, expected.String(), t)
}

func TestRecursion(t *testing.T) {
	fileLocation := "test_data/recursiveCode.java"
	longCode, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		t.Errorf("Could not read in %s", fileLocation)
	}
	code := string(longCode)
	prog, _ := NewExecutable("java", code, nil)
	var expected strings.Builder
	expected.WriteString("Recursive\n")
	expected.WriteString("[0, 1, 0, 0, 1, 0, 1, 0]\n")
	expected.WriteString("[0, 0, 0, 0, 0, 1, 1, 0]\n")
	genericRunCode(prog, expected.String(), t)

}

func TestFileIsDeletedAfter(t *testing.T) {
	prog, _ := NewExecutable("python", "print('Hello World')", nil)
	fileLocation := "../runner_files/PythonRunner.py"
	_, err := os.Stat(fileLocation)
	if err == nil {
		t.Fatalf("%s existed before Run() was called", fileLocation)
	}

	prog.Run()

	_, err = os.Stat(fileLocation)
	if err == nil {
		t.Errorf("%s still exist after Run() was called. It should of been deleted", fileLocation)
	}
}

/***** Test Bad Runs*****/
func TestRunBadJavaCode(t *testing.T) {
	code := "public static void main(String[] args){System.out.println(\"Hello World\")"
	prog, _ := NewExecutable("java", code, nil)
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

func TestRunBadPythonCode(t *testing.T) {
	prog, _ := NewExecutable("python", "print('Hi", nil)
	expected := "  File \"PythonRunner.py\", line 2\n" +
		"    print('Hi\n" +
		"            ^\n" +
		"SyntaxError: EOL while scanning string literal\n"

	genericRunBadCode(prog, expected, t)
}

/***** Supporting Methods *****/
func genericRunCode(prog Executable, expected string, t *testing.T) {
	actual := prog.Run()

	assertEquals(expected, actual, t)
}

func genericRunBadCode(prog Executable, expected string, t *testing.T) {
	actual := prog.Run()
	assertEquals(expected, actual, t)
}
