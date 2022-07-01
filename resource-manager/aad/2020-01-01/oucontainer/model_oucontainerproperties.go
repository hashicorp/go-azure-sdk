package oucontainer

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OuContainerProperties struct {
	Accounts          *[]ContainerAccount `json:"accounts,omitempty"`
	ContainerId       *string             `json:"containerId,omitempty"`
	DeploymentId      *string             `json:"deploymentId,omitempty"`
	DistinguishedName *string             `json:"distinguishedName,omitempty"`
	DomainName        *string             `json:"domainName,omitempty"`
	ProvisioningState *string             `json:"provisioningState,omitempty"`
	ServiceStatus     *string             `json:"serviceStatus,omitempty"`
	TenantId          *string             `json:"tenantId,omitempty"`
}
