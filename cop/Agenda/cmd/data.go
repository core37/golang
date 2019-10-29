package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type User struct {
	Name     string
	Password string
	Email    string
	Phone    string
}

var users []User
var currentUser User
