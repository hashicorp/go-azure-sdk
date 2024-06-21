package caches

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CacheUpgradeStatus struct {
	CurrentFirmwareVersion *string             `json:"currentFirmwareVersion,omitempty"`
	FirmwareUpdateDeadline *string             `json:"firmwareUpdateDeadline,omitempty"`
	FirmwareUpdateStatus   *FirmwareStatusType `json:"firmwareUpdateStatus,omitempty"`
	LastFirmwareUpdate     *string             `json:"lastFirmwareUpdate,omitempty"`
	PendingFirmwareVersion *string             `json:"pendingFirmwareVersion,omitempty"`
}
