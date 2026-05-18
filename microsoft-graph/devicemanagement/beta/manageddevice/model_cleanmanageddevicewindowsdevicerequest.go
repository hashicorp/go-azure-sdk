package manageddevice

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CleanManagedDeviceWindowsDeviceRequest struct {
	KeepUserData *bool `json:"keepUserData,omitempty"`
}
