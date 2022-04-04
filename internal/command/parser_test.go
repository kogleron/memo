package command

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parser_ParseCommand(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		message     string
		expectedCmd *Command
		expectedErr error
	}{
		{
			name:        "empty message",
			message:     "   ",
			expectedCmd: nil,
			expectedErr: errors.New("empty message"),
		},
		{
			name:    "trim message",
			message: " some message  ",
			expectedCmd: &Command{
				Name:    "add",
				Payload: "some message",
			},
			expectedErr: nil,
		},
	}

	for n := range tests {
		tt := &tests[n]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			parser := NewParser()

			actualCmd, actualErr := parser.ParseCommand(tt.message)

			assert.Equal(t, tt.expectedCmd, actualCmd)
			assert.Equal(t, tt.expectedErr, actualErr)
		})
	}
}
