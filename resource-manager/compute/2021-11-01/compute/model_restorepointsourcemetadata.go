package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointSourceMetadata struct {
	DiagnosticsProfile *DiagnosticsProfile                 `json:"diagnosticsProfile"`
	HardwareProfile    *HardwareProfile                    `json:"hardwareProfile"`
	LicenseType        *string                             `json:"licenseType,omitempty"`
	Location           *string                             `json:"location,omitempty"`
	OsProfile          *OSProfile                          `json:"osProfile"`
	SecurityProfile    *SecurityProfile                    `json:"securityProfile"`
	StorageProfile     *RestorePointSourceVMStorageProfile `json:"storageProfile"`
	VmId               *string                             `json:"vmId,omitempty"`
}
