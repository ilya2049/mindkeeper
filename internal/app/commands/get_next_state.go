package commands

type GetNextStateCommandHandler struct {
}

func NewGetNextStateCommandHandler() GetNextStateCommandHandler {
	return GetNextStateCommandHandler{}
}

func (h GetNextStateCommandHandler) Handle() string {
	return ""
}
