package connectedclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedClusterProperties struct {
	AgentPublicKeyCertificate                string              `json:"agentPublicKeyCertificate"`
	AgentVersion                             *string             `json:"agentVersion,omitempty"`
	ConnectivityStatus                       *ConnectivityStatus `json:"connectivityStatus,omitempty"`
	Distribution                             *string             `json:"distribution,omitempty"`
	Infrastructure                           *string             `json:"infrastructure,omitempty"`
	KubernetesVersion                        *string             `json:"kubernetesVersion,omitempty"`
	LastConnectivityTime                     *string             `json:"lastConnectivityTime,omitempty"`
	ManagedIdentityCertificateExpirationTime *string             `json:"managedIdentityCertificateExpirationTime,omitempty"`
	Offering                                 *string             `json:"offering,omitempty"`
	ProvisioningState                        *ProvisioningState  `json:"provisioningState,omitempty"`
	TotalCoreCount                           *int64              `json:"totalCoreCount,omitempty"`
	TotalNodeCount                           *int64              `json:"totalNodeCount,omitempty"`
}
