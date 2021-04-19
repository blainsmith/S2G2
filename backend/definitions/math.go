package definitions

type MathService interface {
	Add(AddRequest) AddResponse
	Subtract(SubtractRequest) SubtractResponse
}

type AddRequest struct {
	A float64
	B float64
}

type AddResponse struct {
	Result float64
}

type SubtractRequest struct {
	A float64
	B float64
}

type SubtractResponse struct {
	Result float64
}
