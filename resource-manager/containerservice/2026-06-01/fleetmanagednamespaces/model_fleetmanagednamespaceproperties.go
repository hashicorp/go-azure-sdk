package fleetmanagednamespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FleetManagedNamespaceProperties struct {
	AdoptionPolicy             AdoptionPolicy                          `json:"adoptionPolicy"`
	DeletePolicy               DeletePolicy                            `json:"deletePolicy"`
	ManagedNamespaceProperties *ManagedNamespaceProperties             `json:"managedNamespaceProperties,omitempty"`
	PortalFqdn                 *string                                 `json:"portalFqdn,omitempty"`
	PropagationPolicy          *PropagationPolicy                      `json:"propagationPolicy,omitempty"`
	ProvisioningState          *FleetManagedNamespaceProvisioningState `json:"provisioningState,omitempty"`
	Status                     *FleetManagedNamespaceStatus            `json:"status,omitempty"`
}
