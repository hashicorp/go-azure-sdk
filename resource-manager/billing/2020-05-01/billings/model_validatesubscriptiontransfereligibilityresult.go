package billings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateSubscriptionTransferEligibilityResult struct {
	ErrorDetails   *ValidateSubscriptionTransferEligibilityError `json:"errorDetails,omitempty"`
	IsMoveEligible *bool                                         `json:"isMoveEligible,omitempty"`
}
