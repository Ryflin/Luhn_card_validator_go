package client

import "testing"

func Test_dammEncode(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name      string
		args      args
		wantCheck string
	}{
		{"Wikipedia example", args{"572"}, "5724"},
		{"Stolen example", args{"43881234567"}, "438812345679"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCheck := dammEncode(tt.args.number); gotCheck != tt.wantCheck {
				t.Errorf("dammEncode() = %v, want %v", gotCheck, tt.wantCheck)
			}
		})
	}
}

func Test_dammDecode(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name        string
		args        args
		wantCorrect bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCorrect := dammDecode(tt.args.number); gotCorrect != tt.wantCorrect {
				t.Errorf("dammDecode() = %v, want %v", gotCorrect, tt.wantCorrect)
			}
		})
	}
}
