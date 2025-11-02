package main

// import (
// 	"testing"
// )

// func init() {
// }

// // TestAdd tests the Add function
// func TestIsEmailValid(t *testing.T) {
// 	testCases := []struct {
// 		name     string
// 		email    string
// 		expected bool
// 	}{
// 		// Valid emails
// 		{"standard email", "example@domain.com", true},
// 		{"with plus addressing", "user+tag@domain.com", true},
// 		{"with dots", "first.last@domain.com", true},
// 		{"with numbers", "user123@domain.com", true},
// 		{"with underscore", "user_name@domain.com", true},
// 		{"with hyphen", "user-name@domain.com", true},
// 		{"subdomain", "user@sub.domain.com", true},
// 		{"country TLD", "user@domain.co.uk", true},
// 		{"long TLD", "user@domain.technology", true},
// 		{"numeric domain", "user@123.com", true},

// 		// Invalid emails
// 		{"missing @ symbol", "exampledomain.com", false},
// 		{"missing domain", "example@", false},
// 		{"missing username", "@domain.com", false},
// 		{"multiple @ symbols", "user@@domain.com", false},
// 		{"space in email", "user name@domain.com", false},
// 		{"leading space", " user@domain.com", false},
// 		{"trailing space", "user@domain.com ", false},
// 		{"missing TLD", "user@domain", false},
// 		{"single character TLD", "user@domain.c", false},
// 		{"special characters", "user!@domain.com", false},
// 		{"consecutive dots", "user..name@domain.com", false},
// 		{"dot at start", ".user@domain.com", false},
// 		{"dot at end", "user.@domain.com", false},
// 		{"domain dot at start", "user@.domain.com", false},
// 		{"domain dot at end", "user@domain.", false},
// 		{"invalid TLD characters", "user@domain.c_m", false},
// 		{"empty string", "", false},
// 		{"just whitespace", "   ", false},
// 		{"null byte", "user@domain.com\x00", false},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			result := validator.IsEmailValid(tc.email)
// 			if result != tc.expected {
// 				t.Errorf("IsEmailValid(%q) = %v; expected %v", tc.email, result, tc.expected)
// 			}
// 		})
// 	}
// }

// func TestIsPassValid(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		pass     string
// 		expected bool
// 	}{
// 		{"basic", "12345678", true},
// 		{"short pass", "1234", false},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			result := validator.IsPasswordLengthValid(tc.pass)
// 			if result != tc.expected {
// 				t.Errorf("IsPassValid(%q) = %v; expected %v ", tc.pass, result, tc.expected)
// 			}
// 		})
// 	}
// }

// func TestLoginApi(t *testing.T) {
// 	testCase := []struct {
// 		name     string
// 		arg1     string
// 		arg2     string
// 		expected model.UserModel
// 	}{
// 		{
// 			name: "user with empty name",
// 			arg1: "jane@example.com",
// 			arg2: "",
// 			expected: model.UserModel{
// 				Id:       1001,
// 				Name:     "Mahadi Hassan",
// 				Password: "1234",
// 				Email:    "mithun.2121@yahoo.com",
// 				Phone:    "1751001003",
// 			},
// 		},
// 	}

// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// result, _ := repository.FindByEmail(tc.arg1)

// 			// if tc.expected.Id != result.Id {
// 			// 	t.Errorf("ERERERERRRRRRR")
// 			// }

// 		})
// 	}
// }
