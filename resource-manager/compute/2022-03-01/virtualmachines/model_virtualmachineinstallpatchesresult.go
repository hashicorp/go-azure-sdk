package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineInstallPatchesResult struct {
	Error                     *ApiError                  `json:"error,omitempty"`
	ExcludedPatchCount        *int64                     `json:"excludedPatchCount,omitempty"`
	FailedPatchCount          *int64                     `json:"failedPatchCount,omitempty"`
	InstallationActivityId    *string                    `json:"installationActivityId,omitempty"`
	InstalledPatchCount       *int64                     `json:"installedPatchCount,omitempty"`
	MaintenanceWindowExceeded *bool                      `json:"maintenanceWindowExceeded,omitempty"`
	NotSelectedPatchCount     *int64                     `json:"notSelectedPatchCount,omitempty"`
	Patches                   *[]PatchInstallationDetail `json:"patches,omitempty"`
	PendingPatchCount         *int64                     `json:"pendingPatchCount,omitempty"`
	RebootStatus              *VMGuestPatchRebootStatus  `json:"rebootStatus,omitempty"`
	StartDateTime             *string                    `json:"startDateTime,omitempty"`
	Status                    *PatchOperationStatus      `json:"status,omitempty"`
}
