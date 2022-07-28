package privatelinkresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkResourceProperties struct {
	GroupId           *string            `json:"groupId,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	RequiredMembers   *[]string          `json:"requiredMembers,omitempty"`
	RequiredZoneNames *[]string          `json:"requiredZoneNames,omitempty"`
}
