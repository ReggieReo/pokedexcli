package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		expect []string
	}{
		{
			input:  "hello world",
			expect: []string{"hello", "world"},
		},
		{
			input:  " hello world  ",
			expect: []string{"hello", "world"},
		},
		{
			input:  "hEllo World",
			expect: []string{"hello", "world"},
		},
		{
			input:  " hELlo woRld ",
			expect: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expect) {
			t.Errorf("Len of output(%v) is not equal to expted len %v", len(actual), len(c.expect))
			continue
		}
		for i := range actual {
			word := actual[i]
			if word != c.expect[i] {
				t.Errorf("Get %v expected %v", word, c.expect[i])
			}
			continue
		}
	}
}
