package executable

//Executable represents program that is ready to execute
type Executable interface {
	Run() string
}

//FileSettings holds the settings for a runner file
type FileSettings struct {
	Imports      string
	ClassName    string
	TrailingCode string
}

//executableState holds the state of an Executable
type executableState struct {
	code       string
	settings   *FileSettings
	createFile fileCreationFunction
}

type fileCreationFunction func(string, *FileSettings) (string, string)
