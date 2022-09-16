package workspacemanagedsqlserversqlusages

type ServerUsageOperationPredicate struct {
	CurrentValue  *float64
	DisplayName   *string
	Limit         *float64
	Name          *string
	NextResetTime *string
	ResourceName  *string
	Unit          *string
}

func (p ServerUsageOperationPredicate) Matches(input ServerUsage) bool {

	if p.CurrentValue != nil && (input.CurrentValue == nil && *p.CurrentValue != *input.CurrentValue) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil && *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Limit != nil && (input.Limit == nil && *p.Limit != *input.Limit) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.NextResetTime != nil && (input.NextResetTime == nil && *p.NextResetTime != *input.NextResetTime) {
		return false
	}

	if p.ResourceName != nil && (input.ResourceName == nil && *p.ResourceName != *input.ResourceName) {
		return false
	}

	if p.Unit != nil && (input.Unit == nil && *p.Unit != *input.Unit) {
		return false
	}

	return true
}
