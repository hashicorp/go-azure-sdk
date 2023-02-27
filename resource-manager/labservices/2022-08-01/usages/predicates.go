// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package usages

type UsageOperationPredicate struct {
	CurrentValue *int64
	Id           *string
	Limit        *int64
}

func (p UsageOperationPredicate) Matches(input Usage) bool {

	if p.CurrentValue != nil && (input.CurrentValue == nil && *p.CurrentValue != *input.CurrentValue) {
		return false
	}

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Limit != nil && (input.Limit == nil && *p.Limit != *input.Limit) {
		return false
	}

	return true
}
