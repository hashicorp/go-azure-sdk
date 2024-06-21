package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointProperties struct {
	ConsistencyMode   *ConsistencyModeTypes       `json:"consistencyMode,omitempty"`
	ExcludeDisks      *[]ApiEntityReference       `json:"excludeDisks,omitempty"`
	ProvisioningState *string                     `json:"provisioningState,omitempty"`
	SourceMetadata    *RestorePointSourceMetadata `json:"sourceMetadata,omitempty"`
	TimeCreated       *string                     `json:"timeCreated,omitempty"`
}
