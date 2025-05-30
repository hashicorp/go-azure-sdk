package managedclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedClusterIngressProfileWebAppRouting struct {
	DnsZoneResourceIds *[]string                          `json:"dnsZoneResourceIds,omitempty"`
	Enabled            *bool                              `json:"enabled,omitempty"`
	Identity           *UserAssignedIdentity              `json:"identity,omitempty"`
	Nginx              *ManagedClusterIngressProfileNginx `json:"nginx,omitempty"`
}
