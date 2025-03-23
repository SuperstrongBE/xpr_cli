package interfaces

type Command interface {
	parse(args []string)
}
