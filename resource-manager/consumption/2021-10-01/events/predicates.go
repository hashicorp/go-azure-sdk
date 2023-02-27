// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package events

type EventSummaryOperationPredicate struct {
	ETag *string
	Id   *string
	Name *string
	Type *string
}

func (p EventSummaryOperationPredicate) Matches(input EventSummary) bool {

	if p.ETag != nil && (input.ETag == nil && *p.ETag != *input.ETag) {
		return false
	}

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
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
