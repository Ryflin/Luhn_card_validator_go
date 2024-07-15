package main

import (
	"reflect"
	"testing"
)

func Test_readJson(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name        string
		args        args
		wantContent map[string]string
		wantErr     bool
	}{
		{
			name: "Test json1",
			args: args{filename: "testing/1.json"},
			wantContent: map[string]string{
				"when":     "did",
				"I":        "become",
				"so":       "numb",
				"when did": "I",
				"become":   "a",
				"shamed":   "lately",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotContent, err := readJson(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("readJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotContent, tt.wantContent) {
				t.Errorf("readJson() = %v, want %v", gotContent, tt.wantContent)
			}
		})
	}
}

func Test_writeJson(t *testing.T) {
	type args struct {
		filename string
		content  map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeJson(tt.args.filename, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("writeJson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_appendJson(t *testing.T) {
	type args struct {
		content    map[string]string
		newContent map[string]string
	}
	tests := []struct {
		name             string
		args             args
		wantFinalContent map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFinalContent := appendJson(tt.args.content, tt.args.newContent); !reflect.DeepEqual(gotFinalContent, tt.wantFinalContent) {
				t.Errorf("appendJson() = %v, want %v", gotFinalContent, tt.wantFinalContent)
			}
		})
	}
}
