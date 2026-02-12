package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceGroup struct {
	Description   *string   `json:"description,omitempty"`
	DisplayName   string    `json:"displayName"`
	Etag          *string   `json:"etag,omitempty"`
	Filter        string    `json:"filter"`
	Id            *string   `json:"id,omitempty"`
	Organizations *[]string `json:"organizations,omitempty"`
}
