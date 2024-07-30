package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceKey struct {
	DeviceId    *string `json:"deviceId,omitempty"`
	KeyMaterial *string `json:"keyMaterial,omitempty"`
	KeyType     *string `json:"keyType,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
