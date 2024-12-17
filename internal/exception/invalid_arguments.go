package exception

type InvalidArguments struct {
	Args string
}

func (args *InvalidArguments) Message() string {
	return "Invalid arguments:" + args.Args
}
