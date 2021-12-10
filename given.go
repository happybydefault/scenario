package scenario

type GivenFragment struct {
	scenario *Scenario
}

func (g GivenFragment) And(description string) GivenFragment {
	return g.scenario.Given(description)
}

func (g GivenFragment) When(description string) *Scenario {
	g.scenario.when = description
	return g.scenario
}
