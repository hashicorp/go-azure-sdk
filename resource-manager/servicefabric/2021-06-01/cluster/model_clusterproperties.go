package cluster

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterProperties struct {
	AddOnFeatures                        *[]AddOnFeatures                      `json:"addOnFeatures,omitempty"`
	ApplicationTypeVersionsCleanupPolicy *ApplicationTypeVersionsCleanupPolicy `json:"applicationTypeVersionsCleanupPolicy,omitempty"`
	AvailableClusterVersions             *[]ClusterVersionDetails              `json:"availableClusterVersions,omitempty"`
	AzureActiveDirectory                 *AzureActiveDirectory                 `json:"azureActiveDirectory,omitempty"`
	Certificate                          *CertificateDescription               `json:"certificate,omitempty"`
	CertificateCommonNames               *ServerCertificateCommonNames         `json:"certificateCommonNames,omitempty"`
	ClientCertificateCommonNames         *[]ClientCertificateCommonName        `json:"clientCertificateCommonNames,omitempty"`
	ClientCertificateThumbprints         *[]ClientCertificateThumbprint        `json:"clientCertificateThumbprints,omitempty"`
	ClusterCodeVersion                   *string                               `json:"clusterCodeVersion,omitempty"`
	ClusterEndpoint                      *string                               `json:"clusterEndpoint,omitempty"`
	ClusterId                            *string                               `json:"clusterId,omitempty"`
	ClusterState                         *ClusterState                         `json:"clusterState,omitempty"`
	DiagnosticsStorageAccountConfig      *DiagnosticsStorageAccountConfig      `json:"diagnosticsStorageAccountConfig,omitempty"`
	EventStoreServiceEnabled             *bool                                 `json:"eventStoreServiceEnabled,omitempty"`
	FabricSettings                       *[]SettingsSectionDescription         `json:"fabricSettings,omitempty"`
	InfrastructureServiceManager         *bool                                 `json:"infrastructureServiceManager,omitempty"`
	ManagementEndpoint                   string                                `json:"managementEndpoint"`
	NodeTypes                            []NodeTypeDescription                 `json:"nodeTypes"`
	Notifications                        *[]Notification                       `json:"notifications,omitempty"`
	ProvisioningState                    *ProvisioningState                    `json:"provisioningState,omitempty"`
	ReliabilityLevel                     *ReliabilityLevel                     `json:"reliabilityLevel,omitempty"`
	ReverseProxyCertificate              *CertificateDescription               `json:"reverseProxyCertificate,omitempty"`
	ReverseProxyCertificateCommonNames   *ServerCertificateCommonNames         `json:"reverseProxyCertificateCommonNames,omitempty"`
	SfZonalUpgradeMode                   *SfZonalUpgradeMode                   `json:"sfZonalUpgradeMode,omitempty"`
	UpgradeDescription                   *ClusterUpgradePolicy                 `json:"upgradeDescription,omitempty"`
	UpgradeMode                          *UpgradeMode                          `json:"upgradeMode,omitempty"`
	UpgradePauseEndTimestampUtc          *string                               `json:"upgradePauseEndTimestampUtc,omitempty"`
	UpgradePauseStartTimestampUtc        *string                               `json:"upgradePauseStartTimestampUtc,omitempty"`
	UpgradeWave                          *ClusterUpgradeCadence                `json:"upgradeWave,omitempty"`
	VMSSZonalUpgradeMode                 *VMSSZonalUpgradeMode                 `json:"vmssZonalUpgradeMode,omitempty"`
	VmImage                              *string                               `json:"vmImage,omitempty"`
	WaveUpgradePaused                    *bool                                 `json:"waveUpgradePaused,omitempty"`
}
