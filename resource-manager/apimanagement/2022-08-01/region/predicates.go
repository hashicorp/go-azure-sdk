package region

type RegionContractOperationPredicate struct {
	IsDeleted      *bool
	IsMasterRegion *bool
	Name           *string
}

func (p RegionContractOperationPredicate) Matches(input RegionContract) bool {

	if p.IsDeleted != nil && (input.IsDeleted == nil && *p.IsDeleted != *input.IsDeleted) {
		return false
	}

	if p.IsMasterRegion != nil && (input.IsMasterRegion == nil && *p.IsMasterRegion != *input.IsMasterRegion) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	return true
}
