// +build !longTests

package executable

import (
	"strings"
	"testing"
)

func TestInfiniteRecursion(t *testing.T) {
	t.Parallel()
	var code strings.Builder
	code.WriteString("public static void main(String[] args) {\n")
	code.WriteString("System.out.println(oops(5));\n")
	code.WriteString("}\n")
	code.WriteString("public static int oops(int x){\n")
	code.WriteString("if(x > 0 ){\n")
	code.WriteString("return oops(++x);\n")
	code.WriteString("}else{")
	code.WriteString("return x;\n")
	code.WriteString("}}")
	exec, err := NewExecutable("java", code.String(), &FileSettings{
		FileNamePrefix: "TestInfiniteRecursion",
	})
	if err != nil {
		t.Error(err)
	}

	expected := "Exception in thread \"main\" java.lang.StackOverflowError\n"
	_, actual := exec.Run()

	if actual == nil {
		t.Fatal("TestInfiniteRecursion's Run did not produce an error.")
	}

	errorMessage := actual.Error()
	newLineIndex := strings.Index(errorMessage, "\n") + 1
	errorMessage = errorMessage[:newLineIndex]

	if !isRuntimeError(actual) {
		t.Errorf("Expected *executable.RuntimeError but got %T", err)
	}
	assertEquals(expected, errorMessage, t)
}

func TestInfiniteLoop(t *testing.T) {
	t.Parallel()
	var code strings.Builder
	code.WriteString("x = 5\n")
	code.WriteString("while(True):\n")
	code.WriteString("\tx+=1\n")
	exec, err := NewExecutable("python", code.String(), &FileSettings{
		FileNamePrefix: "TestInfiniteLoop",
	})
	if err != nil {
		t.Error(err)
	}

	_, actual := exec.Run()
	expected := "Time Limit Exceeded 15s"

	if !isTimeLimitError(actual) {
		t.Errorf("Expected *executable.TimeLimitError but got %T", err)
	}
	assertEquals(expected, actual.Error(), t)
}
