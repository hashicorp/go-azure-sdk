package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertOperationPredicate struct {
	ETag *string
	Id   *string
	Name *string
	Type *string
}

func (p AlertOperationPredicate) Matches(input Alert) bool {

	if p.ETag != nil && (input.ETag == nil || *p.ETag != *input.ETag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
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

type BenefitRecommendationModelOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p BenefitRecommendationModelOperationPredicate) Matches(input BenefitRecommendationModel) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
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

type BenefitUtilizationSummaryOperationPredicate struct {
}

func (p BenefitUtilizationSummaryOperationPredicate) Matches(input BenefitUtilizationSummary) bool {

	return true
}

type DimensionOperationPredicate struct {
	ETag     *string
	Id       *string
	Location *string
	Name     *string
	Sku      *string
	Type     *string
}

func (p DimensionOperationPredicate) Matches(input Dimension) bool {

	if p.ETag != nil && (input.ETag == nil || *p.ETag != *input.ETag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && (input.Location == nil || *p.Location != *input.Location) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Sku != nil && (input.Sku == nil || *p.Sku != *input.Sku) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
