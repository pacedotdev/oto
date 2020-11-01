package pleasantries

type StrangeTypesService interface {
	DoSomethingStrange(DoSomethingStrangeRequest) DoSomethingStrangeResponse
}

type DoSomethingStrangeRequest struct {
	Anything interface{}
}

type DoSomethingStrangeResponse struct {
	Value interface{}
	Size  int
}
