package pleasantries

// Ignorer gets ignored by the tooling.
type Ignorer interface {
	Ignore(IgnoreRequest) IgnoreResponse
}

type IgnoreRequest struct{}

type IgnoreResponse struct{}
