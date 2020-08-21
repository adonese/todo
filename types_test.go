package main

import (
	"reflect"
	"testing"
)

func Test_encodeDecode(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"testing", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := encodeDecode(); (err != nil) != tt.wantErr {
				t.Errorf("encodeDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_readFile(t *testing.T) {
	tests := []struct {
		name    string
		want    Storage
		wantErr bool
	}{
		{"successful test", Storage{Description: "my work"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
