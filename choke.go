// Package choke creates a choke point where actions being performed
// by multiple goroutines are serialized
package choke

import (
	"golang.org/x/net/context"
)

type Choke struct {
	Ch chan func()
}

// New creates a new choke point. The depth sets the
// cap() of the channel. In some use valid cases this
// may need to > 0 to prevent deadlocks.
func New(depth int) (c *Choke) {
	c = &Choke{
		Ch: make(chan func(), depth),
	}
	return // implicit
}

// Do executes the function passed in and is guaranteed
// not to collide with any other Do() in progress that is
// under the same Choke.
func (c *Choke) Do(f func() error) error {
	respCh := make(chan error)
	c.Ch <- func() {
		respCh <- f()
		return //
	}
	return <-respCh
}

// Doer is the goroutine that pulls the actions from
// the channel and executest them.
func (c *Choke) Doer(ctx context.Context) {
LOOP:
	for {
		select {
		case req, ok := <-c.Ch:
			if !ok {
				break LOOP
			}
			req()
		case <-ctx.Done():
			break LOOP
		}
	}
}
