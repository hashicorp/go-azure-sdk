package job

type JobOperationPredicate struct {
	Id *string
}

func (p JobOperationPredicate) Matches(input Job) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	return true
}
