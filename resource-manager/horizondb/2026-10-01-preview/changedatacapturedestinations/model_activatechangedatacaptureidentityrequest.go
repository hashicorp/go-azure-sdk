package changedatacapturedestinations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivateChangeDataCaptureIdentityRequest struct {
	DrainExistingSessions        *bool  `json:"drainExistingSessions,omitempty"`
	IdentityAssignmentResourceId string `json:"identityAssignmentResourceId"`
}
