package executable

import "testing"

func assertEquals(expected string, actual string, t *testing.T) {
	t.Helper()
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

func isRuntimeError(err error) bool {
	_, ok := err.(*RuntimeError)
	return ok
}

func isCompilationError(err error) bool {
	_, ok := err.(*CompilationError)
	return ok
}

func isTimeLimitError(err error) bool {
	_, ok := err.(*TimeLimitExceededError)
	return ok
}

func isUnsupportedLanguageError(err error) bool {
	_, ok := err.(*UnsupportedLanguageError)
	return ok
}
