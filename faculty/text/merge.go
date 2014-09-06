// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package text provides gates for manipulating text.
package text

import (
	"bytes"
	"io"
	// "log"
	"sync"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/see"
)

func init() {
	ns := faculty.Root.Refine("text")
	ns.AddTerminal(see.Name("Merge"), Merge{})
	ns.AddTerminal(see.Name("Form"), Form{})
}

// Merge …
type Merge struct{}

func (Merge) Materialize() be.Reflex {
	_Endo, _Exo := be.NewSynapse()
	firstEndo, firstExo := be.NewSynapse()
	secondEndo, secondExo := be.NewSynapse()
	thirdEndo, thirdExo := be.NewSynapse()
	go func() {
		h := &merge{
			ready: make(chan struct{}),
		}
		h.reply = _Endo.Focus(be.DontCognize)
		close(h.ready)
		firstEndo.Focus(func(v interface{}) { h.CognizeArm(0, v) })
		secondEndo.Focus(func(v interface{}) { h.CognizeArm(1, v) })
		thirdEndo.Focus(func(v interface{}) { h.CognizeArm(2, v) })
	}()
	return be.Reflex{
		"_":      _Exo,
		"First":  firstExo,
		"Second": secondExo,
		"Third":  thirdExo,
	}
}

type merge struct {
	ready chan struct{}
	reply *be.ReCognizer
	sync.Mutex
	arm [3]*bytes.Buffer
}

func (h *merge) CognizeArm(index int, v interface{}) {
	<-h.ready
	h.Lock()
	defer h.Unlock()
	switch t := v.(type) {
	case string:
		h.arm[index] = bytes.NewBufferString(t)
	case []byte:
		h.arm[index] = bytes.NewBuffer(t)
	case byte:
		h.arm[index] = bytes.NewBuffer([]byte{t})
	case rune:
		h.arm[index] = bytes.NewBuffer(nil)
		h.arm[index].WriteRune(t)
	case io.Reader:
		h.arm[index] = bytes.NewBuffer(nil)
		io.Copy(h.arm[index], t)
	default:
		panic("unsupported")
	}
	// merge
	if h.arm[0] == nil || h.arm[1] == nil || h.arm[2] == nil {
		return
	}
	var a bytes.Buffer
	a.Write(h.arm[0].Bytes())
	a.Write(h.arm[1].Bytes())
	a.Write(h.arm[2].Bytes())
	h.reply.ReCognize(a.String())
}
