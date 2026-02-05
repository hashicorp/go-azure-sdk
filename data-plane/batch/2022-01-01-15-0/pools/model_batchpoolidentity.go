package pools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchPoolIdentity struct {
	Type                   PoolIdentityType        `json:"type"`
	UserAssignedIdentities *[]UserAssignedIdentity `json:"userAssignedIdentities,omitempty"`
}
