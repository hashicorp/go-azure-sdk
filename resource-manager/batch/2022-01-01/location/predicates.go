package location

type SupportedSkuOperationPredicate struct {
	FamilyName *string
	Name       *string
}

func (p SupportedSkuOperationPredicate) Matches(input SupportedSku) bool {

	if p.FamilyName != nil && (input.FamilyName == nil && *p.FamilyName != *input.FamilyName) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	return true
}
