package controllers

import (
	"fmt"
	"net/http"

	"github.com/igalsh/usegolang/views"
)

//Exported func
func NewUsers() *Users {
	return &Users{ // Constructor
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

//Exported struct
type Users struct {
	NewView *views.View
}

// This is used to render the form where a user can
// create a new user account.
//
// GET /signup
//Check why i calling to this function without  params
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

// This is used to process the signup form when a user
// tries to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a temporary response.")
}
