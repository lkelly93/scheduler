// +build !longTests

package executable

import (
	"strings"
	"testing"
)

func TestInfiniteRecursion(t *testing.T) {
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
	prog, err := NewExecutable("java", code.String(), nil)
	if err != nil {
		t.Error(err)
	}

	expected := "Exception in thread \"main\" java.lang.StackOverflowError\n"
	_, actual := prog.Run()

	errorMessage := actual.Error()
	newLineIndex := strings.Index(errorMessage, "\n") + 1
	errorMessage = errorMessage[:newLineIndex]

	assertRuntimeError(actual, t)
	assertEquals(expected, errorMessage, t)
}

func TestInfiniteLoop(t *testing.T) {
	var code strings.Builder
	code.WriteString("x = 5\n")
	code.WriteString("while(True):\n")
	code.WriteString("\tx+=1\n")
	prog, _ := NewExecutable("python", code.String(), nil)

	_, actual := prog.Run()
	expected := "Time Limit Exceeded 15s"

	assertTimeLimitError(actual, t)
	assertEquals(expected, actual.Error(), t)
}
