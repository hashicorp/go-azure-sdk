package importedwindowsautopilotdeviceidentity

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedWindowsAutopilotDeviceIdentity struct {
	AssignedUserPrincipalName *string                                      `json:"assignedUserPrincipalName,omitempty"`
	GroupTag                  *string                                      `json:"groupTag,omitempty"`
	HardwareIdentifier        *string                                      `json:"hardwareIdentifier,omitempty"`
	Id                        *string                                      `json:"id,omitempty"`
	ImportId                  *string                                      `json:"importId,omitempty"`
	ODataType                 *string                                      `json:"@odata.type,omitempty"`
	ProductKey                *string                                      `json:"productKey,omitempty"`
	SerialNumber              *string                                      `json:"serialNumber,omitempty"`
	State                     *ImportedWindowsAutopilotDeviceIdentityState `json:"state,omitempty"`
}
