package dig

import (
	"log"

	_dig "go.uber.org/dig"
)

var container *_dig.Container = _dig.New()

func Provide(constructor interface{}) {
	err := container.Provide(constructor)

	if err != nil {
		log.Panic("failed to provide dependency:", err)
	}
}

func Invoke(function interface{}) error {
	err := container.Invoke(function)

	if err != nil {
		log.Panic("failed to invoke function:", err)
	}

	return err
}
