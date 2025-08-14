package netappresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegionInfoAvailabilityZoneMappingsInlined struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	IsAvailable      *bool   `json:"isAvailable,omitempty"`
}
