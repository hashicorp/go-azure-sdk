package standbycontainergrouppoolruntimeviews

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandbyContainerGroupPoolRuntimeViewResourceProperties struct {
	InstanceCountSummary []ContainerGroupInstanceCountSummary `json:"instanceCountSummary"`
	Prediction           *StandbyContainerGroupPoolPrediction `json:"prediction,omitempty"`
	ProvisioningState    *ProvisioningState                   `json:"provisioningState,omitempty"`
	Status               *PoolStatus                          `json:"status,omitempty"`
}
