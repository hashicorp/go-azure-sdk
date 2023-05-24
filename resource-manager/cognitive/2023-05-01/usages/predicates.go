package usages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageOperationPredicate struct {
	CurrentValue  *float64
	Limit         *float64
	NextResetTime *string
	QuotaPeriod   *string
}

func (p UsageOperationPredicate) Matches(input Usage) bool {

	if p.CurrentValue != nil && (input.CurrentValue == nil && *p.CurrentValue != *input.CurrentValue) {
		return false
	}

	if p.Limit != nil && (input.Limit == nil && *p.Limit != *input.Limit) {
		return false
	}

	if p.NextResetTime != nil && (input.NextResetTime == nil && *p.NextResetTime != *input.NextResetTime) {
		return false
	}

	if p.QuotaPeriod != nil && (input.QuotaPeriod == nil && *p.QuotaPeriod != *input.QuotaPeriod) {
		return false
	}

	return true
}
