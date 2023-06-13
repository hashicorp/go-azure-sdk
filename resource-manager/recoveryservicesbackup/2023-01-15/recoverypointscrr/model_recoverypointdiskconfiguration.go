package recoverypointscrr

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointDiskConfiguration struct {
	ExcludedDiskList              *[]DiskInformation `json:"excludedDiskList,omitempty"`
	IncludedDiskList              *[]DiskInformation `json:"includedDiskList,omitempty"`
	NumberOfDisksAttachedToVM     *int64             `json:"numberOfDisksAttachedToVm,omitempty"`
	NumberOfDisksIncludedInBackup *int64             `json:"numberOfDisksIncludedInBackup,omitempty"`
}
