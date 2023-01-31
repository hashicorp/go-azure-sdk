package roles

type RoleOperationPredicate struct {
}

func (p RoleOperationPredicate) Matches(input Role) bool {

	return true
}
