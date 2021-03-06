package configurations

type ServerConfigurationOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ServerConfigurationOperationPredicate) Matches(input ServerConfiguration) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
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

type ServerGroupConfigurationOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ServerGroupConfigurationOperationPredicate) Matches(input ServerGroupConfiguration) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
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
