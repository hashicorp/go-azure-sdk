package workspacemanagedsqlserverblobauditing

type ExtendedServerBlobAuditingPolicyOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ExtendedServerBlobAuditingPolicyOperationPredicate) Matches(input ExtendedServerBlobAuditingPolicy) bool {

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

type ServerBlobAuditingPolicyOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ServerBlobAuditingPolicyOperationPredicate) Matches(input ServerBlobAuditingPolicy) bool {

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
