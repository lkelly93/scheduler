package executable

//Executable represents program that is ready to execute
type Executable interface {
	Run() (string, error)
}

//FileSettings holds the settings for a runner file
type FileSettings struct {
	Imports      string
	ClassName    string
	TrailingCode string
	//FileNamePrefix is the will be put in front of the class name. Only needed if
	//you are planning on running multiple executable in parallel and want to
	//prevent against data races. Defaults to empty string
	FileNamePrefix string
}

//executableState holds the state of an Executable
type executableState struct {
	code       string
	lang       string
	settings   *FileSettings
	createFile fileCreationFunction
}

type fileCreationFunction func(string, *FileSettings) (string, string, error)

var supportedLanguages = map[string]fileCreationFunction{
	"python": createRunnerFilePython,
	"java":   createRunnerFileJava,
}

var fileSettingsDefaults = map[string]FileSettings{
	"python": FileSettings{
		Imports:        "import numpy as np",
		ClassName:      "PythonRunner",
		TrailingCode:   "",
		FileNamePrefix: "",
	},
	"java": FileSettings{
		Imports:        "import java.util.*;",
		ClassName:      "JavaRunner",
		TrailingCode:   "",
		FileNamePrefix: "",
	},
}
