package cluster

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddOnFeatures string

const (
	AddOnFeaturesBackupRestoreService   AddOnFeatures = "BackupRestoreService"
	AddOnFeaturesDnsService             AddOnFeatures = "DnsService"
	AddOnFeaturesRepairManager          AddOnFeatures = "RepairManager"
	AddOnFeaturesResourceMonitorService AddOnFeatures = "ResourceMonitorService"
)

func PossibleValuesForAddOnFeatures() []string {
	return []string{
		string(AddOnFeaturesBackupRestoreService),
		string(AddOnFeaturesDnsService),
		string(AddOnFeaturesRepairManager),
		string(AddOnFeaturesResourceMonitorService),
	}
}

type ClusterState string

const (
	ClusterStateAutoScale                 ClusterState = "AutoScale"
	ClusterStateBaselineUpgrade           ClusterState = "BaselineUpgrade"
	ClusterStateDeploying                 ClusterState = "Deploying"
	ClusterStateEnforcingClusterVersion   ClusterState = "EnforcingClusterVersion"
	ClusterStateReady                     ClusterState = "Ready"
	ClusterStateUpdatingInfrastructure    ClusterState = "UpdatingInfrastructure"
	ClusterStateUpdatingUserCertificate   ClusterState = "UpdatingUserCertificate"
	ClusterStateUpdatingUserConfiguration ClusterState = "UpdatingUserConfiguration"
	ClusterStateUpgradeServiceUnreachable ClusterState = "UpgradeServiceUnreachable"
	ClusterStateWaitingForNodes           ClusterState = "WaitingForNodes"
)

func PossibleValuesForClusterState() []string {
	return []string{
		string(ClusterStateAutoScale),
		string(ClusterStateBaselineUpgrade),
		string(ClusterStateDeploying),
		string(ClusterStateEnforcingClusterVersion),
		string(ClusterStateReady),
		string(ClusterStateUpdatingInfrastructure),
		string(ClusterStateUpdatingUserCertificate),
		string(ClusterStateUpdatingUserConfiguration),
		string(ClusterStateUpgradeServiceUnreachable),
		string(ClusterStateWaitingForNodes),
	}
}

type ClusterUpgradeCadence string

const (
	ClusterUpgradeCadenceWaveOne  ClusterUpgradeCadence = "Wave1"
	ClusterUpgradeCadenceWaveTwo  ClusterUpgradeCadence = "Wave2"
	ClusterUpgradeCadenceWaveZero ClusterUpgradeCadence = "Wave0"
)

func PossibleValuesForClusterUpgradeCadence() []string {
	return []string{
		string(ClusterUpgradeCadenceWaveOne),
		string(ClusterUpgradeCadenceWaveTwo),
		string(ClusterUpgradeCadenceWaveZero),
	}
}

type DurabilityLevel string

const (
	DurabilityLevelBronze DurabilityLevel = "Bronze"
	DurabilityLevelGold   DurabilityLevel = "Gold"
	DurabilityLevelSilver DurabilityLevel = "Silver"
)

func PossibleValuesForDurabilityLevel() []string {
	return []string{
		string(DurabilityLevelBronze),
		string(DurabilityLevelGold),
		string(DurabilityLevelSilver),
	}
}

type Environment string

const (
	EnvironmentLinux   Environment = "Linux"
	EnvironmentWindows Environment = "Windows"
)

func PossibleValuesForEnvironment() []string {
	return []string{
		string(EnvironmentLinux),
		string(EnvironmentWindows),
	}
}

type NotificationCategory string

const (
	NotificationCategoryWaveProgress NotificationCategory = "WaveProgress"
)

func PossibleValuesForNotificationCategory() []string {
	return []string{
		string(NotificationCategoryWaveProgress),
	}
}

type NotificationChannel string

const (
	NotificationChannelEmailSubscription NotificationChannel = "EmailSubscription"
	NotificationChannelEmailUser         NotificationChannel = "EmailUser"
)

func PossibleValuesForNotificationChannel() []string {
	return []string{
		string(NotificationChannelEmailSubscription),
		string(NotificationChannelEmailUser),
	}
}

type NotificationLevel string

const (
	NotificationLevelAll      NotificationLevel = "All"
	NotificationLevelCritical NotificationLevel = "Critical"
)

func PossibleValuesForNotificationLevel() []string {
	return []string{
		string(NotificationLevelAll),
		string(NotificationLevelCritical),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type ReliabilityLevel string

const (
	ReliabilityLevelBronze   ReliabilityLevel = "Bronze"
	ReliabilityLevelGold     ReliabilityLevel = "Gold"
	ReliabilityLevelNone     ReliabilityLevel = "None"
	ReliabilityLevelPlatinum ReliabilityLevel = "Platinum"
	ReliabilityLevelSilver   ReliabilityLevel = "Silver"
)

func PossibleValuesForReliabilityLevel() []string {
	return []string{
		string(ReliabilityLevelBronze),
		string(ReliabilityLevelGold),
		string(ReliabilityLevelNone),
		string(ReliabilityLevelPlatinum),
		string(ReliabilityLevelSilver),
	}
}

type SfZonalUpgradeMode string

const (
	SfZonalUpgradeModeHierarchical SfZonalUpgradeMode = "Hierarchical"
	SfZonalUpgradeModeParallel     SfZonalUpgradeMode = "Parallel"
)

func PossibleValuesForSfZonalUpgradeMode() []string {
	return []string{
		string(SfZonalUpgradeModeHierarchical),
		string(SfZonalUpgradeModeParallel),
	}
}

type UpgradeMode string

const (
	UpgradeModeAutomatic UpgradeMode = "Automatic"
	UpgradeModeManual    UpgradeMode = "Manual"
)

func PossibleValuesForUpgradeMode() []string {
	return []string{
		string(UpgradeModeAutomatic),
		string(UpgradeModeManual),
	}
}

type VMSSZonalUpgradeMode string

const (
	VMSSZonalUpgradeModeHierarchical VMSSZonalUpgradeMode = "Hierarchical"
	VMSSZonalUpgradeModeParallel     VMSSZonalUpgradeMode = "Parallel"
)

func PossibleValuesForVMSSZonalUpgradeMode() []string {
	return []string{
		string(VMSSZonalUpgradeModeHierarchical),
		string(VMSSZonalUpgradeModeParallel),
	}
}

type X509StoreName string

const (
	X509StoreNameAddressBook          X509StoreName = "AddressBook"
	X509StoreNameAuthRoot             X509StoreName = "AuthRoot"
	X509StoreNameCertificateAuthority X509StoreName = "CertificateAuthority"
	X509StoreNameDisallowed           X509StoreName = "Disallowed"
	X509StoreNameMy                   X509StoreName = "My"
	X509StoreNameRoot                 X509StoreName = "Root"
	X509StoreNameTrustedPeople        X509StoreName = "TrustedPeople"
	X509StoreNameTrustedPublisher     X509StoreName = "TrustedPublisher"
)

func PossibleValuesForX509StoreName() []string {
	return []string{
		string(X509StoreNameAddressBook),
		string(X509StoreNameAuthRoot),
		string(X509StoreNameCertificateAuthority),
		string(X509StoreNameDisallowed),
		string(X509StoreNameMy),
		string(X509StoreNameRoot),
		string(X509StoreNameTrustedPeople),
		string(X509StoreNameTrustedPublisher),
	}
}
