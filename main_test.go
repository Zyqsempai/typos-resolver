package typos

import (
	"testing"
)

func TestMakeReverse(t *testing.T) {
	var testCases = []struct {
		toLang string
		input  string
		expect string
	}{
		{"ru", "ghbdtn", "привет"},
		{"en", "руддщ", "hello"},
		{"ru", "Ghbdtn 12", "Привет 12"},
		{"en", "Руддщ 43", "Hello 43"},
		{"ru", ",./[]", "бю/хъ"},
		{"en", "эхъэ", "'[]'"},
	}

	for index, smlt := range testCases {
		got := MakeReverse(smlt.input, smlt.toLang)
		if got != smlt.expect {
			t.Errorf(
				"%d. MakeReverse(%#v, %#v) = %#v; expect %#v",
				index, smlt.input, smlt.toLang, got, smlt.expect)
		}
	}
}

func TestSubstituteTypos(t *testing.T) {
	var testCases = []struct {
		toLang string
		input  string
		expect string
	}{
		{"en", "Неllo", "hello"},
		{"ru", "cok", "сок"},
		{"en", "Неllo 12", "hello 12"},
		{"ru", "cok 43", "сок 43"},
	}

	for index, smlt := range testCases {
		got := SubstituteTypos(smlt.input, smlt.toLang)
		if got != smlt.expect {
			t.Errorf(
				"%d. SubstituteTypos(%#v, %#v) = %#v; expect %#v",
				index, smlt.input, smlt.toLang, got, smlt.expect)
		}
	}
}
