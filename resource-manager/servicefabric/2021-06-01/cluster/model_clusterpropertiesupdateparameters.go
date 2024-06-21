package cluster

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterPropertiesUpdateParameters struct {
	AddOnFeatures                        *[]AddOnFeatures                      `json:"addOnFeatures,omitempty"`
	ApplicationTypeVersionsCleanupPolicy *ApplicationTypeVersionsCleanupPolicy `json:"applicationTypeVersionsCleanupPolicy,omitempty"`
	Certificate                          *CertificateDescription               `json:"certificate,omitempty"`
	CertificateCommonNames               *ServerCertificateCommonNames         `json:"certificateCommonNames,omitempty"`
	ClientCertificateCommonNames         *[]ClientCertificateCommonName        `json:"clientCertificateCommonNames,omitempty"`
	ClientCertificateThumbprints         *[]ClientCertificateThumbprint        `json:"clientCertificateThumbprints,omitempty"`
	ClusterCodeVersion                   *string                               `json:"clusterCodeVersion,omitempty"`
	EventStoreServiceEnabled             *bool                                 `json:"eventStoreServiceEnabled,omitempty"`
	FabricSettings                       *[]SettingsSectionDescription         `json:"fabricSettings,omitempty"`
	InfrastructureServiceManager         *bool                                 `json:"infrastructureServiceManager,omitempty"`
	NodeTypes                            *[]NodeTypeDescription                `json:"nodeTypes,omitempty"`
	Notifications                        *[]Notification                       `json:"notifications,omitempty"`
	ReliabilityLevel                     *ReliabilityLevel                     `json:"reliabilityLevel,omitempty"`
	ReverseProxyCertificate              *CertificateDescription               `json:"reverseProxyCertificate,omitempty"`
	SfZonalUpgradeMode                   *SfZonalUpgradeMode                   `json:"sfZonalUpgradeMode,omitempty"`
	UpgradeDescription                   *ClusterUpgradePolicy                 `json:"upgradeDescription,omitempty"`
	UpgradeMode                          *UpgradeMode                          `json:"upgradeMode,omitempty"`
	UpgradePauseEndTimestampUtc          *string                               `json:"upgradePauseEndTimestampUtc,omitempty"`
	UpgradePauseStartTimestampUtc        *string                               `json:"upgradePauseStartTimestampUtc,omitempty"`
	UpgradeWave                          *ClusterUpgradeCadence                `json:"upgradeWave,omitempty"`
	VMSSZonalUpgradeMode                 *VMSSZonalUpgradeMode                 `json:"vmssZonalUpgradeMode,omitempty"`
	WaveUpgradePaused                    *bool                                 `json:"waveUpgradePaused,omitempty"`
}
