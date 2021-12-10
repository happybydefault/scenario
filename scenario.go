package scenario

import (
	"fmt"
	"strings"
	"testing"
)

type Scenario struct {
	t      *testing.T
	title  string
	givens []string
	when   string
	thens  []*Then
}

// Title returns a pointer to a new BDD Scenario.
func Title(title string) *Scenario {
	return &Scenario{title: title}
}

// Given returns a GivenFragment composed of a new Scenario.
func Given(description string) GivenFragment {
	s := &Scenario{}
	return s.Given(description)
}

func (s *Scenario) Run(t *testing.T) bool {
	t.Helper()

	fmt.Printf("%s\n\n", s)

	ch := make(chan bool, len(s.thens))
	go func() {
		defer close(ch)
		for _, then := range s.thens {
			ch <- t.Run(then.description, func(t *testing.T) {
				t.Helper()
				then.fn(t)
			})
		}
	}()

	success := true
	for v := range ch {
		if !v {
			success = false
		}
	}
	return success
}

func (s *Scenario) Then(description string, fn func(t *testing.T)) *Then {
	then := Then{
		description: description,
		fn:          fn,
	}

	s.thens = append(s.thens, &then)
	return &then
}

func (s *Scenario) Given(description string) GivenFragment {
	s.givens = append(s.givens, description)
	return GivenFragment{scenario: s}
}

// And is an alias for Then.
func (s *Scenario) And(description string, fn func(t *testing.T)) *Then {
	return s.Then(description, fn)
}

func (s *Scenario) String() string {
	givens := make([]string, len(s.givens))
	for i, given := range s.givens {
		word := "Given"
		if i > 0 {
			word = "And"
		}

		givens[i] = fmt.Sprintf("%s %s", word, given)
	}

	thens := make([]string, len(s.thens))
	for i, then := range s.thens {
		word := "Then"
		if i > 0 {
			word = "And"
		}

		thens[i] = fmt.Sprintf("%s %s", word, then.description)
	}

	return fmt.Sprintf("Scenario: %s\n%s\nWhen %s\n%s",
		s.title,
		strings.Join(givens, "\n"),
		s.when,
		strings.Join(thens, "\n"),
	)
}
