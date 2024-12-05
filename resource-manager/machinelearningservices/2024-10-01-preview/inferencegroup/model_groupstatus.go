package inferencegroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupStatus struct {
	ActualCapacityInfo *ActualCapacityInfo `json:"actualCapacityInfo,omitempty"`
	EndpointCount      *int64              `json:"endpointCount,omitempty"`
	RequestedCapacity  *int64              `json:"requestedCapacity,omitempty"`
}
