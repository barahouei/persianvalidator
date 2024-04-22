package persianvalidator

import (
	"testing"
)

var (
	chractersSlice = []rune{
		'ء',
		'آ',
		'أ',
		'ؤ',
		'إ',
		'ئ',
		'ا',
		'ب',
		'ت',
		'ث',
		'ج',
		'ح',
		'خ',
		'د',
		'ذ',
		'ر',
		'ز',
		'س',
		'ش',
		'ص',
		'ض',
		'ط',
		'ظ',
		'ع',
		'غ',
		'ف',
		'ق',
		'ل',
		'م',
		'ن',
		'ه',
		'و',
		'َ',
		'ُ',
		'ِ',
		'ّ',
		'ٕ',
		'پ',
		'چ',
		'ژ',
		'ک',
		'گ',
		'ھ',
		'ی',
	}

	numbersSlice = []rune{
		'۰',
		'۱',
		'۲',
		'۳',
		'۴',
		'۵',
		'۶',
		'۷',
		'۸',
		'۹',
	}

	spacesSlice = []rune{
		' ',
		' ',
		' ',
		' ',
		' ',
		' ',
		' ',
		' ',
		' ',
		' ',
		' ',
		' ',
		'​',
		'‌',
		'‍',
		'‎',
		'‏',
		' ',
		' ',
		'‪',
		'‫',
		'‬',
		'‭',
		'‮',
		' ',
	}

	testCases = []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "persian text 1",
			input: "امضاء",
			want:  true,
		},
		{
			name:  "persian text 2",
			input: "ژاله",
			want:  true,
		},
		{
			name:  "persian text 3",
			input: "مؤمنی",
			want:  true,
		},
		{
			name:  "persian text 4",
			input: "گلمکان",
			want:  true,
		},
		{
			name:  "persian text with space",
			input: "ادب آداب دارد",
			want:  true,
		},
		{
			name:  "persian text with space and period",
			input: "انسان همواره در جستجوی معنا است.",
			want:  false,
		},
		{
			name:  "persian text with non persian text",
			input: "hello دنیا",
			want:  false,
		},
		{
			name:  "non persian text",
			input: "Cogito ergo sum",
			want:  false,
		},
		{
			name:  "persian number",
			input: "۱",
			want:  true,
		},
		{
			name:  "non persian number",
			input: "1",
			want:  false,
		},
		{
			name:  "persian text with persian number",
			input: "یک برابر است با ۱",
			want:  true,
		},
		{
			name:  "persian text with non persian number",
			input: "یک برابر است با 1",
			want:  false,
		},
	}
)

func TestCheckCharacter(t *testing.T) {
	for _, ch := range chractersSlice {
		if value, ok := Characters[ch]; ok {
			if value != ch {
				t.Errorf("value of character %c with unicode %U and rune %d is wrong.", ch, ch, ch)
			}
		}
	}
}

func TestCheckNumbers(t *testing.T) {
	for _, ch := range numbersSlice {
		if value, ok := Numbers[ch]; ok {
			if value != ch {
				t.Errorf("value of character %c with unicode %U and rune %d is wrong.", ch, ch, ch)
			}
		}
	}
}

func TestCheckSpaces(t *testing.T) {
	for _, ch := range spacesSlice {
		if value, ok := Spaces[ch]; ok {
			if value != ch {
				t.Errorf("value of character %c with unicode %U and rune %d is wrong.", ch, ch, ch)
			}
		}
	}
}

func TestIsValid(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValid(tc.input)

			if tc.want != got {
				t.Errorf("want: %t, but got: %t.", tc.want, got)
			}
		})
	}
}
