package pleasantries

// Ignorer gets ignored by the tooling.
type Ignorer interface {
	Ignore(IgnoreRequest) IgnoreResponse
}

// IgnoreRequest should get ignored.
type IgnoreRequest struct{}

// IgnoreResponse should get ignored.
type IgnoreResponse struct{}
