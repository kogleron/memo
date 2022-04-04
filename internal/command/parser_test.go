package command_test

import (
	"errors"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
			expectedErr: errors.New("empty message"), //nolint: goerr113
		},
		{
			name:    "default command",
			message: " some message  ",
			expectedCmd: &command.Command{
				Name:    "add",
				Payload: "some message",
				Message: &tgbotapi.Message{
					Text: " some message  ",
				},
			},
			expectedErr: nil,
		},
		{
			name:    "parsed command",
			message: "/cmd some message  ",
			expectedCmd: &command.Command{
				Name:    "cmd",
				Payload: "some message",
				Message: &tgbotapi.Message{
					Text: "/cmd some message  ",
				},
			},
			expectedErr: nil,
		},
	}

	for n := range tests {
		tt := &tests[n]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			parser := command.NewParser()
			message := tgbotapi.Message{
				Text: tt.message,
			}

			actualCmd, actualErr := parser.ParseCommand(&message)

			assert.Equal(t, tt.expectedCmd, actualCmd)
			assert.Equal(t, tt.expectedErr, actualErr)
		})
	}
}
