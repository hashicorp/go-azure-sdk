package domains

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainPurchaseConsent struct {
	AgreedAt      *string   `json:"agreedAt,omitempty"`
	AgreedBy      *string   `json:"agreedBy,omitempty"`
	AgreementKeys *[]string `json:"agreementKeys,omitempty"`
}
