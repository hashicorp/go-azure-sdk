package operationalizationclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllocationState string

const (
	AllocationStateResizing AllocationState = "Resizing"
	AllocationStateSteady   AllocationState = "Steady"
)

func PossibleValuesForAllocationState() []string {
	return []string{
		string(AllocationStateResizing),
		string(AllocationStateSteady),
	}
}

type ApplicationSharingPolicy string

const (
	ApplicationSharingPolicyPersonal ApplicationSharingPolicy = "Personal"
	ApplicationSharingPolicyShared   ApplicationSharingPolicy = "Shared"
)

func PossibleValuesForApplicationSharingPolicy() []string {
	return []string{
		string(ApplicationSharingPolicyPersonal),
		string(ApplicationSharingPolicyShared),
	}
}

type Autosave string

const (
	AutosaveLocal  Autosave = "Local"
	AutosaveNone   Autosave = "None"
	AutosaveRemote Autosave = "Remote"
)

func PossibleValuesForAutosave() []string {
	return []string{
		string(AutosaveLocal),
		string(AutosaveNone),
		string(AutosaveRemote),
	}
}

type Caching string

const (
	CachingNone      Caching = "None"
	CachingReadOnly  Caching = "ReadOnly"
	CachingReadWrite Caching = "ReadWrite"
)

func PossibleValuesForCaching() []string {
	return []string{
		string(CachingNone),
		string(CachingReadOnly),
		string(CachingReadWrite),
	}
}

type ClusterPurpose string

const (
	ClusterPurposeDenseProd ClusterPurpose = "DenseProd"
	ClusterPurposeDevTest   ClusterPurpose = "DevTest"
	ClusterPurposeFastProd  ClusterPurpose = "FastProd"
)

func PossibleValuesForClusterPurpose() []string {
	return []string{
		string(ClusterPurposeDenseProd),
		string(ClusterPurposeDevTest),
		string(ClusterPurposeFastProd),
	}
}

type ComputeInstanceAuthorizationType string

const (
	ComputeInstanceAuthorizationTypePersonal ComputeInstanceAuthorizationType = "personal"
)

func PossibleValuesForComputeInstanceAuthorizationType() []string {
	return []string{
		string(ComputeInstanceAuthorizationTypePersonal),
	}
}

type ComputeInstanceState string

const (
	ComputeInstanceStateCreateFailed    ComputeInstanceState = "CreateFailed"
	ComputeInstanceStateCreating        ComputeInstanceState = "Creating"
	ComputeInstanceStateDeleting        ComputeInstanceState = "Deleting"
	ComputeInstanceStateJobRunning      ComputeInstanceState = "JobRunning"
	ComputeInstanceStateRestarting      ComputeInstanceState = "Restarting"
	ComputeInstanceStateRunning         ComputeInstanceState = "Running"
	ComputeInstanceStateSettingUp       ComputeInstanceState = "SettingUp"
	ComputeInstanceStateSetupFailed     ComputeInstanceState = "SetupFailed"
	ComputeInstanceStateStarting        ComputeInstanceState = "Starting"
	ComputeInstanceStateStopped         ComputeInstanceState = "Stopped"
	ComputeInstanceStateStopping        ComputeInstanceState = "Stopping"
	ComputeInstanceStateUnknown         ComputeInstanceState = "Unknown"
	ComputeInstanceStateUnusable        ComputeInstanceState = "Unusable"
	ComputeInstanceStateUserSettingUp   ComputeInstanceState = "UserSettingUp"
	ComputeInstanceStateUserSetupFailed ComputeInstanceState = "UserSetupFailed"
)

func PossibleValuesForComputeInstanceState() []string {
	return []string{
		string(ComputeInstanceStateCreateFailed),
		string(ComputeInstanceStateCreating),
		string(ComputeInstanceStateDeleting),
		string(ComputeInstanceStateJobRunning),
		string(ComputeInstanceStateRestarting),
		string(ComputeInstanceStateRunning),
		string(ComputeInstanceStateSettingUp),
		string(ComputeInstanceStateSetupFailed),
		string(ComputeInstanceStateStarting),
		string(ComputeInstanceStateStopped),
		string(ComputeInstanceStateStopping),
		string(ComputeInstanceStateUnknown),
		string(ComputeInstanceStateUnusable),
		string(ComputeInstanceStateUserSettingUp),
		string(ComputeInstanceStateUserSetupFailed),
	}
}

type ComputePowerAction string

const (
	ComputePowerActionStart ComputePowerAction = "Start"
	ComputePowerActionStop  ComputePowerAction = "Stop"
)

func PossibleValuesForComputePowerAction() []string {
	return []string{
		string(ComputePowerActionStart),
		string(ComputePowerActionStop),
	}
}

type ComputeType string

const (
	ComputeTypeAKS               ComputeType = "AKS"
	ComputeTypeAmlCompute        ComputeType = "AmlCompute"
	ComputeTypeComputeInstance   ComputeType = "ComputeInstance"
	ComputeTypeDataFactory       ComputeType = "DataFactory"
	ComputeTypeDataLakeAnalytics ComputeType = "DataLakeAnalytics"
	ComputeTypeDatabricks        ComputeType = "Databricks"
	ComputeTypeHDInsight         ComputeType = "HDInsight"
	ComputeTypeKubernetes        ComputeType = "Kubernetes"
	ComputeTypeSynapseSpark      ComputeType = "SynapseSpark"
	ComputeTypeVirtualMachine    ComputeType = "VirtualMachine"
)

func PossibleValuesForComputeType() []string {
	return []string{
		string(ComputeTypeAKS),
		string(ComputeTypeAmlCompute),
		string(ComputeTypeComputeInstance),
		string(ComputeTypeDataFactory),
		string(ComputeTypeDataLakeAnalytics),
		string(ComputeTypeDatabricks),
		string(ComputeTypeHDInsight),
		string(ComputeTypeKubernetes),
		string(ComputeTypeSynapseSpark),
		string(ComputeTypeVirtualMachine),
	}
}

type LoadBalancerType string

const (
	LoadBalancerTypeInternalLoadBalancer LoadBalancerType = "InternalLoadBalancer"
	LoadBalancerTypePublicIP             LoadBalancerType = "PublicIp"
)

func PossibleValuesForLoadBalancerType() []string {
	return []string{
		string(LoadBalancerTypeInternalLoadBalancer),
		string(LoadBalancerTypePublicIP),
	}
}

type MountAction string

const (
	MountActionMount   MountAction = "Mount"
	MountActionUnmount MountAction = "Unmount"
)

func PossibleValuesForMountAction() []string {
	return []string{
		string(MountActionMount),
		string(MountActionUnmount),
	}
}

type MountState string

const (
	MountStateMountFailed      MountState = "MountFailed"
	MountStateMountRequested   MountState = "MountRequested"
	MountStateMounted          MountState = "Mounted"
	MountStateUnmountFailed    MountState = "UnmountFailed"
	MountStateUnmountRequested MountState = "UnmountRequested"
	MountStateUnmounted        MountState = "Unmounted"
)

func PossibleValuesForMountState() []string {
	return []string{
		string(MountStateMountFailed),
		string(MountStateMountRequested),
		string(MountStateMounted),
		string(MountStateUnmountFailed),
		string(MountStateUnmountRequested),
		string(MountStateUnmounted),
	}
}

type Network string

const (
	NetworkBridge Network = "Bridge"
	NetworkHost   Network = "Host"
)

func PossibleValuesForNetwork() []string {
	return []string{
		string(NetworkBridge),
		string(NetworkHost),
	}
}

type OperationName string

const (
	OperationNameCreate  OperationName = "Create"
	OperationNameDelete  OperationName = "Delete"
	OperationNameReimage OperationName = "Reimage"
	OperationNameRestart OperationName = "Restart"
	OperationNameStart   OperationName = "Start"
	OperationNameStop    OperationName = "Stop"
)

func PossibleValuesForOperationName() []string {
	return []string{
		string(OperationNameCreate),
		string(OperationNameDelete),
		string(OperationNameReimage),
		string(OperationNameRestart),
		string(OperationNameStart),
		string(OperationNameStop),
	}
}

type OperationStatus string

const (
	OperationStatusCreateFailed  OperationStatus = "CreateFailed"
	OperationStatusDeleteFailed  OperationStatus = "DeleteFailed"
	OperationStatusInProgress    OperationStatus = "InProgress"
	OperationStatusReimageFailed OperationStatus = "ReimageFailed"
	OperationStatusRestartFailed OperationStatus = "RestartFailed"
	OperationStatusStartFailed   OperationStatus = "StartFailed"
	OperationStatusStopFailed    OperationStatus = "StopFailed"
	OperationStatusSucceeded     OperationStatus = "Succeeded"
)

func PossibleValuesForOperationStatus() []string {
	return []string{
		string(OperationStatusCreateFailed),
		string(OperationStatusDeleteFailed),
		string(OperationStatusInProgress),
		string(OperationStatusReimageFailed),
		string(OperationStatusRestartFailed),
		string(OperationStatusStartFailed),
		string(OperationStatusStopFailed),
		string(OperationStatusSucceeded),
	}
}

type OperationTrigger string

const (
	OperationTriggerIdleShutdown OperationTrigger = "IdleShutdown"
	OperationTriggerSchedule     OperationTrigger = "Schedule"
	OperationTriggerUser         OperationTrigger = "User"
)

func PossibleValuesForOperationTrigger() []string {
	return []string{
		string(OperationTriggerIdleShutdown),
		string(OperationTriggerSchedule),
		string(OperationTriggerUser),
	}
}

type OsType string

const (
	OsTypeLinux   OsType = "Linux"
	OsTypeWindows OsType = "Windows"
)

func PossibleValuesForOsType() []string {
	return []string{
		string(OsTypeLinux),
		string(OsTypeWindows),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUnknown),
		string(ProvisioningStateUpdating),
	}
}

type ProvisioningStatus string

const (
	ProvisioningStatusCompleted    ProvisioningStatus = "Completed"
	ProvisioningStatusFailed       ProvisioningStatus = "Failed"
	ProvisioningStatusProvisioning ProvisioningStatus = "Provisioning"
)

func PossibleValuesForProvisioningStatus() []string {
	return []string{
		string(ProvisioningStatusCompleted),
		string(ProvisioningStatusFailed),
		string(ProvisioningStatusProvisioning),
	}
}

type RemoteLoginPortPublicAccess string

const (
	RemoteLoginPortPublicAccessDisabled     RemoteLoginPortPublicAccess = "Disabled"
	RemoteLoginPortPublicAccessEnabled      RemoteLoginPortPublicAccess = "Enabled"
	RemoteLoginPortPublicAccessNotSpecified RemoteLoginPortPublicAccess = "NotSpecified"
)

func PossibleValuesForRemoteLoginPortPublicAccess() []string {
	return []string{
		string(RemoteLoginPortPublicAccessDisabled),
		string(RemoteLoginPortPublicAccessEnabled),
		string(RemoteLoginPortPublicAccessNotSpecified),
	}
}

type ScheduleProvisioningState string

const (
	ScheduleProvisioningStateCompleted    ScheduleProvisioningState = "Completed"
	ScheduleProvisioningStateFailed       ScheduleProvisioningState = "Failed"
	ScheduleProvisioningStateProvisioning ScheduleProvisioningState = "Provisioning"
)

func PossibleValuesForScheduleProvisioningState() []string {
	return []string{
		string(ScheduleProvisioningStateCompleted),
		string(ScheduleProvisioningStateFailed),
		string(ScheduleProvisioningStateProvisioning),
	}
}

type ScheduleStatus string

const (
	ScheduleStatusDisabled ScheduleStatus = "Disabled"
	ScheduleStatusEnabled  ScheduleStatus = "Enabled"
)

func PossibleValuesForScheduleStatus() []string {
	return []string{
		string(ScheduleStatusDisabled),
		string(ScheduleStatusEnabled),
	}
}

type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierFree     SkuTier = "Free"
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierFree),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}

type SourceType string

const (
	SourceTypeDataset   SourceType = "Dataset"
	SourceTypeDatastore SourceType = "Datastore"
	SourceTypeURI       SourceType = "URI"
)

func PossibleValuesForSourceType() []string {
	return []string{
		string(SourceTypeDataset),
		string(SourceTypeDatastore),
		string(SourceTypeURI),
	}
}

type SshPublicAccess string

const (
	SshPublicAccessDisabled SshPublicAccess = "Disabled"
	SshPublicAccessEnabled  SshPublicAccess = "Enabled"
)

func PossibleValuesForSshPublicAccess() []string {
	return []string{
		string(SshPublicAccessDisabled),
		string(SshPublicAccessEnabled),
	}
}

type SslConfigStatus string

const (
	SslConfigStatusAuto     SslConfigStatus = "Auto"
	SslConfigStatusDisabled SslConfigStatus = "Disabled"
	SslConfigStatusEnabled  SslConfigStatus = "Enabled"
)

func PossibleValuesForSslConfigStatus() []string {
	return []string{
		string(SslConfigStatusAuto),
		string(SslConfigStatusDisabled),
		string(SslConfigStatusEnabled),
	}
}

type StorageAccountType string

const (
	StorageAccountTypePremiumLRS  StorageAccountType = "Premium_LRS"
	StorageAccountTypeStandardLRS StorageAccountType = "Standard_LRS"
)

func PossibleValuesForStorageAccountType() []string {
	return []string{
		string(StorageAccountTypePremiumLRS),
		string(StorageAccountTypeStandardLRS),
	}
}

type UnderlyingResourceAction string

const (
	UnderlyingResourceActionDelete UnderlyingResourceAction = "Delete"
	UnderlyingResourceActionDetach UnderlyingResourceAction = "Detach"
)

func PossibleValuesForUnderlyingResourceAction() []string {
	return []string{
		string(UnderlyingResourceActionDelete),
		string(UnderlyingResourceActionDetach),
	}
}

type VMPriority string

const (
	VMPriorityDedicated   VMPriority = "Dedicated"
	VMPriorityLowPriority VMPriority = "LowPriority"
)

func PossibleValuesForVMPriority() []string {
	return []string{
		string(VMPriorityDedicated),
		string(VMPriorityLowPriority),
	}
}
