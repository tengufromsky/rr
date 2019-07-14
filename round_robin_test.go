package rr

import (
	"testing"
)

func TestNewRoundRobin(t *testing.T) {
	// errors while create
	for _, tc := range []struct {
		name string
		in   []StringElement
	}{
		{
			name: "one parametrs with 101 percents",
			in: []StringElement{
				{Percent: 101, Value: "one and only"},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewRoundRobinString(tc.in)
			assertMustError(t, err)
		})
	}

	// positive scenario with truly weights
	for _, tc := range []struct {
		name string
		in   []StringElement
	}{
		{
			name: "one parametr with 100 percents",
			in: []StringElement{
				{Percent: 100, Value: "one and only"},
			},
		},
		{
			name: "two parametrs with equal percents",
			in: []StringElement{
				{Percent: 50, Value: "first"},
				{Percent: 50, Value: "second"},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewRoundRobinString(tc.in)
			assertNoError(t, err)
		})
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func assertMustError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
}
