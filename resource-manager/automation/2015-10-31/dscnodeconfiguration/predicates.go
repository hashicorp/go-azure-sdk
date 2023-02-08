package dscnodeconfiguration

type DscNodeConfigurationOperationPredicate struct {
	CreationTime     *string
	Id               *string
	LastModifiedTime *string
	Name             *string
	Type             *string
}

func (p DscNodeConfigurationOperationPredicate) Matches(input DscNodeConfiguration) bool {

	if p.CreationTime != nil && (input.CreationTime == nil && *p.CreationTime != *input.CreationTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedTime != nil && (input.LastModifiedTime == nil && *p.LastModifiedTime != *input.LastModifiedTime) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}
