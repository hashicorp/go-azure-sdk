package netappresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegionInfoResourceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p RegionInfoResourceOperationPredicate) Matches(input RegionInfoResource) bool {

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

type SubscriptionQuotaItemOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p SubscriptionQuotaItemOperationPredicate) Matches(input SubscriptionQuotaItem) bool {

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

type UsageResultOperationPredicate struct {
	Id *string
}

func (p UsageResultOperationPredicate) Matches(input UsageResult) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	return true
}
