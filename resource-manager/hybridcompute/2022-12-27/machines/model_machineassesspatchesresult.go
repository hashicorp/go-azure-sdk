package machines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineAssessPatchesResult struct {
	AssessmentActivityId                *string                              `json:"assessmentActivityId,omitempty"`
	AvailablePatchCountByClassification *AvailablePatchCountByClassification `json:"availablePatchCountByClassification,omitempty"`
	ErrorDetails                        *ErrorDetail                         `json:"errorDetails,omitempty"`
	LastModifiedDateTime                *string                              `json:"lastModifiedDateTime,omitempty"`
	OsType                              *OsType                              `json:"osType,omitempty"`
	PatchServiceUsed                    *PatchServiceUsed                    `json:"patchServiceUsed,omitempty"`
	RebootPending                       *bool                                `json:"rebootPending,omitempty"`
	StartDateTime                       *string                              `json:"startDateTime,omitempty"`
	StartedBy                           *PatchOperationStartedBy             `json:"startedBy,omitempty"`
	Status                              *PatchOperationStatus                `json:"status,omitempty"`
}
