package triggers

type TriggerOperationPredicate struct {
}

func (p TriggerOperationPredicate) Matches(input Trigger) bool {

	return true
}
