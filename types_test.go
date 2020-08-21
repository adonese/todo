package main

import "testing"

func TestMessage_tokenize(t *testing.T) {
	type fields struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		sep    string
		want   string
	}{
		{"to_test", fields{message: "remindme to drink water at two hours from now"}, "to", "drink water at two hours from now"},
		{"at_test", fields{message: "remindme to drink water at two hours from now"}, "at", "drink water at two hours from now"},
		{"np_test", fields{message: "remindme to drink water at two hours from now"}, "noop", "drink water at two hours from now"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				message: tt.fields.message,
			}
			if got := m.tokenize(" to"); got[0] != tt.want {
				t.Errorf("Message.tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
