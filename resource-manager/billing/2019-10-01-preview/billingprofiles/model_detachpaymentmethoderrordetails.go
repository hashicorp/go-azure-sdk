package billingprofiles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetachPaymentMethodErrorDetails struct {
	Code    *DetachPaymentMethodEligibilityErrorCode `json:"code,omitempty"`
	Message *string                                  `json:"message,omitempty"`
}
