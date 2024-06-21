package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailablePatchSummary struct {
	AssessmentActivityId          *string               `json:"assessmentActivityId,omitempty"`
	CriticalAndSecurityPatchCount *int64                `json:"criticalAndSecurityPatchCount,omitempty"`
	Error                         *ApiError             `json:"error,omitempty"`
	LastModifiedTime              *string               `json:"lastModifiedTime,omitempty"`
	OtherPatchCount               *int64                `json:"otherPatchCount,omitempty"`
	RebootPending                 *bool                 `json:"rebootPending,omitempty"`
	StartTime                     *string               `json:"startTime,omitempty"`
	Status                        *PatchOperationStatus `json:"status,omitempty"`
}
