package appliances

import "strings"

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

func parseAccessProfileType(input string) (*AccessProfileType, error) {
	vals := map[string]AccessProfileType{
		"clustercustomeruser": AccessProfileTypeClusterCustomerUser,
		"clusteruser":         AccessProfileTypeClusterUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessProfileType(input)
	return &out, nil
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

func parseDistro(input string) (*Distro, error) {
	vals := map[string]Distro{
		"aksedge": DistroAKSEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Distro(input)
	return &out, nil
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

func parseProvider(input string) (*Provider, error) {
	vals := map[string]Provider{
		"hci":       ProviderHCI,
		"kubevirt":  ProviderKubeVirt,
		"openstack": ProviderOpenStack,
		"scvmm":     ProviderSCVMM,
		"vmware":    ProviderVMWare,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Provider(input)
	return &out, nil
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

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"connected":                             StatusConnected,
		"connecting":                            StatusConnecting,
		"imagedeprovisioning":                   StatusImageDeprovisioning,
		"imagedownloaded":                       StatusImageDownloaded,
		"imagedownloading":                      StatusImageDownloading,
		"imagepending":                          StatusImagePending,
		"imageprovisioned":                      StatusImageProvisioned,
		"imageprovisioning":                     StatusImageProvisioning,
		"imageunknown":                          StatusImageUnknown,
		"none":                                  StatusNone,
		"offline":                               StatusOffline,
		"postupgrade":                           StatusPostUpgrade,
		"preupgrade":                            StatusPreUpgrade,
		"preparingforupgrade":                   StatusPreparingForUpgrade,
		"running":                               StatusRunning,
		"updatingcapi":                          StatusUpdatingCAPI,
		"updatingcloudoperator":                 StatusUpdatingCloudOperator,
		"updatingcluster":                       StatusUpdatingCluster,
		"upgradeclusterextensionfailedtodelete": StatusUpgradeClusterExtensionFailedToDelete,
		"upgradecomplete":                       StatusUpgradeComplete,
		"upgradefailed":                         StatusUpgradeFailed,
		"upgradeprerequisitescompleted":         StatusUpgradePrerequisitesCompleted,
		"upgradingkvaio":                        StatusUpgradingKVAIO,
		"validating":                            StatusValidating,
		"waitingforcloudoperator":               StatusWaitingForCloudOperator,
		"waitingforheartbeat":                   StatusWaitingForHeartbeat,
		"waitingforkvaio":                       StatusWaitingForKVAIO,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}
