package exception

type InvalidFormat struct {
	Args string
}

func (i *InvalidFormat) Message() string {
	return "Something went wrong parsing your request: " + i.Args
}
