package command_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"memo/internal/command"
)

func Test_Parser_ParseCommand(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		message     string
		expectedCmd *command.Command
		expectedErr error
	}{
		{
			name:        "empty message",
			message:     "   ",
			expectedCmd: nil,
			expectedErr: command.ErrEmptyMessage,
		},
		{
			name:    "trim message",
			message: " some message  ",
			expectedCmd: &command.Command{
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
			parser := command.NewParser()

			actualCmd, actualErr := parser.ParseCommand(tt.message)

			assert.Equal(t, tt.expectedCmd, actualCmd)
			assert.Equal(t, tt.expectedErr, actualErr)
		})
	}
}
