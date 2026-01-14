package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityOperationPredicate struct {
	Name   *string
	Reason *string
}

func (p CapabilityOperationPredicate) Matches(input Capability) bool {

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Reason != nil && (input.Reason == nil || *p.Reason != *input.Reason) {
		return false
	}

	return true
}

type QuotaUsageOperationPredicate struct {
	CurrentValue *int64
	Id           *string
	Limit        *int64
	Unit         *string
}

func (p QuotaUsageOperationPredicate) Matches(input QuotaUsage) bool {

	if p.CurrentValue != nil && (input.CurrentValue == nil || *p.CurrentValue != *input.CurrentValue) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Limit != nil && (input.Limit == nil || *p.Limit != *input.Limit) {
		return false
	}

	if p.Unit != nil && (input.Unit == nil || *p.Unit != *input.Unit) {
		return false
	}

	return true
}
