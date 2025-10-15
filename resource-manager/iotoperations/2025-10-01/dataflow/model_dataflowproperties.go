package dataflow

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowProperties struct {
	HealthState            *ResourceHealthState `json:"healthState,omitempty"`
	Mode                   *OperationalMode     `json:"mode,omitempty"`
	Operations             []DataflowOperation  `json:"operations"`
	ProvisioningState      *ProvisioningState   `json:"provisioningState,omitempty"`
	RequestDiskPersistence *OperationalMode     `json:"requestDiskPersistence,omitempty"`
}
