package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedWindowsAutopilotDeviceIdentityUpload struct {
	CreatedDateTimeUtc *string                                             `json:"createdDateTimeUtc,omitempty"`
	DeviceIdentities   *[]ImportedWindowsAutopilotDeviceIdentity           `json:"deviceIdentities,omitempty"`
	Id                 *string                                             `json:"id,omitempty"`
	ODataType          *string                                             `json:"@odata.type,omitempty"`
	Status             *ImportedWindowsAutopilotDeviceIdentityUploadStatus `json:"status,omitempty"`
}
