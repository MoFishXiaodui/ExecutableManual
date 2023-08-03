package __1_6unitTest

import (
	"testing"
)

func TestHelloTom(t *testing.T) {
	// T is a type passed to Test functions to manage test state and support formatted test logs.
	output := HelloTom()
	expectOutput := "Tom"

	if output != expectOutput {
		t.Errorf("Expected %s do not match actual %s", expectOutput, output)
	}
}

//func TestMain(m *testing.M) {
// M is a type passed to a TestMain function to run the actual tests.
//	code := m.Run()
//	os.Exit(code)
//}
