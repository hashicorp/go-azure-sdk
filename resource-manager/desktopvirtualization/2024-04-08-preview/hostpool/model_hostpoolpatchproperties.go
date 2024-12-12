package hostpool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostPoolPatchProperties struct {
	AgentUpdate                   *AgentUpdatePatchProperties    `json:"agentUpdate,omitempty"`
	CustomRdpProperty             *string                        `json:"customRdpProperty,omitempty"`
	Description                   *string                        `json:"description,omitempty"`
	DirectUDP                     *DirectUDP                     `json:"directUDP,omitempty"`
	FriendlyName                  *string                        `json:"friendlyName,omitempty"`
	LoadBalancerType              *LoadBalancerType              `json:"loadBalancerType,omitempty"`
	ManagedPrivateUDP             *ManagedPrivateUDP             `json:"managedPrivateUDP,omitempty"`
	MaxSessionLimit               *int64                         `json:"maxSessionLimit,omitempty"`
	PersonalDesktopAssignmentType *PersonalDesktopAssignmentType `json:"personalDesktopAssignmentType,omitempty"`
	PreferredAppGroupType         *PreferredAppGroupType         `json:"preferredAppGroupType,omitempty"`
	PublicNetworkAccess           *HostpoolPublicNetworkAccess   `json:"publicNetworkAccess,omitempty"`
	PublicUDP                     *PublicUDP                     `json:"publicUDP,omitempty"`
	RegistrationInfo              *RegistrationInfoPatch         `json:"registrationInfo,omitempty"`
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
