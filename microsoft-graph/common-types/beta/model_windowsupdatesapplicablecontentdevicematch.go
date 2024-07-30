package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesApplicableContentDeviceMatch struct {
	DeviceId      *string   `json:"deviceId,omitempty"`
	ODataType     *string   `json:"@odata.type,omitempty"`
	RecommendedBy *[]string `json:"recommendedBy,omitempty"`
}
