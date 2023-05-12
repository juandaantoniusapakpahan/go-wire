//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabasePostgresql, NewDatabaseMonggoDb, NewDatabaseRepository)
	return nil
}

var foo = wire.NewSet(NewFooRepository, NewFooService)
var bar = wire.NewSet(NewBarRepository, NewBarService)

func InitializedServiceFooBar() *FooBarService {
	wire.Build(
		foo,
		bar,
		NewFooBarService,
	)
	return nil
}

var sayHelloSet = wire.NewSet(NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)))

func InitializedSayHello() *SayHelloService {
	wire.Build(sayHelloSet, NewSayHelloService)
	return nil
}
