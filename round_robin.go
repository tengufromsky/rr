package rr

import (
	"errors"
	"math/rand"
	"sync/atomic"
)

// RoundRobinString can separate some elements by percents, but only integer percents
// inner representation of elements must contain 100 elements for better and
// easier counter calculation
type RoundRobinString struct {
	counter  uint32
	elements []string
}

// StringElement contains info about element of the inner round robin queue
type StringElement struct {
	Percent int32
	Value   string
}

// Load change inner state of RoundRobinString as indicated parameters and
// reset inner counter
func (rr *RoundRobinString) Load(elements []StringElement) error {
	atomic.StoreUint32(&rr.counter, 0)
	el, err := createHosts(elements)
	if err != nil {
		return err
	}
	rr.elements = el
	return nil
}

// Next return next query item value
func (rr *RoundRobinString) Next() string {
	c := atomic.AddUint32(&rr.counter, 1) % 100
	return rr.elements[c]
}

// NewRoundRobinString create initialized RoundRobinString and return pointer
// to it, or error if parameters invalid
func NewRoundRobinString(elements []StringElement) (*RoundRobinString, error) {
	el, err := createHosts(elements)
	if err != nil {
		return nil, err
	}
	w := &RoundRobinString{elements: el}
	return w, nil
}

func createHosts(elements []StringElement) ([]string, error) {
	if err := checkPercents(elements); err != nil {
		return nil, err
	}

	var result = make([]string, 0, 100)
	for _, element := range elements {
		for i := int32(0); i < element.Percent; i++ {
			result = append(result, element.Value)
		}
	}

	// TODO replace shuffling to weighted spreading
	rand.Shuffle(100, func(i, j int) { result[i], result[j] = result[j], result[i] })

	return result, nil
}

func checkPercents(elements []StringElement) error {
	var sum int32
	for _, e := range elements {
		sum += e.Percent
	}
	if sum != 100.0 {
		return ErrWrongPercents
	}
	return nil
}

var (
	ErrWrongPercents = errors.New("summary of percentages not equal 100")
)
