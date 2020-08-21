package main

import "testing"

func Test_capture(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"successful_finding", args{"to drink water at two hours from now"}, "drink water at two hours from now"},
		{"successful_finding", args{"tod drink water at two hours from now"}, "drink water at two hours from now"},
		{"successful_finding", args{"tonn drink water at two hours from now"}, "drink water at two hours from now"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := capture(tt.args.s); got != tt.want {
				t.Errorf("capture() = %v, want %v", got, tt.want)
			}
		})
	}
}
