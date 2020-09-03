package executable

import "fmt"

//UnsupportedLanguageError is the error returned by NewExecutable if
//the language provided is not supported.
type UnsupportedLanguageError struct {
	lang string
}

func (ule *UnsupportedLanguageError) Error() string {
	return fmt.Sprintf("%s is not a supported language", ule.lang)
}
