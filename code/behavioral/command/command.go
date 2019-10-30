package command

// Command : 基本的には1つの命令を表すのでExecute(実行)だけを持つ
type Command interface {
	Execute() bool
}

type command struct {
}

func NewCommand() *command {
	return &command{}
}

func (c *command) Execute() bool {
	return true
}
