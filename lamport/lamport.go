// Package lamport implements a lamport clock using the clock interface
package lamport

import (
	"strconv"
)

type Lamport struct {
	time int
}

func (l *Lamport) Increment() {
	l.time += 1
}

func (l Lamport) Send() string {
	l.Increment()
	return strconv.Itoa(l.time)
}

func (l *Lamport) Receive(s string) error {
	othertime64, err := strconv.ParseInt(s, 10, 64)
	othertime := int(othertime64)
	if err == nil && othertime > l.time {
		l.time = othertime
	}
	return err
}
