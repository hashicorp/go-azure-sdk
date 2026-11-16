package nodes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeProperties struct {
	AvailabilityZone string   `json:"availabilityZone"`
	Role             NodeRole `json:"role"`
}
