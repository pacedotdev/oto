package pleasantries

type Welcomer interface {
	Welcome(WelcomeRequest) WelcomeResponse
}

type WelcomeRequest struct {
	To   string
	Name string
}

type WelcomeResponse struct {
	Message string
}
