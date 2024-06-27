package billingprofile

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteBillingProfileEligibilityResult struct {
	EligibilityDetails *[]DeleteBillingProfileEligibilityDetail `json:"eligibilityDetails,omitempty"`
	EligibilityStatus  *DeleteBillingProfileEligibilityStatus   `json:"eligibilityStatus,omitempty"`
}
