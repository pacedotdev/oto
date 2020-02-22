package pleasantries

type Ignorer interface {
	Ignore(IgnoreRequest) IgnoreResponse
}

type IgnoreRequest struct{}

type IgnoreResponse struct{}
