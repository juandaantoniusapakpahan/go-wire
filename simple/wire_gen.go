// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

import (
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializedService(isError bool) (*SimpleService, error) {
	simpleRepository := NewSimpleRepository(isError)
	simpleService, err := NewSimpleService(simpleRepository)
	if err != nil {
		return nil, err
	}
	return simpleService, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	databasePostgresql := NewDatabasePostgresql()
	databaseMonggoDb := NewDatabaseMonggoDb()
	databaseRepository := NewDatabaseRepository(databasePostgresql, databaseMonggoDb)
	return databaseRepository
}

func InitializedServiceFooBar() *FooBarService {
	fooRepository := NewFooRepository()
	fooService := NewFooService(fooRepository)
	barRepository := NewBarRepository()
	barService := NewBarService(barRepository)
	fooBarService := NewFooBarService(fooService, barService)
	return fooBarService
}

func InitializedSayHello() *SayHelloService {
	sayHelloImpl := NewSayHelloImpl()
	sayHelloService := NewSayHelloService(sayHelloImpl)
	return sayHelloService
}

// injector.go:

var foo = wire.NewSet(NewFooRepository, NewFooService)

var bar = wire.NewSet(NewBarRepository, NewBarService)

var sayHelloSet = wire.NewSet(NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)))