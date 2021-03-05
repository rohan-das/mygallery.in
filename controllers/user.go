package controllers

import (
	"fmt"
	"net/http"

	"mygallery.in/views"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// intial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using Controller")
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
