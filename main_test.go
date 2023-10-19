package main

import "testing"

func TestMaskLinks(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "Here's my spammy page: http://hehefouls.netHAHAHA see you.",
			expected: "Here's my spammy page: http://******************* see you.",
		},
		{
			input:    "Here's my spammy page: hTTp://youth-elixir.com",
			expected: "Here's my spammy page: hTTp://youth-elixir.com",
		},
		{
			input:    "No links here",
			expected: "No links here",
		},
		{
			input:    "http://www.example.com is a link",
			expected: "http://*************** is a link",
		},
		{
			input:    "http://https://http://",
			expected: "http://********http://",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := MaskLinks(test.input)
			if result != test.expected {
				t.Errorf("Expected: %s, Got: %s", test.expected, result)
			}
		})
	}
}
