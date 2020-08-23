package program_test

import (
	"github.com/lkelly93/code-runner/program"
	"testing"
	"os"
)

func TestPythonCreateFile(t *testing.T){
	prog := program.NewProgram("python", "print('Hello World')")
	runnerFileFunctor := program.GetFunctor(prog.Lang)

	actual := runnerFileFunctor(&prog)

	expected := "python3 ../runnerFiles/PythonRunner.py"
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected , actual)
	}

	os.Remove("../runnerFiles/PythonRunner.py")
}