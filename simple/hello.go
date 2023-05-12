package simple

// binding interface

type SayHello interface {
	Hello(name string) string
}

type SayHelloImpl struct {
}

func (sayhello *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

type SayHelloService struct {
	SayHello
}

func NewSayHelloService(sayhello SayHello) *SayHelloService {
	return &SayHelloService{
		SayHello: sayhello,
	}
}
