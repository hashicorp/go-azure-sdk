package fleetmanagednamespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedNamespaceProperties struct {
	Annotations          *map[string]string `json:"annotations,omitempty"`
	DefaultNetworkPolicy *NetworkPolicy     `json:"defaultNetworkPolicy,omitempty"`
	DefaultResourceQuota *ResourceQuota     `json:"defaultResourceQuota,omitempty"`
	Labels               *map[string]string `json:"labels,omitempty"`
}
