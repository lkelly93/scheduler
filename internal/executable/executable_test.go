//Package executable represents a program written in a generic language.
//This package can run the given program and return the result
package executable

import "testing"

func TestDebug(t *testing.T) {
	lang := "python"
	code := "os.system(\"hostname\")"
	settings := FileSettings{
		Imports: "import os",
	}
	exe, _ := NewExecutable(lang, code, &settings)

	out, err := exe.Run()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Errorf("Output: %s", out)
}

// import (
// 	"io/ioutil"
// 	"os"
// 	"reflect"
// 	"strings"
// 	"testing"
// )

// func TestNewExecutable(t *testing.T) {
// 	type args struct {
// 		lang     string
// 		code     string
// 		settings *FileSettings
// 	}
// 	allTests := []struct {
// 		name          string
// 		args          args
// 		expected      Executable
// 		expectedError error
// 	}{
// 		{
// 			name: "NewExecutable",
// 			args: args{
// 				lang: "python",
// 				code: "print('Hello World')",
// 			},
// 			expected: &executableState{
// 				code:     "print('Hello World')",
// 				lang:     "python",
// 				settings: nil,
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name: "NewExecutableFail",
// 			args: args{
// 				lang: "Not a Language",
// 				code: "Not Code",
// 			},
// 			expected: nil,
// 			expectedError: &UnsupportedLanguageError{
// 				lang: "Not a Language",
// 			},
// 		},
// 	}
// 	for _, test := range allTests {
// 		//The below line of code is needed because of a Go gotcha hidden inside
// 		//the go testing framework.
// 		//Read https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
// 		//For more information.
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			got, err := NewExecutable(test.args.lang, test.args.code, test.args.settings)
// 			if (err != nil) && (test.expectedError != nil) {
// 				if isUnsupportedLanguageError(test.expectedError) {
// 					if !isUnsupportedLanguageError(err) {
// 						t.Errorf("Expected UnsupportedLanguageError but got %T", err)
// 						return
// 					}
// 				} else {
// 					t.Errorf("%s returned an error that was not expected, error was %T", test.name, err)
// 					return
// 				}
// 				assertEquals(test.expectedError.Error(), err.Error(), t)
// 				return
// 				//Check if error happened and it should not have.
// 			} else if (err != nil) && (test.expectedError == nil) {
// 				t.Error(err)
// 				return
// 			}
// 			//Have to set got's function field to nil because DeeplyEquals fails if
// 			//function values are not equal to nil. I will test this manually in
// 			//another method.
// 			gotState := got.(*executableState)
// 			gotState.createFile = nil

// 			if !reflect.DeepEqual(got, test.expected) {
// 				t.Errorf("NewExecutable() was %v, want %v", got, test.expected)
// 			}
// 		})
// 	}
// }

// func TestRuns(t *testing.T) {
// 	var JavaCustomSettingsCode strings.Builder
// 	JavaCustomSettingsCode.WriteString("public static void main(String[] args){\n")
// 	JavaCustomSettingsCode.WriteString("HashMap<Integer,Integer> x = new HashMap<>();\n")
// 	JavaCustomSettingsCode.WriteString("x.put(5,4);\n")
// 	JavaCustomSettingsCode.WriteString("System.out.println(x.remove(5));\n")
// 	JavaCustomSettingsCode.WriteString("pi();\n")
// 	JavaCustomSettingsCode.WriteString("}\n")

// 	var BadJavaCodeExpectedMessage strings.Builder
// 	BadJavaCodeExpectedMessage.WriteString("JavaRunner.java:3: error: ';' expected\n")
// 	BadJavaCodeExpectedMessage.WriteString("public static void main(String[] args){System.out.println(\"Hello World\")\n")
// 	BadJavaCodeExpectedMessage.WriteString("                                                                        ^\n")
// 	BadJavaCodeExpectedMessage.WriteString("JavaRunner.java:5: error: reached end of file while parsing\n")
// 	BadJavaCodeExpectedMessage.WriteString("}\n")
// 	BadJavaCodeExpectedMessage.WriteString(" ^\n")
// 	BadJavaCodeExpectedMessage.WriteString("2 errors\n")
// 	BadJavaCodeExpectedMessage.WriteString("error: compilation failed\n")

// 	var BadPythonCodeExpectedMessage strings.Builder
// 	BadPythonCodeExpectedMessage.WriteString("  File \"PythonRunner.py\", line 2\n")
// 	BadPythonCodeExpectedMessage.WriteString("    print('Hi\n")
// 	BadPythonCodeExpectedMessage.WriteString("            ^\n")
// 	BadPythonCodeExpectedMessage.WriteString("SyntaxError: EOL while scanning string literal\n")

// 	pythonLongCodeLocation := "test_data/longPythonCode.py"
// 	pythonLongCodeData, err := ioutil.ReadFile(pythonLongCodeLocation)
// 	if err != nil {
// 		t.Errorf("Could not read in %s", pythonLongCodeLocation)
// 	}
// 	javaLongCodeLocation := "test_data/longJavaCode.java"
// 	javaLongCodeData, err := ioutil.ReadFile(javaLongCodeLocation)
// 	if err != nil {
// 		t.Errorf("Could not read in %s", javaLongCodeLocation)
// 	}

// 	recursiveFileLocation := "test_data/recursiveCode.java"
// 	recursiveFileData, err := ioutil.ReadFile(recursiveFileLocation)
// 	if err != nil {
// 		t.Errorf("Could not read in %s", recursiveFileLocation)
// 	}

// 	type args struct {
// 		lang     string
// 		code     string
// 		settings *FileSettings
// 	}
// 	allTests := []struct {
// 		name          string
// 		args          args
// 		expected      string
// 		expectedError error
// 	}{
// 		{
// 			name: "TestRunPythonCode",
// 			args: args{
// 				lang:     "python",
// 				code:     "print('Hello World')",
// 				settings: nil,
// 			},
// 			expected:      "Hello World\n",
// 			expectedError: nil,
// 		},
// 		{
// 			name: "TestRunPythonCodeCustomFileSettings",
// 			args: args{
// 				lang: "python",
// 				code: "print(np.e)",
// 				settings: &FileSettings{
// 					Imports:        "import math\nimport numpy as np",
// 					ClassName:      "",
// 					TrailingCode:   "print(math.tau)",
// 					FileNamePrefix: "",
// 				},
// 			},
// 			expected:      "2.718281828459045\n6.283185307179586\n",
// 			expectedError: nil,
// 		},
// 		{
// 			name: "TestRunJavaCode",
// 			args: args{
// 				lang:     "java",
// 				code:     "public static void main(String[] args){System.out.println(\"Hello World\");}",
// 				settings: nil,
// 			},
// 			expected:      "Hello World\n",
// 			expectedError: nil,
// 		},
// 		{
// 			name: "TestRunJavaCodeCustomFileSettings",
// 			args: args{
// 				lang: "java",
// 				code: JavaCustomSettingsCode.String(),
// 				settings: &FileSettings{
// 					Imports:        "import java.lang.*;\n import java.util.*;",
// 					ClassName:      "",
// 					TrailingCode:   "public static void pi(){System.out.println(Math.PI);}",
// 					FileNamePrefix: "",
// 				},
// 			},
// 			expected:      "4\n3.141592653589793\n",
// 			expectedError: nil,
// 		},
// 		{
// 			name: "TestRunPythonCodeLonger",
// 			args: args{
// 				lang:     "python",
// 				code:     string(pythonLongCodeData),
// 				settings: nil,
// 			},
// 			expected:      "Male\n",
// 			expectedError: nil,
// 		},
// 		{
// 			name: "TestRunJavaCodeLonger",
// 			args: args{
// 				lang:     "java",
// 				code:     string(javaLongCodeData),
// 				settings: nil,
// 			},
// 			expected:      "NonRecursive\n[0, 1, 0, 0, 1, 0, 1, 0]\n[0, 0, 0, 0, 0, 1, 1, 0]\n",
// 			expectedError: nil,
// 		},
// 		{
// 			name: "TestRecursion",
// 			args: args{
// 				lang:     "java",
// 				code:     string(recursiveFileData),
// 				settings: nil,
// 			},
// 			expected:      "Recursive\n[0, 1, 0, 0, 1, 0, 1, 0]\n[0, 0, 0, 0, 0, 1, 1, 0]\n",
// 			expectedError: nil,
// 		},
// 		{
// 			name: "TestBadJavaCode",
// 			args: args{
// 				lang:     "java",
// 				code:     "public static void main(String[] args){System.out.println(\"Hello World\")",
// 				settings: nil,
// 			},
// 			expected: "",
// 			expectedError: &RuntimeError{
// 				errMessage: BadJavaCodeExpectedMessage.String(),
// 			},
// 		},
// 		{
// 			name: "TestBadPythonCode",
// 			args: args{
// 				lang:     "python",
// 				code:     "print('Hi",
// 				settings: nil,
// 			},
// 			expected: "",
// 			expectedError: &RuntimeError{
// 				errMessage: BadPythonCodeExpectedMessage.String(),
// 			},
// 		},
// 	}
// 	for _, test := range allTests {
// 		//The below line of code is needed because of a Go gotcha hidden inside
// 		//the go testing framework.
// 		//Read https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
// 		//For more information.
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			if test.args.settings == nil {
// 				test.args.settings = &FileSettings{}
// 			}
// 			test.args.settings.FileNamePrefix = test.name
// 			exec, err := NewExecutable(test.args.lang,
// 				test.args.code,
// 				test.args.settings)
// 			if err != nil {
// 				t.Error(err)
// 				return
// 			}
// 			actual, err := exec.Run()
// 			if (err != nil) && test.expectedError != nil {
// 				//Check error types
// 				if isRuntimeError(test.expectedError) {
// 					if !isRuntimeError(err) {
// 						t.Errorf("%s returned an unexpected error type, err was %T", test.name, err)
// 						return
// 					}
// 				} else if isCompilationError(test.expectedError) {
// 					if !isCompilationError(err) {
// 						t.Errorf("%s returned an unexpected error type, err was %T", test.name, err)
// 						return
// 					}
// 				} else {
// 					t.Errorf("%s returned an error that was not expected, error was %T", test.name, err)
// 					return
// 				}

// 				assertEquals(test.expectedError.Error(), err.Error(), t)
// 				return
// 			} else if (err != nil) && (test.expectedError == nil) {
// 				t.Error(err)
// 				return
// 			}
// 			assertEquals(test.expected, actual, t)
// 		})
// 	}
// }

// func TestFileIsDeletedAfterRun(t *testing.T) {
// 	t.Parallel()
// 	exec, err := NewExecutable("python", "print('Hello World')", &FileSettings{
// 		FileNamePrefix: "TestFileIsDeletedAfterRun",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fileLocation := "../runner_files/TestFileIsDeletedAfterRunPythonRunner.py"
// 	_, err = os.Stat(fileLocation)
// 	if err == nil {
// 		t.Fatalf("%s existed before Run() was called", fileLocation)
// 	}

// 	exec.Run()

// 	_, err = os.Stat(fileLocation)
// 	if err == nil {
// 		t.Errorf("%s still exist after Run() was called. It should of been deleted", fileLocation)
// 	}
// }
