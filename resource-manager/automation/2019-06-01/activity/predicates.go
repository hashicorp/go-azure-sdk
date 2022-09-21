package activity

type ActivityOperationPredicate struct {
	Id   *string
	Name *string
}

func (p ActivityOperationPredicate) Matches(input Activity) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	return true
}
