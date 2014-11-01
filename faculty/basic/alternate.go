// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

type Alternate struct {
	flow []chan struct{} // flow tokens for both channels
}

func (a *Alternate) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	a.flow = make([]chan struct{}, 2)
	a.flow[0] = make(chan struct{}, 1)
	a.flow[1] = make(chan struct{}, 1)
	return nil
}

func (a *Alternate) OverCognize(eye *be.Eye, valve Name, value interface{}) {
	switch valve.(string) {
	case "SX":
		a.flow[0] <- struct{}{} // obtain token to send
		eye.Show("TX", value)
		<-a.flow[1] // grant token to other side
	case "SY":
		a.flow[1] <- struct{}{} // obtain token to send
		eye.Show("TY", value)
		<-a.flow[0] // grant token to other side
	case "TX", "TY":
	default:
		panic("invalid valve name on alternation gate")
	}
}
