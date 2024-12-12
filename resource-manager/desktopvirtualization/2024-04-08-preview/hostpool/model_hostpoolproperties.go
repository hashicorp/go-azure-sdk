package hostpool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostPoolProperties struct {
	AgentUpdate                   *AgentUpdateProperties         `json:"agentUpdate,omitempty"`
	AppAttachPackageReferences    *[]string                      `json:"appAttachPackageReferences,omitempty"`
	ApplicationGroupReferences    *[]string                      `json:"applicationGroupReferences,omitempty"`
	CloudPcResource               *bool                          `json:"cloudPcResource,omitempty"`
	CustomRdpProperty             *string                        `json:"customRdpProperty,omitempty"`
	Description                   *string                        `json:"description,omitempty"`
	DirectUDP                     *DirectUDP                     `json:"directUDP,omitempty"`
	FriendlyName                  *string                        `json:"friendlyName,omitempty"`
	HostPoolType                  HostPoolType                   `json:"hostPoolType"`
	LoadBalancerType              LoadBalancerType               `json:"loadBalancerType"`
	ManagedPrivateUDP             *ManagedPrivateUDP             `json:"managedPrivateUDP,omitempty"`
	ManagementType                *ManagementType                `json:"managementType,omitempty"`
	MaxSessionLimit               *int64                         `json:"maxSessionLimit,omitempty"`
	ObjectId                      *string                        `json:"objectId,omitempty"`
	PersonalDesktopAssignmentType *PersonalDesktopAssignmentType `json:"personalDesktopAssignmentType,omitempty"`
	PreferredAppGroupType         PreferredAppGroupType          `json:"preferredAppGroupType"`
	PrivateEndpointConnections    *[]PrivateEndpointConnection   `json:"privateEndpointConnections,omitempty"`
	PublicNetworkAccess           *HostpoolPublicNetworkAccess   `json:"publicNetworkAccess,omitempty"`
	PublicUDP                     *PublicUDP                     `json:"publicUDP,omitempty"`
	RegistrationInfo              *RegistrationInfo              `json:"registrationInfo,omitempty"`
	RelayUDP                      *RelayUDP                      `json:"relayUDP,omitempty"`
	Ring                          *int64                         `json:"ring,omitempty"`
	SsoClientId                   *string                        `json:"ssoClientId,omitempty"`
	SsoClientSecretKeyVaultPath   *string                        `json:"ssoClientSecretKeyVaultPath,omitempty"`
	SsoSecretType                 *SSOSecretType                 `json:"ssoSecretType,omitempty"`
	SsoadfsAuthority              *string                        `json:"ssoadfsAuthority,omitempty"`
	StartVMOnConnect              *bool                          `json:"startVMOnConnect,omitempty"`
	VMTemplate                    *string                        `json:"vmTemplate,omitempty"`
	ValidationEnvironment         *bool                          `json:"validationEnvironment,omitempty"`
}
