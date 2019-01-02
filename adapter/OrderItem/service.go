package OrderItem

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Store an
func (s *Service) GetAllOrderCount() int {
	return s.repo.GetAllOrderCount()
}
