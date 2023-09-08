package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDevicesSummary struct {
	ComanagedCount                   *int64  `json:"comanagedCount,omitempty"`
	EligibleButNotAzureAdJoinedCount *int64  `json:"eligibleButNotAzureAdJoinedCount,omitempty"`
	EligibleCount                    *int64  `json:"eligibleCount,omitempty"`
	IneligibleCount                  *int64  `json:"ineligibleCount,omitempty"`
	NeedsOsUpdateCount               *int64  `json:"needsOsUpdateCount,omitempty"`
	ODataType                        *string `json:"@odata.type,omitempty"`
	ScheduledForEnrollmentCount      *int64  `json:"scheduledForEnrollmentCount,omitempty"`
}
