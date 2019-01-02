package OrderItem

//Reader read from db
type Reader interface {
	GetAllOrderCount() int
}

type Writer interface {
}

//Repository db interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface (Service for access)
type UseCase interface {
	Reader
	Writer
}
