// Package twofer implements a simple function to share things.
package twofer

// ShareWith returns a string with a message about sharing with the given name.
// If name is empty, it uses "you" as the recipient.
func ShareWith(name string) string {
    if name == "" {
        name = "you"
    }
    return "One for " + name + ", one for me."
}
