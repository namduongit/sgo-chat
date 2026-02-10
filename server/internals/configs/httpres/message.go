package httpres

type MessageResponse string

const (
	InternalServerError MessageResponse = "Internal Server Error"
	BadRequest          MessageResponse = "Bad Request"
	BodyMissing         MessageResponse = "Request Body Missing"
	NotFound            MessageResponse = "Resource Not Found"
	Unauthorized        MessageResponse = "Unauthorized Access"
	Forbidden           MessageResponse = "Forbidden Access"
	Success             MessageResponse = "Request Successful"
)
