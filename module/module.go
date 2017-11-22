// Package module implement some function
package module

// PublicType is
type PublicType struct {
	Name  string
	value int
}

type NewType struct {
	PublicType
	email string
}

type InterfaceOne interface {
	One() string
}

type InterfaceTwo interface {
	Two() string
}

type InterfaceThree interface {
	InterfaceOne
	InterfaceTwo
}

type privateType struct {
	name  string
	value int
}

func privateFunc() {
	// do something
	return
}

// PublicFunc do this
func PublicFunc() {
	return
}
