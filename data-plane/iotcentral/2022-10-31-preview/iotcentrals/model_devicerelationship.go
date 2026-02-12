package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceRelationship struct {
	Id     *string `json:"id,omitempty"`
	Source *string `json:"source,omitempty"`
	Target *string `json:"target,omitempty"`
}
