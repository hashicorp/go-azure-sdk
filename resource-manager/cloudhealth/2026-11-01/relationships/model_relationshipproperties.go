package relationships

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelationshipProperties struct {
	ChildEntityName   string                        `json:"childEntityName"`
	DiscoveredBy      *string                       `json:"discoveredBy,omitempty"`
	DisplayName       *string                       `json:"displayName,omitempty"`
	ParentEntityName  string                        `json:"parentEntityName"`
	ProvisioningState *HealthModelProvisioningState `json:"provisioningState,omitempty"`
	Tags              *map[string]string            `json:"tags,omitempty"`
}
