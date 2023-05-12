package simple

type FooRepository struct {
}

func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type FooService struct {
	*FooRepository
}

// provider
func NewFooService(fooRepository *FooRepository) *FooService {
	return &FooService{FooRepository: fooRepository}
}
