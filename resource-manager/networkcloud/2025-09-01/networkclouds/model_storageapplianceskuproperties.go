package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageApplianceSkuProperties struct {
	CapacityGB *int64  `json:"capacityGB,omitempty"`
	Model      *string `json:"model,omitempty"`
}
