package usages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageOperationPredicate struct {
	CurrentValue   *int64
	Id             *string
	Limit          *int64
	RequestedLimit *int64
	Type           *string
	Unit           *string
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

	if p.RequestedLimit != nil && (input.RequestedLimit == nil && *p.RequestedLimit != *input.RequestedLimit) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	if p.Unit != nil && (input.Unit == nil && *p.Unit != *input.Unit) {
		return false
	}

	return true
}
