package main

import "testing"

// https://leetcode.com/problems/valid-palindrome/
func Test_checkPalindrome(t *testing.T) {
	type args struct {
		str *string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Positive", args: args{str: stringPtr("A man, a plan, a canal: Panama")}, want: true},
		{name: "Negative", args: args{str: stringPtr("race a car")}, want: false},
		{name: "Positive", args: args{str: stringPtr(" ")}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPalindrome(tt.args.str); got != tt.want {
				t.Errorf("checkPalindrome() = %v, want %v on args: %v", got, tt.want, *tt.args.str)
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}
