package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WatermarkProtectionValues struct {
	IsEnabledForContentSharing *bool   `json:"isEnabledForContentSharing,omitempty"`
	IsEnabledForVideo          *bool   `json:"isEnabledForVideo,omitempty"`
	ODataType                  *string `json:"@odata.type,omitempty"`
}
