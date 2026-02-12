package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceTemplate struct {
	CapabilityModel interface{} `json:"capabilityModel"`
	Description     *string     `json:"description,omitempty"`
	DisplayName     *string     `json:"displayName,omitempty"`
	Etag            *string     `json:"etag,omitempty"`
	Id              *string     `json:"@id,omitempty"`
	Type            []string    `json:"@type"`
}
