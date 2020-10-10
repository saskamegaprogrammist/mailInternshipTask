package requestParser

type EmptyStruct struct{}

// request model

type Request struct {
	id       int
	resource string
	count    int
}
