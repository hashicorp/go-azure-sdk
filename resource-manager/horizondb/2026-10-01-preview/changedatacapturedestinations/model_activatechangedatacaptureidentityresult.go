package changedatacapturedestinations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivateChangeDataCaptureIdentityResult struct {
	PreviousIdentityAssignmentResourceId *string `json:"previousIdentityAssignmentResourceId,omitempty"`
	Status                               Status  `json:"status"`
	TargetIdentityAssignmentResourceId   string  `json:"targetIdentityAssignmentResourceId"`
}
