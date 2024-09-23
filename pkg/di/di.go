package di

import (
	"log"
	"net/http"

	"github.com/samber/do"
)

func Provide[T any](provider func() T) {
	do.Provide(nil, func(*do.Injector) (T, error) {
		return provider(), nil
	})
}

func ProvideNamed[T any](name string, provider func() T) {
	do.ProvideNamed(nil, name, func(*do.Injector) (T, error) {
		return provider(), nil
	})
}

func Invoke[T any]() T {
	return do.MustInvoke[T](nil)
}

func InvokeNamed[T any](name string) T {
	return do.MustInvokeNamed[T](nil, name)
}

func ProvideHTTP(path string, provider func() http.Handler) {
	ProvideNamed(path, provider)
}

func RegisterHTTP(path string) {
	log.Println("serving HTTP: " + path)
	http.Handle(path, InvokeNamed[http.Handler](path))
}
