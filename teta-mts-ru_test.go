package main

import (
	"testing"
)

func TestCounter(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"Test 1", "dcaaabbbccccc", "a3b3c6d1"},
		{"Test 2", "aaabbbcccccaaaaa", "a8b3c5"},
		{"Test 3", "zzzzcccUUUuu", "U3c3u2z4"},
		{"Test 4", "ЯЯЯБББддд", "Б3Я3д3"},
		{"Test 5", "Я  + Я + ЯБ ББдд  ! @@ ##### д ", " 12!1#5+2@2Б3Я3д3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Counter(tt.arg); got != tt.want {
				t.Errorf("Counter() = %v, want %v", got, tt.want)
			}
		})
	}
}
