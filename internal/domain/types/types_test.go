package types

import "testing"

type PhoneStructTest struct {
	name     string
	phone    PhoneNumber
	expected bool
}

func TestPhoneValidate(t *testing.T) {
	tests := []PhoneStructTest{
		{
			name:     "valid +7",
			phone:    PhoneNumber("+79230322122"),
			expected: true,
		},
		{
			name:     "valid 8",
			phone:    PhoneNumber("89230322122"),
			expected: true,
		},
		{
			name:     "valid no prefix",
			phone:    PhoneNumber("9230322122"),
			expected: true,
		},
		{
			name:     "too short",
			phone:    PhoneNumber("8923"),
			expected: false,
		},
		{
			name:     "too long",
			phone:    PhoneNumber("8923032212223"),
			expected: false,
		},
		{
			name:     "invalid prefix",
			phone:    PhoneNumber("+19230322122"),
			expected: false,
		},
		{
			name:     "invalid middle digit",
			phone:    PhoneNumber("+72230322122"),
			expected: false,
		},
		{
			name:     "contains letters",
			phone:    PhoneNumber("+792s0a2212b"),
			expected: false,
		},
		{
			name:     "only prefix",
			phone:    PhoneNumber("+7"),
			expected: false,
		},
		{
			name:     "empty",
			phone:    PhoneNumber(""),
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.phone.ValidatePhone()
			if result != test.expected {
				t.Errorf("test %s failed: expected %v, got %v",
					test.name, test.expected, result)
			}
		})
	}
}
