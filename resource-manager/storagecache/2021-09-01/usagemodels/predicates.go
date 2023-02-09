package usagemodels

type UsageModelOperationPredicate struct {
	ModelName  *string
	TargetType *string
}

func (p UsageModelOperationPredicate) Matches(input UsageModel) bool {

	if p.ModelName != nil && (input.ModelName == nil && *p.ModelName != *input.ModelName) {
		return false
	}

	if p.TargetType != nil && (input.TargetType == nil && *p.TargetType != *input.TargetType) {
		return false
	}

	return true
}
