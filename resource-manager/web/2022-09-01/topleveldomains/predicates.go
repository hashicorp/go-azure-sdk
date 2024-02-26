package topleveldomains

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TldLegalAgreementCollectionOperationPredicate struct {
	NextLink *string
}

func (p TldLegalAgreementCollectionOperationPredicate) Matches(input TldLegalAgreementCollection) bool {

	if p.NextLink != nil && (input.NextLink == nil || *p.NextLink != *input.NextLink) {
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
