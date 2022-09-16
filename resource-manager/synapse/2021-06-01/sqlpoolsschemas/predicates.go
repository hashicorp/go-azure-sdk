package sqlpoolsschemas

type ResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ResourceOperationPredicate) Matches(input Resource) bool {

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

type SqlPoolVulnerabilityAssessmentOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p SqlPoolVulnerabilityAssessmentOperationPredicate) Matches(input SqlPoolVulnerabilityAssessment) bool {

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

type VulnerabilityAssessmentScanRecordOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p VulnerabilityAssessmentScanRecordOperationPredicate) Matches(input VulnerabilityAssessmentScanRecord) bool {

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
