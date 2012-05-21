// Package fbapp provides a Facebook Application type and some related
// utilities.
package fbapp

import (
	"flag"
)

type App interface {
	ID() uint64
	Secret() string
	Namespace() string
	SecretByte() []byte
}

type app struct {
	id         uint64
	secret     string
	namespace  string
	secretByte []byte
}

var Default = Flag("fbapp")

// Defines flags to configure an application.
func Flag(name string) App {
	app := &app{}
	flag.Uint64Var(&app.id, name+".id", 0, "Facebook application ID.")
	flag.StringVar(
		&app.secret, name+".secret", "", "Facebook application secret.")
	flag.StringVar(
		&app.namespace, name+".namespace", "", "Facebook application namespace.")
	return app
}

// Create a new App with the given configuration.
func New(id uint64, secret string, namespace string) App {
	return &app{
		id: id,
		secret: secret,
		namespace: namespace,
	}
}

func (a *app) ID() uint64 {
	return a.id
}

func (a *app) Secret() string {
	return a.secret
}

func (a *app) Namespace() string {
	return a.namespace
}

func (a *app) SecretByte() []byte {
	if a.secretByte == nil {
		a.secretByte = []byte(a.secret)
	}
	return a.secretByte
}
