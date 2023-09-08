package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCGalleryImageOperationPredicate struct {
	DisplayName      *string
	EndDate          *string
	ExpirationDate   *string
	Id               *string
	ODataType        *string
	Offer            *string
	OfferDisplayName *string
	Publisher        *string
	RecommendedSku   *string
	SizeInGB         *int64
	Sku              *string
	SkuDisplayName   *string
	StartDate        *string
}

func (p CloudPCGalleryImageOperationPredicate) Matches(input CloudPCGalleryImage) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EndDate != nil && (input.EndDate == nil || *p.EndDate != *input.EndDate) {
		return false
	}

	if p.ExpirationDate != nil && (input.ExpirationDate == nil || *p.ExpirationDate != *input.ExpirationDate) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Offer != nil && (input.Offer == nil || *p.Offer != *input.Offer) {
		return false
	}

	if p.OfferDisplayName != nil && (input.OfferDisplayName == nil || *p.OfferDisplayName != *input.OfferDisplayName) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	if p.RecommendedSku != nil && (input.RecommendedSku == nil || *p.RecommendedSku != *input.RecommendedSku) {
		return false
	}

	if p.SizeInGB != nil && (input.SizeInGB == nil || *p.SizeInGB != *input.SizeInGB) {
		return false
	}

	if p.Sku != nil && (input.Sku == nil || *p.Sku != *input.Sku) {
		return false
	}

	if p.SkuDisplayName != nil && (input.SkuDisplayName == nil || *p.SkuDisplayName != *input.SkuDisplayName) {
		return false
	}

	if p.StartDate != nil && (input.StartDate == nil || *p.StartDate != *input.StartDate) {
		return false
	}

	return true
}
