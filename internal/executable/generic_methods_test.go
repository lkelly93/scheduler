package executable

import "testing"

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

func assertRuntimeError(err error, t *testing.T) {
	if _, ok := err.(*RuntimeError); !ok {
		t.Errorf("Expected RuntimeError but got %T", err)
	}
}

func assertCompilationError(err error, t *testing.T) {
	if _, ok := err.(*CompilationError); !ok {
		t.Errorf("Expected CompilationError but got %T", err)
	}
}

func assertTimeLimitError(err error, t *testing.T) {
	if _, ok := err.(*TimeLimitExceededError); !ok {
		t.Errorf("Expected TimeLimitExceededError but got %T", err)
	}
}

func assertUnsupportedLanguageError(err error, t *testing.T) {
	if _, ok := err.(*UnsupportedLanguageError); !ok {
		t.Errorf("Expected UnsupportedLanguageError but got %T", err)
	}
}
