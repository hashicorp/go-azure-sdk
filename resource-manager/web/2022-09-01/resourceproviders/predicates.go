package resourceproviders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingMeterOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p BillingMeterOperationPredicate) Matches(input BillingMeter) bool {

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

type CustomHostnameSitesOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p CustomHostnameSitesOperationPredicate) Matches(input CustomHostnameSites) bool {

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

type GeoRegionOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p GeoRegionOperationPredicate) Matches(input GeoRegion) bool {

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

type IdentifierOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p IdentifierOperationPredicate) Matches(input Identifier) bool {

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

type PremierAddOnOfferOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p PremierAddOnOfferOperationPredicate) Matches(input PremierAddOnOffer) bool {

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

type SourceControlOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p SourceControlOperationPredicate) Matches(input SourceControl) bool {

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
