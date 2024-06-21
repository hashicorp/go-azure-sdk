package cluster

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterReportedProperties struct {
	ClusterId             *string          `json:"clusterId,omitempty"`
	ClusterName           *string          `json:"clusterName,omitempty"`
	ClusterType           *ClusterNodeType `json:"clusterType,omitempty"`
	ClusterVersion        *string          `json:"clusterVersion,omitempty"`
	DiagnosticLevel       *DiagnosticLevel `json:"diagnosticLevel,omitempty"`
	ImdsAttestation       *ImdsAttestation `json:"imdsAttestation,omitempty"`
	LastUpdated           *string          `json:"lastUpdated,omitempty"`
	Manufacturer          *string          `json:"manufacturer,omitempty"`
	Nodes                 *[]ClusterNode   `json:"nodes,omitempty"`
	OemActivation         *OemActivation   `json:"oemActivation,omitempty"`
	SupportedCapabilities *[]string        `json:"supportedCapabilities,omitempty"`
}
