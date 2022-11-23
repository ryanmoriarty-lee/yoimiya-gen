package contract

const IDKey = "yoimiya:id"

type IDService interface {
	NewID() string
}
