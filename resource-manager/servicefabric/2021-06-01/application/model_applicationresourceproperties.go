package application

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationResourceProperties struct {
	ManagedIdentities         *[]ApplicationUserAssignedIdentity `json:"managedIdentities,omitempty"`
	MaximumNodes              *int64                             `json:"maximumNodes,omitempty"`
	Metrics                   *[]ApplicationMetricDescription    `json:"metrics,omitempty"`
	MinimumNodes              *int64                             `json:"minimumNodes,omitempty"`
	Parameters                *map[string]string                 `json:"parameters,omitempty"`
	ProvisioningState         *string                            `json:"provisioningState,omitempty"`
	RemoveApplicationCapacity *bool                              `json:"removeApplicationCapacity,omitempty"`
	TypeName                  *string                            `json:"typeName,omitempty"`
	TypeVersion               *string                            `json:"typeVersion,omitempty"`
	UpgradePolicy             *ApplicationUpgradePolicy          `json:"upgradePolicy,omitempty"`
}
