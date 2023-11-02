package htmg

// Node is the most basic of children in a DOM and represents something that can
// be written to an io.Writer and possess children.
type Node interface {
	Writable

	Children() []Node
	PrependChild(child Node) Node
	AppendChild(child Node) Node
}
