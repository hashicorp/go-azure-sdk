package customimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomImagePropertiesFromVM struct {
	LinuxOsInfo   *LinuxOsInfo   `json:"linuxOsInfo,omitempty"`
	SourceVMId    *string        `json:"sourceVmId,omitempty"`
	WindowsOsInfo *WindowsOsInfo `json:"windowsOsInfo,omitempty"`
}
