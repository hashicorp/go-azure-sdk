package billingaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Enrollment struct {
	BillingCycle *string             `json:"billingCycle,omitempty"`
	Channel      *string             `json:"channel,omitempty"`
	CountryCode  *string             `json:"countryCode,omitempty"`
	Currency     *string             `json:"currency,omitempty"`
	EndDate      *string             `json:"endDate,omitempty"`
	Language     *string             `json:"language,omitempty"`
	Policies     *EnrollmentPolicies `json:"policies,omitempty"`
	StartDate    *string             `json:"startDate,omitempty"`
	Status       *string             `json:"status,omitempty"`
}
