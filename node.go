package htmg

// Node is the most basic of children in a DOM and represents something that can
// be written to an io.Writer and possess children.
type Node interface {
	Writable

	// Children retrieves the current slice of applied child nodes
	Children() []Node

	// Prepend adds the given children to the beginning of the children list
	Prepend(children ...Node) Node

	// Append adds the given children to the end of the children list
	Append(children ...Node) Node
}
