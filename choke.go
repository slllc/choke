// Package choke creates a choke point where actions being performed
// by multiple goroutines are serialized
package choke

import (
	"golang.org/x/net/context"
)

type Choke struct {
	Ch    chan func()
	width int
}

// New creates a new choke point.
func New(width int) (c *Choke) {
	c = &Choke{
		Ch:    make(chan func()),
		width: width,
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

func (c *Choke) Start(ctx context.Context) *Choke {
	for i := 0; i < c.width; i++ {
		go c.Doer1(ctx)
	}
	return c
}

// Doer is the goroutine that pulls the actions from
// the channel and executest them.
func (c *Choke) Doer1(ctx context.Context) {
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
