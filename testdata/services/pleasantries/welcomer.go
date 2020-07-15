package pleasantries

// Welcomer welcomes people.
type Welcomer interface {
	// Welcome makes a welcome message for somebody.
	Welcome(WelcomeRequest) WelcomeResponse
}

// WelcomeRequest is the request object for Welcomer.Welcome.
type WelcomeRequest struct {
	// To is the address of the person to send the message to.
	To string
	// Name is the name of the person to welcome.
	Name string
}

// WelcomeResponse is the response object for Welcomer.Welcome.
type WelcomeResponse struct {
	// Message is the welcome message.
	Message string
}
