package actor

type Decider func(child *PID, cause interface{}) Directive

type SupervisionStrategy interface {
	Handle(child *PID, cause interface{}) Directive
}

type DefaultStrategy struct {
	maxNrOfRetries              int
	withinTimeRangeMilliseconds int
	decider                     Decider
}

func (strategy *DefaultStrategy) Handle(child *PID, reason interface{}) Directive {
	return strategy.decider(child, reason)
}

func NewDefaultStrategy(maxNrOfRetries int, withinTimeRangeMilliseconds int, decider Decider) SupervisionStrategy {
	return &DefaultStrategy{
		maxNrOfRetries:              maxNrOfRetries,
		withinTimeRangeMilliseconds: withinTimeRangeMilliseconds,
		decider:                     decider,
	}
}

func DefaultDecider(child *PID, reason interface{}) Directive {
	return Directive_RestartDirective
}

func DefaultSupervisionStrategy() SupervisionStrategy {
	return NewDefaultStrategy(10, 30000, DefaultDecider)
}
