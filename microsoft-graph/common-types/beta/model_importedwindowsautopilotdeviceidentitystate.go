package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedWindowsAutopilotDeviceIdentityState struct {
	DeviceErrorCode      *int64                                                         `json:"deviceErrorCode,omitempty"`
	DeviceErrorName      *string                                                        `json:"deviceErrorName,omitempty"`
	DeviceImportStatus   *ImportedWindowsAutopilotDeviceIdentityStateDeviceImportStatus `json:"deviceImportStatus,omitempty"`
	DeviceRegistrationId *string                                                        `json:"deviceRegistrationId,omitempty"`
	ODataType            *string                                                        `json:"@odata.type,omitempty"`
}
