package exception

type ProcessingException struct {
	Args string
}

func (e *ProcessingException) Message() string {
	return "Something went wrong wrong processing your request: " + e.Args
}
