package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventSummaryOperationPredicate struct {
	ETag *string
	Id   *string
	Name *string
	Type *string
}

func (p EventSummaryOperationPredicate) Matches(input EventSummary) bool {

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

type LotSummaryOperationPredicate struct {
	ETag *string
	Id   *string
	Name *string
	Type *string
}

func (p LotSummaryOperationPredicate) Matches(input LotSummary) bool {

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

type MarketplaceOperationPredicate struct {
	Etag *string
	Id   *string
	Name *string
	Type *string
}

func (p MarketplaceOperationPredicate) Matches(input Marketplace) bool {

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
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

type ModernReservationTransactionOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ModernReservationTransactionOperationPredicate) Matches(input ModernReservationTransaction) bool {

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

type ReservationDetailOperationPredicate struct {
	Etag *string
	Id   *string
	Name *string
	Type *string
}

func (p ReservationDetailOperationPredicate) Matches(input ReservationDetail) bool {

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
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

type ReservationRecommendationOperationPredicate struct {
}

func (p ReservationRecommendationOperationPredicate) Matches(input ReservationRecommendation) bool {

	return true
}

type ReservationSummaryOperationPredicate struct {
	Etag *string
	Id   *string
	Name *string
	Type *string
}

func (p ReservationSummaryOperationPredicate) Matches(input ReservationSummary) bool {

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
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

type ReservationTransactionOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ReservationTransactionOperationPredicate) Matches(input ReservationTransaction) bool {

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
