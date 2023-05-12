package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

// Function provider
func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{Error: isError}
}

type SimpleService struct {
	Repository *SimpleRepository
}

// Function provider
func NewSimpleService(respository *SimpleRepository) (*SimpleService, error) {

	if respository.Error {
		return nil, errors.New("failed create")
	} else {
		return &SimpleService{Repository: respository}, nil
	}

}
