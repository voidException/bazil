package get

import (
	"os"

	"bazil.org/bazil/cas/flagx"
	clicas "bazil.org/bazil/cli/debug/cas"
	"bazil.org/bazil/cliutil/subcommands"
)

type getCommand struct {
	subcommands.Description
	Arguments struct {
		Type  string
		Level uint8
		Key   flagx.KeyParam
	}
}

func (c *getCommand) Run() error {
	chunk, err := clicas.CAS.State.Store.Get(
		c.Arguments.Key.Key(),
		c.Arguments.Type,
		c.Arguments.Level,
	)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(chunk.Buf)
	if err != nil {
		return err
	}
	return nil
}

var get = getCommand{
	Description: "get a chunk from CAS",
}

func init() {
	subcommands.Register(&get)
}
