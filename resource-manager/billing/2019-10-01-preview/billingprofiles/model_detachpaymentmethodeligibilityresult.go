package billingprofiles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetachPaymentMethodEligibilityResult struct {
	ErrorDetails *[]DetachPaymentMethodErrorDetails `json:"errorDetails,omitempty"`
	IsEligible   *bool                              `json:"isEligible,omitempty"`
}
