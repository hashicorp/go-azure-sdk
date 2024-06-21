package scopemaps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopeMapProperties struct {
	Actions           []string           `json:"actions"`
	CreationDate      *string            `json:"creationDate,omitempty"`
	Description       *string            `json:"description,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	Type              *string            `json:"type,omitempty"`
}
