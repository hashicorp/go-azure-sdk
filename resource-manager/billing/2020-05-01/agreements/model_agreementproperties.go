package agreements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementProperties struct {
	AcceptanceMode     *AcceptanceMode     `json:"acceptanceMode,omitempty"`
	AgreementLink      *string             `json:"agreementLink,omitempty"`
	BillingProfileInfo *BillingProfileInfo `json:"billingProfileInfo,omitempty"`
	Category           *Category           `json:"category,omitempty"`
	EffectiveDate      *string             `json:"effectiveDate,omitempty"`
	ExpirationDate     *string             `json:"expirationDate,omitempty"`
	Participants       *[]Participants     `json:"participants,omitempty"`
	Status             *string             `json:"status,omitempty"`
}
