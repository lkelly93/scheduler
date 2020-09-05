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
	exe := getNewExecutableForTesting(lang, code, t)

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
	assertUnsupportedLanguageError(err, t)
}

// /***** Test Good Runs *****/
func TestRunPythonCode(t *testing.T) {
	lang := "python"
	code := "print('Hello World')"
	prog := getNewExecutableForTesting(lang, code, t)
	expected := "Hello World\n"
	genericRunCode(prog, expected, t)
}

func TestRunPythonCodeCustomFileSettings(t *testing.T) {
	lang := "python"
	code := "print(np.e)"

	settings := FileSettings{
		Imports:        "import math\nimport numpy as np",
		ClassName:      "",
		TrailingCode:   "print(math.tau)",
		FileNamePrefix: "EXECUTABLE_TESTS",
	}
	exec, _ := NewExecutable(lang, code, &settings)

	actual, _ := exec.Run()
	expected := "2.718281828459045\n6.283185307179586\n"
	assertEquals(expected, actual, t)
}
func TestRunJavaCodeCustomFileSettings(t *testing.T) {
	lang := "java"
	var code strings.Builder
	code.WriteString("public static void main(String[] args){\n")
	code.WriteString("HashMap<Integer,Integer> x = new HashMap<>();\n")
	code.WriteString("x.put(5,4);\n")
	code.WriteString("System.out.println(x.remove(5));\n")
	code.WriteString("pi();\n")
	code.WriteString("}\n")

	settings := FileSettings{
		Imports:        "import java.lang.*;\n import java.util.*;",
		ClassName:      "",
		TrailingCode:   "public static void pi(){System.out.println(Math.PI);}",
		FileNamePrefix: "EXECUTABLE_TESTS",
	}
	exec, _ := NewExecutable(lang, code.String(), &settings)

	actual, err := exec.Run()

	if err != nil {
		t.Fatal(err)
	}

	expected := "4\n3.141592653589793\n"
	assertEquals(expected, actual, t)
}

func TestRunPythonCodeLonger(t *testing.T) {
	fileLocation := "test_data/longPythonCode.py"
	longCodeFile, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		t.Errorf("Could not read in %s", fileLocation)
	}
	code := string(longCodeFile)
	exec := getNewExecutableForTesting("python", code, t)
	expected := "Male\n"
	genericRunCode(exec, expected, t)
}

func TestRunJavaCode(t *testing.T) {
	code := "public static void main(String[] args){System.out.println(\"Hello World\");}"
	exec := getNewExecutableForTesting("java", code, t)
	expected := "Hello World\n"
	genericRunCode(exec, expected, t)
}

func TestRunJavaCodeLonger(t *testing.T) {
	fileLocation := "test_data/longJavaCode.java"
	longCode, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		t.Errorf("Could not read in %s", fileLocation)
	}
	code := string(longCode)
	exec := getNewExecutableForTesting("java", code, t)
	var expected strings.Builder
	expected.WriteString("NonRecursive\n")
	expected.WriteString("[0, 1, 0, 0, 1, 0, 1, 0]\n")
	expected.WriteString("[0, 0, 0, 0, 0, 1, 1, 0]\n")
	genericRunCode(exec, expected.String(), t)
}

func TestRecursion(t *testing.T) {
	fileLocation := "test_data/recursiveCode.java"
	longCode, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		t.Errorf("Could not read in %s", fileLocation)
	}
	code := string(longCode)
	exec := getNewExecutableForTesting("java", code, t)
	var expected strings.Builder
	expected.WriteString("Recursive\n")
	expected.WriteString("[0, 1, 0, 0, 1, 0, 1, 0]\n")
	expected.WriteString("[0, 0, 0, 0, 0, 1, 1, 0]\n")
	genericRunCode(exec, expected.String(), t)

}

func TestFileIsDeletedAfter(t *testing.T) {
	exec := getNewExecutableForTesting("python", "print('Hello World')", t)
	fileLocation := "../runner_files/DeletedAfterTestPythonRunner.py"
	_, err := os.Stat(fileLocation)
	if err == nil {
		t.Fatalf("%s existed before Run() was called", fileLocation)
	}

	exec.Run()

	_, err = os.Stat(fileLocation)
	if err == nil {
		t.Errorf("%s still exist after Run() was called. It should of been deleted", fileLocation)
	}
}

/***** Test Bad Runs*****/
func TestRunBadJavaCode(t *testing.T) {
	code := "public static void main(String[] args){System.out.println(\"Hello World\")"
	exec := getNewExecutableForTesting("java", code, t)
	expected := "EXECUTABLE_TESTSJavaRunner.java:3: error: ';' expected\n" +
		"public static void main(String[] args){System.out.println(\"Hello World\")\n" +
		"                                                                        ^\n" +
		"EXECUTABLE_TESTSJavaRunner.java:5: error: reached end of file while parsing\n" +
		"}\n" +
		" ^\n" +
		"2 errors\n" +
		"error: compilation failed\n"

	genericRuntimeErrorTest(exec, expected, t)

}

func TestRunBadPythonCode(t *testing.T) {
	exec := getNewExecutableForTesting("python", "print('Hi", t)
	expected := "  File \"EXECUTABLE_TESTSPythonRunner.py\", line 2\n" +
		"    print('Hi\n" +
		"            ^\n" +
		"SyntaxError: EOL while scanning string literal\n"

	genericRuntimeErrorTest(exec, expected, t)
}

/***** Supporting Methods *****/
func genericRunCode(prog Executable, expected string, t *testing.T) {
	actual, _ := prog.Run()

	assertEquals(expected, actual, t)
}

func genericRuntimeErrorTest(prog Executable, expected string, t *testing.T) {
	_, actual := prog.Run()
	assertRuntimeError(actual, t)
	assertEquals(expected, actual.Error(), t)
}
