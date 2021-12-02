package cmd

import (
	"reflect"
	"testing"

	"github.com/slack-go/slack"
)

func Test_getClient(t *testing.T) {
	tests := []struct {
		name string
		want *slack.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
