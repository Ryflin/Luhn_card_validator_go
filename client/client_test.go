package main

import "testing"

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
		{"American Express rand", args{"3782 822463 10005"}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValid := luhn_digit(tt.args.digits); gotValid != tt.wantValid {
				t.Errorf("luhn_digit() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}
