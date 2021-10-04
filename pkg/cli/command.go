package cli

//Command interface that cli commands should conform to
type Command interface {
	Run() error
}
