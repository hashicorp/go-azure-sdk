package syncgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSyncParticipantIdentity struct {
	TenantId               *string                                             `json:"tenantId,omitempty"`
	Type                   DataSyncParticipantIdentityType                     `json:"type"`
	UserAssignedIdentities *map[string]DataSyncParticipantUserAssignedIdentity `json:"userAssignedIdentities,omitempty"`
}
