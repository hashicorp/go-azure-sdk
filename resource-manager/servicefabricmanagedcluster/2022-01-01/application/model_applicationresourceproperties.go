package application

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationResourceProperties struct {
	ManagedIdentities *[]ApplicationUserAssignedIdentity `json:"managedIdentities,omitempty"`
	Parameters        *map[string]string                 `json:"parameters,omitempty"`
	ProvisioningState *string                            `json:"provisioningState,omitempty"`
	UpgradePolicy     *ApplicationUpgradePolicy          `json:"upgradePolicy"`
	Version           *string                            `json:"version,omitempty"`
}
