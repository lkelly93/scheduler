package executable

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	type args struct {
		lang     string
		code     string
		settings *FileSettings
	}
	allTests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "TestPythonCreateFile",
			args: args{
				lang:     "python",
				code:     "print('Hello World')",
				settings: nil,
			},
			expected: "python3 " + getRunnerFileLocation("TestPythonCreateFilePythonRunner.py"),
		},
		{
			name: "TestPythonCreateFileCustomFileSetttings",
			args: args{
				lang: "python",
				code: "print('Hello World')",
				settings: &FileSettings{
					Imports:        "",
					ClassName:      "SillyPythonName",
					TrailingCode:   "",
					FileNamePrefix: "TestPythonCreateFileCustomFileSetttings",
				},
			},
			expected: "python3 " + getRunnerFileLocation("TestPythonCreateFileCustomFileSetttingsSillyPythonName.py"),
		},
		{
			name: "TestJavaCreateFile",
			args: args{
				lang:     "java",
				code:     "public static void main(String[] args){System.out.println(\"Hello World\");}",
				settings: nil,
			},
			expected: "java " + getRunnerFileLocation("TestJavaCreateFileJavaRunner.java"),
		},
		{
			name: "TestJavaCreateFileCustomFileSettings",
			args: args{
				lang: "java",
				code: "public static void main(String[] args){System.out.println(\"Hello World\");}",
				settings: &FileSettings{
					Imports:        "",
					ClassName:      "SillyJavaName",
					TrailingCode:   "",
					FileNamePrefix: "TestJavaCreateFileCustomFileSettings",
				},
			},
			expected: "java " + getRunnerFileLocation("TestJavaCreateFileCustomFileSettingsSillyJavaName.java"),
		},
	}

	for _, test := range allTests {
		//The below line of code is needed because of a Go gotcha hidden inside
		//the go testing framework.
		//Read https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
		//For more information.
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.args.settings = fillRestOfFileSettings(test.args.lang, test.args.settings)
			test.args.settings.FileNamePrefix = test.name
			// if test.args.settings == nil {
			// 	test.args.settings = &FileSettings{}
			// 	test.args.settings.FileNamePrefix = test.name
			// }

			createFileFunction := getFileCreationFunction(test.args.lang)
			sysCommand, fileLocation, _ := createFileFunction(test.args.code, test.args.settings)
			defer os.Remove(fileLocation)
			actual := sysCommand + " " + fileLocation

			assertEquals(test.expected, actual, t)
		})
	}
}

func TestCreateRunnerFile(t *testing.T) {
	t.Parallel()
	lang := "python"
	code := "print('Hello World')"

	createFileFunction := getFileCreationFunction(lang)
	_, fileLocation, _ := createFileFunction(code, &FileSettings{
		FileNamePrefix: "TestCreateRunnerFile",
	})
	defer os.Remove(fileLocation)

	_, err := os.Stat(fileLocation)
	if err != nil {
		t.Errorf("%s file was not created", fileLocation)
	}
}

func TestRemoveFilePath(t *testing.T) {
	t.Parallel()
	message := "/path/to/runner/file/PythonRunner.py had an error Python();<_aRunner.py"
	expected := "PythonRunner.py had an error Python();<_aRunner.py"
	mockFilePath := "/path/to/runner/file/PythonRunner.py"
	actual := removeFilePath(message, mockFilePath)

	assertEquals(expected, actual, t)
}

func TestAllSupportedDefaultSettingsJava(t *testing.T) {
	t.Parallel()
	actual := fileSettingsDefaults["java"]
	expected := FileSettings{
		Imports:        "import java.util.*;",
		ClassName:      "JavaRunner",
		TrailingCode:   "",
		FileNamePrefix: "",
	}

	assertEquals(expected.Imports, actual.Imports, t)
	assertEquals(expected.ClassName, actual.ClassName, t)
	assertEquals(expected.TrailingCode, actual.TrailingCode, t)
	assertEquals(expected.FileNamePrefix, actual.FileNamePrefix, t)
}

func TestAllSupportedDefaultSettingsPython(t *testing.T) {
	t.Parallel()
	actual := fileSettingsDefaults["python"]
	expected := FileSettings{
		Imports:        "import numpy as np",
		ClassName:      "PythonRunner",
		TrailingCode:   "",
		FileNamePrefix: "",
	}

	assertEquals(expected.Imports, actual.Imports, t)
	assertEquals(expected.ClassName, actual.ClassName, t)
	assertEquals(expected.TrailingCode, actual.TrailingCode, t)
	assertEquals(expected.FileNamePrefix, actual.FileNamePrefix, t)
}

func TestCreateFileAndAddCode(t *testing.T) {
	t.Parallel()
	outFile := "TestCreateFileAndAddCode.txt"
	code := "This is a test\n"

	err := createFileAndAddCode(outFile, code)
	defer os.Remove(outFile)

	if err != nil {
		t.Errorf("Could not create a runner file")
	}

	file, err := ioutil.ReadFile(outFile)
	if err != nil {
		t.Errorf("Couldn't open runner file after it was created.")
	}

	fileText := string(file)
	assertEquals(code, fileText, t)
}

func TestGetRunnerFileLocation(t *testing.T) {
	t.Parallel()
	suffix := "Test.txt"
	expected := "/securefs/runner_files/Test.txt"

	actual := getRunnerFileLocation(suffix)

	assertEquals(expected, actual, t)
}
