package standbyvirtualmachinepoolruntimeviews

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandbyVirtualMachinePoolRuntimeViewResourceProperties struct {
	InstanceCountSummary []VirtualMachineInstanceCountSummary `json:"instanceCountSummary"`
	Prediction           *StandbyVirtualMachinePoolPrediction `json:"prediction,omitempty"`
	ProvisioningState    *ProvisioningState                   `json:"provisioningState,omitempty"`
	Status               *PoolStatus                          `json:"status,omitempty"`
}
