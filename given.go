package scenario

type Given struct {
	scenario *Scenario
}

func (g Given) Given(description string) Given {
	return g.scenario.Given(description)
}

func (g Given) And(description string) Given {
	return g.Given(description)
}

func (g Given) When(description string) *Scenario {
	g.scenario.when = description
	return g.scenario
}
