package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlacementProfile struct {
	ClusterId      *string `json:"clusterId,omitempty"`
	DatastoreId    *string `json:"datastoreId,omitempty"`
	HostId         *string `json:"hostId,omitempty"`
	ResourcePoolId *string `json:"resourcePoolId,omitempty"`
}
