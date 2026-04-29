package caches

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CacheProperties struct {
	ActualThroughputMibps             *float64                      `json:"actualThroughputMibps,omitempty"`
	CacheState                        *CacheLifeCycleState          `json:"cacheState,omitempty"`
	CacheSubnetResourceId             string                        `json:"cacheSubnetResourceId"`
	CifsChangeNotifications           *CifsChangeNotifyState        `json:"cifsChangeNotifications,omitempty"`
	Encryption                        *EncryptionState              `json:"encryption,omitempty"`
	EncryptionKeySource               EncryptionKeySource           `json:"encryptionKeySource"`
	ExportPolicy                      *CachePropertiesExportPolicy  `json:"exportPolicy,omitempty"`
	FilePath                          string                        `json:"filePath"`
	GlobalFileLocking                 *GlobalFileLockingState       `json:"globalFileLocking,omitempty"`
	Kerberos                          *KerberosState                `json:"kerberos,omitempty"`
	KeyVaultPrivateEndpointResourceId *string                       `json:"keyVaultPrivateEndpointResourceId,omitempty"`
	Language                          *VolumeLanguage               `json:"language,omitempty"`
	Ldap                              *LdapState                    `json:"ldap,omitempty"`
	LdapServerType                    *LdapServerType               `json:"ldapServerType,omitempty"`
	MaximumNumberOfFiles              *int64                        `json:"maximumNumberOfFiles,omitempty"`
	MountTargets                      *[]CacheMountTargetProperties `json:"mountTargets,omitempty"`
	OriginClusterInformation          OriginClusterInformation      `json:"originClusterInformation"`
	PeeringSubnetResourceId           string                        `json:"peeringSubnetResourceId"`
	ProtocolTypes                     *[]ProtocolTypes              `json:"protocolTypes,omitempty"`
	ProvisioningState                 *CacheProvisioningState       `json:"provisioningState,omitempty"`
	Size                              int64                         `json:"size"`
	SmbSettings                       *SmbSettings                  `json:"smbSettings,omitempty"`
	ThroughputMibps                   *float64                      `json:"throughputMibps,omitempty"`
	WriteBack                         *EnableWriteBackState         `json:"writeBack,omitempty"`
}
