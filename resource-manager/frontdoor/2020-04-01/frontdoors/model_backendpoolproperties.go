package frontdoors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackendPoolProperties struct {
	Backends              *[]Backend              `json:"backends,omitempty"`
	HealthProbeSettings   *SubResource            `json:"healthProbeSettings"`
	LoadBalancingSettings *SubResource            `json:"loadBalancingSettings"`
	ResourceState         *FrontDoorResourceState `json:"resourceState,omitempty"`
}
