package managednamespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceProperties struct {
	AdoptionPolicy       *AdoptionPolicy             `json:"adoptionPolicy,omitempty"`
	Annotations          *map[string]string          `json:"annotations,omitempty"`
	DefaultNetworkPolicy *NetworkPolicies            `json:"defaultNetworkPolicy,omitempty"`
	DefaultResourceQuota *ResourceQuota              `json:"defaultResourceQuota,omitempty"`
	DeletePolicy         *DeletePolicy               `json:"deletePolicy,omitempty"`
	Labels               *map[string]string          `json:"labels,omitempty"`
	PortalFqdn           *string                     `json:"portalFqdn,omitempty"`
	ProvisioningState    *NamespaceProvisioningState `json:"provisioningState,omitempty"`
}
