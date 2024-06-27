package billingaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PaymentTermsEligibilityDetail struct {
	Code    *PaymentTermsEligibilityCode `json:"code,omitempty"`
	Message *string                      `json:"message,omitempty"`
}
