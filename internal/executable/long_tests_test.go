//+build !longTests

package executable_test

import (
	"strings"
	"testing"

	"github.com/lkelly93/scheduler/internal/executable"
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
	prog, err := executable.NewExecutable("java", code.String())
	if err != nil {
		t.Error(err)
	}

	expected := "Exception in thread \"main\" java.lang.StackOverflowError\n"
	actual := prog.Run()

	newLineIndex := strings.Index(actual, "\n") + 1
	actual = actual[:newLineIndex]

	assertEquals(expected, actual, t)
}

func TestInfiniteLoop(t *testing.T) {
	var code strings.Builder
	code.WriteString("x = 5\n")
	code.WriteString("while(True):\n")
	code.WriteString("\tx+=1\n")
	prog, _ := executable.NewExecutable("python", code.String())

	actual := prog.Run()
	expected := "Time Limit Exceeded 15s"

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
