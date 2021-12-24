package scenario

import (
	"flag"
	"fmt"
	"strings"
	"testing"

	"github.com/gookit/color"
)

var pretty bool

func init() {
	flag.BoolVar(&pretty, "scenario.pretty", false, "Enable pretty output.")
}

type Scenario struct {
	t      *testing.T
	title  string
	givens []string
	when   string
	thens  []*Then
}

// New returns a pointer to a new BDD Scenario.
func New(title string) *Scenario {
	return &Scenario{title: title}
}

func (s *Scenario) Run(t *testing.T) bool {
	t.Helper()

	str := s.String()
	if pretty {
		str = color.FgLightBlue.Sprint(s)
	}
	fmt.Printf("%s\n\n", str)

	return t.Run(s.title, func(t *testing.T) {
		t.Helper()
		for _, then := range s.thens {
			t.Run(then.description, func(t *testing.T) {
				t.Helper()
				then.fn(t)
			})
		}
	})
}

func (s *Scenario) Then(description string, fn func(t *testing.T)) *Then {
	then := Then{
		description: description,
		fn:          fn,
	}

	s.thens = append(s.thens, &then)
	return &then
}

func (s *Scenario) Given(description string) Given {
	s.givens = append(s.givens, description)
	return Given{scenario: s}
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
