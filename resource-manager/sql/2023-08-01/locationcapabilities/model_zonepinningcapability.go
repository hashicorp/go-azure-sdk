package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZonePinningCapability struct {
	AvailabilityZone *string           `json:"availabilityZone,omitempty"`
	Reason           *string           `json:"reason,omitempty"`
	Status           *CapabilityStatus `json:"status,omitempty"`
}
