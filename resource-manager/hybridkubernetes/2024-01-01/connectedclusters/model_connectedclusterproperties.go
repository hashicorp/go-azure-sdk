package connectedclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedClusterProperties struct {
	AadProfile                               *AadProfile         `json:"aadProfile,omitempty"`
	AgentPublicKeyCertificate                string              `json:"agentPublicKeyCertificate"`
	AgentVersion                             *string             `json:"agentVersion,omitempty"`
	ArcAgentProfile                          *ArcAgentProfile    `json:"arcAgentProfile,omitempty"`
	AzureHybridBenefit                       *AzureHybridBenefit `json:"azureHybridBenefit,omitempty"`
	ConnectivityStatus                       *ConnectivityStatus `json:"connectivityStatus,omitempty"`
	Distribution                             *string             `json:"distribution,omitempty"`
	DistributionVersion                      *string             `json:"distributionVersion,omitempty"`
	Infrastructure                           *string             `json:"infrastructure,omitempty"`
	KubernetesVersion                        *string             `json:"kubernetesVersion,omitempty"`
	LastConnectivityTime                     *string             `json:"lastConnectivityTime,omitempty"`
	ManagedIdentityCertificateExpirationTime *string             `json:"managedIdentityCertificateExpirationTime,omitempty"`
	MiscellaneousProperties                  *map[string]string  `json:"miscellaneousProperties,omitempty"`
	Offering                                 *string             `json:"offering,omitempty"`
	PrivateLinkScopeResourceId               *string             `json:"privateLinkScopeResourceId,omitempty"`
	PrivateLinkState                         *PrivateLinkState   `json:"privateLinkState,omitempty"`
	ProvisioningState                        *ProvisioningState  `json:"provisioningState,omitempty"`
	TotalCoreCount                           *int64              `json:"totalCoreCount,omitempty"`
	TotalNodeCount                           *int64              `json:"totalNodeCount,omitempty"`
}
