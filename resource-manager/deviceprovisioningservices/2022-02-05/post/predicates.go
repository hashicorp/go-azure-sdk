package post

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOperationPredicate struct {
	KeyName      *string
	PrimaryKey   *string
	SecondaryKey *string
}

func (p SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOperationPredicate) Matches(input SharedAccessSignatureAuthorizationRuleAccessRightsDescription) bool {

	if p.KeyName != nil && *p.KeyName != input.KeyName {
		return false
	}

	if p.PrimaryKey != nil && (input.PrimaryKey == nil && *p.PrimaryKey != *input.PrimaryKey) {
		return false
	}

	if p.SecondaryKey != nil && (input.SecondaryKey == nil && *p.SecondaryKey != *input.SecondaryKey) {
		return false
	}

	return true
}
