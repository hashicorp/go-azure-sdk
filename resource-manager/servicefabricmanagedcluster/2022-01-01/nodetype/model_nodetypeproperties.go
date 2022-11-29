package nodetype

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeTypeProperties struct {
	AdditionalDataDisks          *[]VmssDataDisk                   `json:"additionalDataDisks,omitempty"`
	ApplicationPorts             *EndpointRangeDescription         `json:"applicationPorts,omitempty"`
	Capacities                   *map[string]string                `json:"capacities,omitempty"`
	DataDiskLetter               *string                           `json:"dataDiskLetter,omitempty"`
	DataDiskSizeGB               *int64                            `json:"dataDiskSizeGB,omitempty"`
	DataDiskType                 *DiskType                         `json:"dataDiskType,omitempty"`
	EnableAcceleratedNetworking  *bool                             `json:"enableAcceleratedNetworking,omitempty"`
	EnableEncryptionAtHost       *bool                             `json:"enableEncryptionAtHost,omitempty"`
	EnableOverProvisioning       *bool                             `json:"enableOverProvisioning,omitempty"`
	EphemeralPorts               *EndpointRangeDescription         `json:"ephemeralPorts,omitempty"`
	FrontendConfigurations       *[]FrontendConfiguration          `json:"frontendConfigurations,omitempty"`
	IsPrimary                    bool                              `json:"isPrimary"`
	IsStateless                  *bool                             `json:"isStateless,omitempty"`
	MultiplePlacementGroups      *bool                             `json:"multiplePlacementGroups,omitempty"`
	NetworkSecurityRules         *[]NetworkSecurityRule            `json:"networkSecurityRules,omitempty"`
	PlacementProperties          *map[string]string                `json:"placementProperties,omitempty"`
	ProvisioningState            *ManagedResourceProvisioningState `json:"provisioningState,omitempty"`
	UseDefaultPublicLoadBalancer *bool                             `json:"useDefaultPublicLoadBalancer,omitempty"`
	UseTempDataDisk              *bool                             `json:"useTempDataDisk,omitempty"`
	VmExtensions                 *[]VMSSExtension                  `json:"vmExtensions,omitempty"`
	VmImageOffer                 *string                           `json:"vmImageOffer,omitempty"`
	VmImagePublisher             *string                           `json:"vmImagePublisher,omitempty"`
	VmImageSku                   *string                           `json:"vmImageSku,omitempty"`
	VmImageVersion               *string                           `json:"vmImageVersion,omitempty"`
	VmInstanceCount              int64                             `json:"vmInstanceCount"`
	VmManagedIdentity            *identity.UserAssignedList        `json:"vmManagedIdentity,omitempty"`
	VmSecrets                    *[]VaultSecretGroup               `json:"vmSecrets,omitempty"`
	VmSize                       *string                           `json:"vmSize,omitempty"`
}
