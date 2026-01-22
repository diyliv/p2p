package command

import (
	"strings"

	"github.com/diyliv/p2p/internal/models"
)

func ParseCommand(line string) *models.Command {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return &models.Command{}
	}
	return &models.Command{
		Cmd:  parts[0],
		Args: parts[1:],
	}
}
