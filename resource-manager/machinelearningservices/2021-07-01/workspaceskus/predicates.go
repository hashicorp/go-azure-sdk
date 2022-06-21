package workspaceskus

type WorkspaceSkuOperationPredicate struct {
	Name         *string
	ResourceType *string
	Tier         *string
}

func (p WorkspaceSkuOperationPredicate) Matches(input WorkspaceSku) bool {

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.ResourceType != nil && (input.ResourceType == nil && *p.ResourceType != *input.ResourceType) {
		return false
	}

	if p.Tier != nil && (input.Tier == nil && *p.Tier != *input.Tier) {
		return false
	}

	return true
}
