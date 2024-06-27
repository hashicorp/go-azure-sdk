package billingprofile

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteBillingProfileEligibilityDetail struct {
	Code    *DeleteBillingProfileEligibilityCode `json:"code,omitempty"`
	Message *string                              `json:"message,omitempty"`
}
