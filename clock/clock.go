// Package clock implements a library for syncing distributed processes
package clock

// Clock is the interface that wraps the basic methods required of any clock object
//
// Increment takes one step forward in time
// Send produces a string used for syncing by other clock objects
// Receive consumes a string produced by other clock objects and updates the internal
// state of the clock object
type Clock interface {
	Increment()
	Send() string
	Sync(string) error
}
