package sourcecontrolsyncjobstreams

type SourceControlSyncJobStreamOperationPredicate struct {
	Id *string
}

func (p SourceControlSyncJobStreamOperationPredicate) Matches(input SourceControlSyncJobStream) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	return true
}
