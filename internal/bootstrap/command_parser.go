package bootstrap

import (
	"memo/internal/api/telegram/command"
	"memo/internal/domain"
)

func NewParser(repo domain.UserRepository) command.Parser {
	parser := command.NewParser(repo)
	defaultParser := command.NewDefaultAddParser(parser, repo)

	return defaultParser
}
