package devices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationType string

const (
	AuthenticationTypeAzureActiveDirectory AuthenticationType = "AzureActiveDirectory"
	AuthenticationTypeInvalid              AuthenticationType = "Invalid"
)

func PossibleValuesForAuthenticationType() []string {
	return []string{
		string(AuthenticationTypeAzureActiveDirectory),
		string(AuthenticationTypeInvalid),
	}
}

type ClusterWitnessType string

const (
	ClusterWitnessTypeCloud     ClusterWitnessType = "Cloud"
	ClusterWitnessTypeFileShare ClusterWitnessType = "FileShare"
	ClusterWitnessTypeNone      ClusterWitnessType = "None"
)

func PossibleValuesForClusterWitnessType() []string {
	return []string{
		string(ClusterWitnessTypeCloud),
		string(ClusterWitnessTypeFileShare),
		string(ClusterWitnessTypeNone),
	}
}

type DataBoxEdgeDeviceKind string

const (
	DataBoxEdgeDeviceKindAzureDataBoxGateway    DataBoxEdgeDeviceKind = "AzureDataBoxGateway"
	DataBoxEdgeDeviceKindAzureModularDataCentre DataBoxEdgeDeviceKind = "AzureModularDataCentre"
	DataBoxEdgeDeviceKindAzureStackEdge         DataBoxEdgeDeviceKind = "AzureStackEdge"
	DataBoxEdgeDeviceKindAzureStackHub          DataBoxEdgeDeviceKind = "AzureStackHub"
)

func PossibleValuesForDataBoxEdgeDeviceKind() []string {
	return []string{
		string(DataBoxEdgeDeviceKindAzureDataBoxGateway),
		string(DataBoxEdgeDeviceKindAzureModularDataCentre),
		string(DataBoxEdgeDeviceKindAzureStackEdge),
		string(DataBoxEdgeDeviceKindAzureStackHub),
	}
}

type DataBoxEdgeDeviceStatus string

const (
	DataBoxEdgeDeviceStatusDisconnected          DataBoxEdgeDeviceStatus = "Disconnected"
	DataBoxEdgeDeviceStatusMaintenance           DataBoxEdgeDeviceStatus = "Maintenance"
	DataBoxEdgeDeviceStatusNeedsAttention        DataBoxEdgeDeviceStatus = "NeedsAttention"
	DataBoxEdgeDeviceStatusOffline               DataBoxEdgeDeviceStatus = "Offline"
	DataBoxEdgeDeviceStatusOnline                DataBoxEdgeDeviceStatus = "Online"
	DataBoxEdgeDeviceStatusPartiallyDisconnected DataBoxEdgeDeviceStatus = "PartiallyDisconnected"
	DataBoxEdgeDeviceStatusReadyToSetup          DataBoxEdgeDeviceStatus = "ReadyToSetup"
)

func PossibleValuesForDataBoxEdgeDeviceStatus() []string {
	return []string{
		string(DataBoxEdgeDeviceStatusDisconnected),
		string(DataBoxEdgeDeviceStatusMaintenance),
		string(DataBoxEdgeDeviceStatusNeedsAttention),
		string(DataBoxEdgeDeviceStatusOffline),
		string(DataBoxEdgeDeviceStatusOnline),
		string(DataBoxEdgeDeviceStatusPartiallyDisconnected),
		string(DataBoxEdgeDeviceStatusReadyToSetup),
	}
}

type DataResidencyType string

const (
	DataResidencyTypeGeoZoneReplication DataResidencyType = "GeoZoneReplication"
	DataResidencyTypeZoneReplication    DataResidencyType = "ZoneReplication"
)

func PossibleValuesForDataResidencyType() []string {
	return []string{
		string(DataResidencyTypeGeoZoneReplication),
		string(DataResidencyTypeZoneReplication),
	}
}

type DeviceType string

const (
	DeviceTypeDataBoxEdgeDevice DeviceType = "DataBoxEdgeDevice"
)

func PossibleValuesForDeviceType() []string {
	return []string{
		string(DeviceTypeDataBoxEdgeDevice),
	}
}

type EncryptionAlgorithm string

const (
	EncryptionAlgorithmAESTwoFiveSix        EncryptionAlgorithm = "AES256"
	EncryptionAlgorithmNone                 EncryptionAlgorithm = "None"
	EncryptionAlgorithmRSAESPKCSOneVOneFive EncryptionAlgorithm = "RSAES_PKCS1_v_1_5"
)

func PossibleValuesForEncryptionAlgorithm() []string {
	return []string{
		string(EncryptionAlgorithmAESTwoFiveSix),
		string(EncryptionAlgorithmNone),
		string(EncryptionAlgorithmRSAESPKCSOneVOneFive),
	}
}

type InstallRebootBehavior string

const (
	InstallRebootBehaviorNeverReboots   InstallRebootBehavior = "NeverReboots"
	InstallRebootBehaviorRequestReboot  InstallRebootBehavior = "RequestReboot"
	InstallRebootBehaviorRequiresReboot InstallRebootBehavior = "RequiresReboot"
)

func PossibleValuesForInstallRebootBehavior() []string {
	return []string{
		string(InstallRebootBehaviorNeverReboots),
		string(InstallRebootBehaviorRequestReboot),
		string(InstallRebootBehaviorRequiresReboot),
	}
}

type InstallationImpact string

const (
	InstallationImpactDeviceRebooted          InstallationImpact = "DeviceRebooted"
	InstallationImpactKubernetesWorkloadsDown InstallationImpact = "KubernetesWorkloadsDown"
	InstallationImpactNone                    InstallationImpact = "None"
)

func PossibleValuesForInstallationImpact() []string {
	return []string{
		string(InstallationImpactDeviceRebooted),
		string(InstallationImpactKubernetesWorkloadsDown),
		string(InstallationImpactNone),
	}
}

type JobStatus string

const (
	JobStatusCanceled  JobStatus = "Canceled"
	JobStatusFailed    JobStatus = "Failed"
	JobStatusInvalid   JobStatus = "Invalid"
	JobStatusPaused    JobStatus = "Paused"
	JobStatusRunning   JobStatus = "Running"
	JobStatusScheduled JobStatus = "Scheduled"
	JobStatusSucceeded JobStatus = "Succeeded"
)

func PossibleValuesForJobStatus() []string {
	return []string{
		string(JobStatusCanceled),
		string(JobStatusFailed),
		string(JobStatusInvalid),
		string(JobStatusPaused),
		string(JobStatusRunning),
		string(JobStatusScheduled),
		string(JobStatusSucceeded),
	}
}

type KeyVaultSyncStatus string

const (
	KeyVaultSyncStatusKeyVaultNotConfigured KeyVaultSyncStatus = "KeyVaultNotConfigured"
	KeyVaultSyncStatusKeyVaultNotSynced     KeyVaultSyncStatus = "KeyVaultNotSynced"
	KeyVaultSyncStatusKeyVaultSyncFailed    KeyVaultSyncStatus = "KeyVaultSyncFailed"
	KeyVaultSyncStatusKeyVaultSyncPending   KeyVaultSyncStatus = "KeyVaultSyncPending"
	KeyVaultSyncStatusKeyVaultSynced        KeyVaultSyncStatus = "KeyVaultSynced"
	KeyVaultSyncStatusKeyVaultSyncing       KeyVaultSyncStatus = "KeyVaultSyncing"
)

func PossibleValuesForKeyVaultSyncStatus() []string {
	return []string{
		string(KeyVaultSyncStatusKeyVaultNotConfigured),
		string(KeyVaultSyncStatusKeyVaultNotSynced),
		string(KeyVaultSyncStatusKeyVaultSyncFailed),
		string(KeyVaultSyncStatusKeyVaultSyncPending),
		string(KeyVaultSyncStatusKeyVaultSynced),
		string(KeyVaultSyncStatusKeyVaultSyncing),
	}
}

type MsiIdentityType string

const (
	MsiIdentityTypeNone           MsiIdentityType = "None"
	MsiIdentityTypeSystemAssigned MsiIdentityType = "SystemAssigned"
	MsiIdentityTypeUserAssigned   MsiIdentityType = "UserAssigned"
)

func PossibleValuesForMsiIdentityType() []string {
	return []string{
		string(MsiIdentityTypeNone),
		string(MsiIdentityTypeSystemAssigned),
		string(MsiIdentityTypeUserAssigned),
	}
}

type NetworkAdapterDHCPStatus string

const (
	NetworkAdapterDHCPStatusDisabled NetworkAdapterDHCPStatus = "Disabled"
	NetworkAdapterDHCPStatusEnabled  NetworkAdapterDHCPStatus = "Enabled"
)

func PossibleValuesForNetworkAdapterDHCPStatus() []string {
	return []string{
		string(NetworkAdapterDHCPStatusDisabled),
		string(NetworkAdapterDHCPStatusEnabled),
	}
}

type NetworkAdapterRDMAStatus string

const (
	NetworkAdapterRDMAStatusCapable   NetworkAdapterRDMAStatus = "Capable"
	NetworkAdapterRDMAStatusIncapable NetworkAdapterRDMAStatus = "Incapable"
)

func PossibleValuesForNetworkAdapterRDMAStatus() []string {
	return []string{
		string(NetworkAdapterRDMAStatusCapable),
		string(NetworkAdapterRDMAStatusIncapable),
	}
}

type NetworkAdapterStatus string

const (
	NetworkAdapterStatusActive   NetworkAdapterStatus = "Active"
	NetworkAdapterStatusInactive NetworkAdapterStatus = "Inactive"
)

func PossibleValuesForNetworkAdapterStatus() []string {
	return []string{
		string(NetworkAdapterStatusActive),
		string(NetworkAdapterStatusInactive),
	}
}

type NetworkGroup string

const (
	NetworkGroupNonRDMA NetworkGroup = "NonRDMA"
	NetworkGroupNone    NetworkGroup = "None"
	NetworkGroupRDMA    NetworkGroup = "RDMA"
)

func PossibleValuesForNetworkGroup() []string {
	return []string{
		string(NetworkGroupNonRDMA),
		string(NetworkGroupNone),
		string(NetworkGroupRDMA),
	}
}

type ResourceMoveStatus string

const (
	ResourceMoveStatusNone                   ResourceMoveStatus = "None"
	ResourceMoveStatusResourceMoveFailed     ResourceMoveStatus = "ResourceMoveFailed"
	ResourceMoveStatusResourceMoveInProgress ResourceMoveStatus = "ResourceMoveInProgress"
)

func PossibleValuesForResourceMoveStatus() []string {
	return []string{
		string(ResourceMoveStatusNone),
		string(ResourceMoveStatusResourceMoveFailed),
		string(ResourceMoveStatusResourceMoveInProgress),
	}
}

type RoleTypes string

const (
	RoleTypesASA                 RoleTypes = "ASA"
	RoleTypesCloudEdgeManagement RoleTypes = "CloudEdgeManagement"
	RoleTypesCognitive           RoleTypes = "Cognitive"
	RoleTypesFunctions           RoleTypes = "Functions"
	RoleTypesIOT                 RoleTypes = "IOT"
	RoleTypesKubernetes          RoleTypes = "Kubernetes"
	RoleTypesMEC                 RoleTypes = "MEC"
)

func PossibleValuesForRoleTypes() []string {
	return []string{
		string(RoleTypesASA),
		string(RoleTypesCloudEdgeManagement),
		string(RoleTypesCognitive),
		string(RoleTypesFunctions),
		string(RoleTypesIOT),
		string(RoleTypesKubernetes),
		string(RoleTypesMEC),
	}
}

type SkuName string

const (
	SkuNameEPTwoOneTwoEightGPUOneMxOneW   SkuName = "EP2_128_GPU1_Mx1_W"
	SkuNameEPTwoOneTwoEightOneTFourMxOneW SkuName = "EP2_128_1T4_Mx1_W"
	SkuNameEPTwoSixFourMxOneW             SkuName = "EP2_64_Mx1_W"
	SkuNameEPTwoSixFourOneVPUW            SkuName = "EP2_64_1VPU_W"
	SkuNameEPTwoTwoFiveSixGPUTwoMxOne     SkuName = "EP2_256_GPU2_Mx1"
	SkuNameEPTwoTwoFiveSixTwoTFourW       SkuName = "EP2_256_2T4_W"
	SkuNameEdge                           SkuName = "Edge"
	SkuNameEdgeMRMini                     SkuName = "EdgeMR_Mini"
	SkuNameEdgeMRTCP                      SkuName = "EdgeMR_TCP"
	SkuNameEdgePBase                      SkuName = "EdgeP_Base"
	SkuNameEdgePHigh                      SkuName = "EdgeP_High"
	SkuNameEdgePRBase                     SkuName = "EdgePR_Base"
	SkuNameEdgePRBaseUPS                  SkuName = "EdgePR_Base_UPS"
	SkuNameGPU                            SkuName = "GPU"
	SkuNameGateway                        SkuName = "Gateway"
	SkuNameManagement                     SkuName = "Management"
	SkuNameRCALarge                       SkuName = "RCA_Large"
	SkuNameRCASmall                       SkuName = "RCA_Small"
	SkuNameRDC                            SkuName = "RDC"
	SkuNameTCALarge                       SkuName = "TCA_Large"
	SkuNameTCASmall                       SkuName = "TCA_Small"
	SkuNameTDC                            SkuName = "TDC"
	SkuNameTEAFourNodeHeater              SkuName = "TEA_4Node_Heater"
	SkuNameTEAFourNodeUPSHeater           SkuName = "TEA_4Node_UPS_Heater"
	SkuNameTEAOneNode                     SkuName = "TEA_1Node"
	SkuNameTEAOneNodeHeater               SkuName = "TEA_1Node_Heater"
	SkuNameTEAOneNodeUPS                  SkuName = "TEA_1Node_UPS"
	SkuNameTEAOneNodeUPSHeater            SkuName = "TEA_1Node_UPS_Heater"
	SkuNameTMA                            SkuName = "TMA"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameEPTwoOneTwoEightGPUOneMxOneW),
		string(SkuNameEPTwoOneTwoEightOneTFourMxOneW),
		string(SkuNameEPTwoSixFourMxOneW),
		string(SkuNameEPTwoSixFourOneVPUW),
		string(SkuNameEPTwoTwoFiveSixGPUTwoMxOne),
		string(SkuNameEPTwoTwoFiveSixTwoTFourW),
		string(SkuNameEdge),
		string(SkuNameEdgeMRMini),
		string(SkuNameEdgeMRTCP),
		string(SkuNameEdgePBase),
		string(SkuNameEdgePHigh),
		string(SkuNameEdgePRBase),
		string(SkuNameEdgePRBaseUPS),
		string(SkuNameGPU),
		string(SkuNameGateway),
		string(SkuNameManagement),
		string(SkuNameRCALarge),
		string(SkuNameRCASmall),
		string(SkuNameRDC),
		string(SkuNameTCALarge),
		string(SkuNameTCASmall),
		string(SkuNameTDC),
		string(SkuNameTEAFourNodeHeater),
		string(SkuNameTEAFourNodeUPSHeater),
		string(SkuNameTEAOneNode),
		string(SkuNameTEAOneNodeHeater),
		string(SkuNameTEAOneNodeUPS),
		string(SkuNameTEAOneNodeUPSHeater),
		string(SkuNameTMA),
	}
}

type SkuTier string

const (
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierStandard),
	}
}

type SubscriptionState string

const (
	SubscriptionStateDeleted      SubscriptionState = "Deleted"
	SubscriptionStateRegistered   SubscriptionState = "Registered"
	SubscriptionStateSuspended    SubscriptionState = "Suspended"
	SubscriptionStateUnregistered SubscriptionState = "Unregistered"
	SubscriptionStateWarned       SubscriptionState = "Warned"
)

func PossibleValuesForSubscriptionState() []string {
	return []string{
		string(SubscriptionStateDeleted),
		string(SubscriptionStateRegistered),
		string(SubscriptionStateSuspended),
		string(SubscriptionStateUnregistered),
		string(SubscriptionStateWarned),
	}
}

type UpdateOperation string

const (
	UpdateOperationDownload UpdateOperation = "Download"
	UpdateOperationInstall  UpdateOperation = "Install"
	UpdateOperationNone     UpdateOperation = "None"
	UpdateOperationScan     UpdateOperation = "Scan"
)

func PossibleValuesForUpdateOperation() []string {
	return []string{
		string(UpdateOperationDownload),
		string(UpdateOperationInstall),
		string(UpdateOperationNone),
		string(UpdateOperationScan),
	}
}

type UpdateStatus string

const (
	UpdateStatusDownloadCompleted UpdateStatus = "DownloadCompleted"
	UpdateStatusDownloadPending   UpdateStatus = "DownloadPending"
	UpdateStatusDownloadStarted   UpdateStatus = "DownloadStarted"
	UpdateStatusInstallCompleted  UpdateStatus = "InstallCompleted"
	UpdateStatusInstallStarted    UpdateStatus = "InstallStarted"
)

func PossibleValuesForUpdateStatus() []string {
	return []string{
		string(UpdateStatusDownloadCompleted),
		string(UpdateStatusDownloadPending),
		string(UpdateStatusDownloadStarted),
		string(UpdateStatusInstallCompleted),
		string(UpdateStatusInstallStarted),
	}
}

type UpdateType string

const (
	UpdateTypeFirmware   UpdateType = "Firmware"
	UpdateTypeKubernetes UpdateType = "Kubernetes"
	UpdateTypeSoftware   UpdateType = "Software"
)

func PossibleValuesForUpdateType() []string {
	return []string{
		string(UpdateTypeFirmware),
		string(UpdateTypeKubernetes),
		string(UpdateTypeSoftware),
	}
}
