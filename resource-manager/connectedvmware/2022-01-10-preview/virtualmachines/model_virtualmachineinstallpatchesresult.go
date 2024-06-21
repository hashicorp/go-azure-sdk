package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineInstallPatchesResult struct {
	ErrorDetails              *ErrorDetail              `json:"errorDetails,omitempty"`
	ExcludedPatchCount        *int64                    `json:"excludedPatchCount,omitempty"`
	FailedPatchCount          *int64                    `json:"failedPatchCount,omitempty"`
	InstallationActivityId    *string                   `json:"installationActivityId,omitempty"`
	InstalledPatchCount       *int64                    `json:"installedPatchCount,omitempty"`
	LastModifiedDateTime      *string                   `json:"lastModifiedDateTime,omitempty"`
	MaintenanceWindowExceeded *bool                     `json:"maintenanceWindowExceeded,omitempty"`
	NotSelectedPatchCount     *int64                    `json:"notSelectedPatchCount,omitempty"`
	OsType                    *OsTypeUM                 `json:"osType,omitempty"`
	PatchServiceUsed          *PatchServiceUsed         `json:"patchServiceUsed,omitempty"`
	PendingPatchCount         *int64                    `json:"pendingPatchCount,omitempty"`
	RebootStatus              *VMGuestPatchRebootStatus `json:"rebootStatus,omitempty"`
	StartDateTime             *string                   `json:"startDateTime,omitempty"`
	StartedBy                 *PatchOperationStartedBy  `json:"startedBy,omitempty"`
	Status                    *PatchOperationStatus     `json:"status,omitempty"`
}
