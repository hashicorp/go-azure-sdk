package topleveldomains

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TldLegalAgreementOperationPredicate struct {
	AgreementKey *string
	Content      *string
	Title        *string
	Url          *string
}

func (p TldLegalAgreementOperationPredicate) Matches(input TldLegalAgreement) bool {

	if p.AgreementKey != nil && *p.AgreementKey != input.AgreementKey {
		return false
	}

	if p.Content != nil && *p.Content != input.Content {
		return false
	}

	if p.Title != nil && *p.Title != input.Title {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	return true
}

type TopLevelDomainOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p TopLevelDomainOperationPredicate) Matches(input TopLevelDomain) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Kind != nil && (input.Kind == nil || *p.Kind != *input.Kind) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
