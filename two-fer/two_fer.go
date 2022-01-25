// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

const you = "you"

/*
Two-fer or 2-fer is short for two for one.
One for you and one for me.

Given a name, return a string with the message:
One for name, one for me.
Where "name" is the given name.
However, if the name is missing, return the string:
One for you, one for me.
*/
func ShareWith(name string) string {
	var who string

	switch {
	case len(name) > 0:
		who = name
	default:
		who = you
	}

	return "One for " + who + ", one for me."
}
