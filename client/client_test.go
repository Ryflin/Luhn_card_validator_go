package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

// nothing to test it was just reading input
func Test_read_digit(t *testing.T) {
	tests := []struct {
		name       string
		wantNumber string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumber := read_digit(); gotNumber != tt.wantNumber {
				t.Errorf("read_digit() = %v, want %v", gotNumber, tt.wantNumber)
			}
		})
	}
}

func Test_luhn_digit(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name      string
		args      args
		wantValid bool
	}{
		{"luhn digit wiki example", args{"17893729974"}, true},
		{"luhn digiti wiki example incorrect", args{"17893729984"}, false},
		{"empty string special", args{""}, false},
		{"one digit false", args{"2"}, false},
		{"American Express rand", args{"3782 822463 10005"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValid := luhn_digit(tt.args.digits); gotValid != tt.wantValid {
				t.Errorf("luhn_digit() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}

func Test_validate_with_server(t *testing.T) {
	type args struct {
		username string
		password string
		number   string
	}
	tests := []struct {
		name      string
		args      args
		wantValid bool
		wantLimit uint64
	}{
		{"Test default invalid", args{"bad", "bad", "bad"}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValid, gotLimit := validate_with_server(tt.args.username, tt.args.password, tt.args.number)
			if gotValid != tt.wantValid {
				t.Errorf("validate_with_server() gotValid = %v, want %v", gotValid, tt.wantValid)
			}
			if gotLimit != tt.wantLimit {
				t.Errorf("validate_with_server() gotLimit = %v, want %v", gotLimit, tt.wantLimit)
			}
		})
	}
}

func Test_check_card_Provider(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check_card_Provider()
		})
	}
}
