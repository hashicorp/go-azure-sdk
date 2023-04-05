package appliances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessProfileType string

const (
	AccessProfileTypeClusterCustomerUser AccessProfileType = "clusterCustomerUser"
	AccessProfileTypeClusterUser         AccessProfileType = "clusterUser"
)

func PossibleValuesForAccessProfileType() []string {
	return []string{
		string(AccessProfileTypeClusterCustomerUser),
		string(AccessProfileTypeClusterUser),
	}
}

type Distro string

const (
	DistroAKSEdge Distro = "AKSEdge"
)

func PossibleValuesForDistro() []string {
	return []string{
		string(DistroAKSEdge),
	}
}

type Provider string

const (
	ProviderHCI       Provider = "HCI"
	ProviderKubeVirt  Provider = "KubeVirt"
	ProviderOpenStack Provider = "OpenStack"
	ProviderSCVMM     Provider = "SCVMM"
	ProviderVMWare    Provider = "VMWare"
)

func PossibleValuesForProvider() []string {
	return []string{
		string(ProviderHCI),
		string(ProviderKubeVirt),
		string(ProviderOpenStack),
		string(ProviderSCVMM),
		string(ProviderVMWare),
	}
}

type Status string

const (
	StatusConnected                             Status = "Connected"
	StatusConnecting                            Status = "Connecting"
	StatusImageDeprovisioning                   Status = "ImageDeprovisioning"
	StatusImageDownloaded                       Status = "ImageDownloaded"
	StatusImageDownloading                      Status = "ImageDownloading"
	StatusImagePending                          Status = "ImagePending"
	StatusImageProvisioned                      Status = "ImageProvisioned"
	StatusImageProvisioning                     Status = "ImageProvisioning"
	StatusImageUnknown                          Status = "ImageUnknown"
	StatusNone                                  Status = "None"
	StatusOffline                               Status = "Offline"
	StatusPostUpgrade                           Status = "PostUpgrade"
	StatusPreUpgrade                            Status = "PreUpgrade"
	StatusPreparingForUpgrade                   Status = "PreparingForUpgrade"
	StatusRunning                               Status = "Running"
	StatusUpdatingCAPI                          Status = "UpdatingCAPI"
	StatusUpdatingCloudOperator                 Status = "UpdatingCloudOperator"
	StatusUpdatingCluster                       Status = "UpdatingCluster"
	StatusUpgradeClusterExtensionFailedToDelete Status = "UpgradeClusterExtensionFailedToDelete"
	StatusUpgradeComplete                       Status = "UpgradeComplete"
	StatusUpgradeFailed                         Status = "UpgradeFailed"
	StatusUpgradePrerequisitesCompleted         Status = "UpgradePrerequisitesCompleted"
	StatusUpgradingKVAIO                        Status = "UpgradingKVAIO"
	StatusValidating                            Status = "Validating"
	StatusWaitingForCloudOperator               Status = "WaitingForCloudOperator"
	StatusWaitingForHeartbeat                   Status = "WaitingForHeartbeat"
	StatusWaitingForKVAIO                       Status = "WaitingForKVAIO"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusConnected),
		string(StatusConnecting),
		string(StatusImageDeprovisioning),
		string(StatusImageDownloaded),
		string(StatusImageDownloading),
		string(StatusImagePending),
		string(StatusImageProvisioned),
		string(StatusImageProvisioning),
		string(StatusImageUnknown),
		string(StatusNone),
		string(StatusOffline),
		string(StatusPostUpgrade),
		string(StatusPreUpgrade),
		string(StatusPreparingForUpgrade),
		string(StatusRunning),
		string(StatusUpdatingCAPI),
		string(StatusUpdatingCloudOperator),
		string(StatusUpdatingCluster),
		string(StatusUpgradeClusterExtensionFailedToDelete),
		string(StatusUpgradeComplete),
		string(StatusUpgradeFailed),
		string(StatusUpgradePrerequisitesCompleted),
		string(StatusUpgradingKVAIO),
		string(StatusValidating),
		string(StatusWaitingForCloudOperator),
		string(StatusWaitingForHeartbeat),
		string(StatusWaitingForKVAIO),
	}
}
