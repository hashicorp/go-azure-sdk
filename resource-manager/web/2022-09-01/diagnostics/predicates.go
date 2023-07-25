package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnalysisDefinitionOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p AnalysisDefinitionOperationPredicate) Matches(input AnalysisDefinition) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Kind != nil && (input.Kind == nil && *p.Kind != *input.Kind) {
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

type DetectorDefinitionResourceOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p DetectorDefinitionResourceOperationPredicate) Matches(input DetectorDefinitionResource) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Kind != nil && (input.Kind == nil && *p.Kind != *input.Kind) {
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

type DetectorResponseOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p DetectorResponseOperationPredicate) Matches(input DetectorResponse) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Kind != nil && (input.Kind == nil && *p.Kind != *input.Kind) {
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

type DiagnosticCategoryOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p DiagnosticCategoryOperationPredicate) Matches(input DiagnosticCategory) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Kind != nil && (input.Kind == nil && *p.Kind != *input.Kind) {
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
