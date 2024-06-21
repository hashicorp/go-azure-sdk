package registries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryProperties struct {
	AdminUserEnabled           *bool                        `json:"adminUserEnabled,omitempty"`
	AnonymousPullEnabled       *bool                        `json:"anonymousPullEnabled,omitempty"`
	CreationDate               *string                      `json:"creationDate,omitempty"`
	DataEndpointEnabled        *bool                        `json:"dataEndpointEnabled,omitempty"`
	DataEndpointHostNames      *[]string                    `json:"dataEndpointHostNames,omitempty"`
	Encryption                 *EncryptionProperty          `json:"encryption,omitempty"`
	LoginServer                *string                      `json:"loginServer,omitempty"`
	NetworkRuleBypassOptions   *NetworkRuleBypassOptions    `json:"networkRuleBypassOptions,omitempty"`
	NetworkRuleSet             *NetworkRuleSet              `json:"networkRuleSet,omitempty"`
	Policies                   *Policies                    `json:"policies,omitempty"`
	PrivateEndpointConnections *[]PrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *ProvisioningState           `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *PublicNetworkAccess         `json:"publicNetworkAccess,omitempty"`
	Status                     *Status                      `json:"status,omitempty"`
	ZoneRedundancy             *ZoneRedundancy              `json:"zoneRedundancy,omitempty"`
}
