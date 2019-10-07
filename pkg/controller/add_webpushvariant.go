package controller

import "github.com/aerogear/unifiedpush-operator/pkg/controller/webpushvariant"

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, webpushvariant.Add)
}
