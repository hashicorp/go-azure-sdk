package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDevicesSummaryOperationPredicate struct {
	ComanagedCount                   *int64
	EligibleButNotAzureAdJoinedCount *int64
	EligibleCount                    *int64
	IneligibleCount                  *int64
	NeedsOsUpdateCount               *int64
	ODataType                        *string
	ScheduledForEnrollmentCount      *int64
}

func (p ComanagementEligibleDevicesSummaryOperationPredicate) Matches(input ComanagementEligibleDevicesSummary) bool {

	if p.ComanagedCount != nil && (input.ComanagedCount == nil || *p.ComanagedCount != *input.ComanagedCount) {
		return false
	}

	if p.EligibleButNotAzureAdJoinedCount != nil && (input.EligibleButNotAzureAdJoinedCount == nil || *p.EligibleButNotAzureAdJoinedCount != *input.EligibleButNotAzureAdJoinedCount) {
		return false
	}

	if p.EligibleCount != nil && (input.EligibleCount == nil || *p.EligibleCount != *input.EligibleCount) {
		return false
	}

	if p.IneligibleCount != nil && (input.IneligibleCount == nil || *p.IneligibleCount != *input.IneligibleCount) {
		return false
	}

	if p.NeedsOsUpdateCount != nil && (input.NeedsOsUpdateCount == nil || *p.NeedsOsUpdateCount != *input.NeedsOsUpdateCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ScheduledForEnrollmentCount != nil && (input.ScheduledForEnrollmentCount == nil || *p.ScheduledForEnrollmentCount != *input.ScheduledForEnrollmentCount) {
		return false
	}

	return true
}
