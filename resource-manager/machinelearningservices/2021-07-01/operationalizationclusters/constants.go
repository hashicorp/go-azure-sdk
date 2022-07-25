package operationalizationclusters

import "strings"

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

func parseAllocationState(input string) (*AllocationState, error) {
	vals := map[string]AllocationState{
		"resizing": AllocationStateResizing,
		"steady":   AllocationStateSteady,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllocationState(input)
	return &out, nil
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

func parseApplicationSharingPolicy(input string) (*ApplicationSharingPolicy, error) {
	vals := map[string]ApplicationSharingPolicy{
		"personal": ApplicationSharingPolicyPersonal,
		"shared":   ApplicationSharingPolicyShared,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplicationSharingPolicy(input)
	return &out, nil
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

func parseClusterPurpose(input string) (*ClusterPurpose, error) {
	vals := map[string]ClusterPurpose{
		"denseprod": ClusterPurposeDenseProd,
		"devtest":   ClusterPurposeDevTest,
		"fastprod":  ClusterPurposeFastProd,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterPurpose(input)
	return &out, nil
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

func parseComputeInstanceAuthorizationType(input string) (*ComputeInstanceAuthorizationType, error) {
	vals := map[string]ComputeInstanceAuthorizationType{
		"personal": ComputeInstanceAuthorizationTypePersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComputeInstanceAuthorizationType(input)
	return &out, nil
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

func parseComputeInstanceState(input string) (*ComputeInstanceState, error) {
	vals := map[string]ComputeInstanceState{
		"createfailed":    ComputeInstanceStateCreateFailed,
		"creating":        ComputeInstanceStateCreating,
		"deleting":        ComputeInstanceStateDeleting,
		"jobrunning":      ComputeInstanceStateJobRunning,
		"restarting":      ComputeInstanceStateRestarting,
		"running":         ComputeInstanceStateRunning,
		"settingup":       ComputeInstanceStateSettingUp,
		"setupfailed":     ComputeInstanceStateSetupFailed,
		"starting":        ComputeInstanceStateStarting,
		"stopped":         ComputeInstanceStateStopped,
		"stopping":        ComputeInstanceStateStopping,
		"unknown":         ComputeInstanceStateUnknown,
		"unusable":        ComputeInstanceStateUnusable,
		"usersettingup":   ComputeInstanceStateUserSettingUp,
		"usersetupfailed": ComputeInstanceStateUserSetupFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComputeInstanceState(input)
	return &out, nil
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

func parseComputeType(input string) (*ComputeType, error) {
	vals := map[string]ComputeType{
		"aks":               ComputeTypeAKS,
		"amlcompute":        ComputeTypeAmlCompute,
		"computeinstance":   ComputeTypeComputeInstance,
		"datafactory":       ComputeTypeDataFactory,
		"datalakeanalytics": ComputeTypeDataLakeAnalytics,
		"databricks":        ComputeTypeDatabricks,
		"hdinsight":         ComputeTypeHDInsight,
		"kubernetes":        ComputeTypeKubernetes,
		"synapsespark":      ComputeTypeSynapseSpark,
		"virtualmachine":    ComputeTypeVirtualMachine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComputeType(input)
	return &out, nil
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

func parseLoadBalancerType(input string) (*LoadBalancerType, error) {
	vals := map[string]LoadBalancerType{
		"internalloadbalancer": LoadBalancerTypeInternalLoadBalancer,
		"publicip":             LoadBalancerTypePublicIP,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoadBalancerType(input)
	return &out, nil
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

func parseOperationName(input string) (*OperationName, error) {
	vals := map[string]OperationName{
		"create":  OperationNameCreate,
		"delete":  OperationNameDelete,
		"reimage": OperationNameReimage,
		"restart": OperationNameRestart,
		"start":   OperationNameStart,
		"stop":    OperationNameStop,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationName(input)
	return &out, nil
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

func parseOperationStatus(input string) (*OperationStatus, error) {
	vals := map[string]OperationStatus{
		"createfailed":  OperationStatusCreateFailed,
		"deletefailed":  OperationStatusDeleteFailed,
		"inprogress":    OperationStatusInProgress,
		"reimagefailed": OperationStatusReimageFailed,
		"restartfailed": OperationStatusRestartFailed,
		"startfailed":   OperationStatusStartFailed,
		"stopfailed":    OperationStatusStopFailed,
		"succeeded":     OperationStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationStatus(input)
	return &out, nil
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

func parseOsType(input string) (*OsType, error) {
	vals := map[string]OsType{
		"linux":   OsTypeLinux,
		"windows": OsTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OsType(input)
	return &out, nil
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

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":  ProvisioningStateCanceled,
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
		"unknown":   ProvisioningStateUnknown,
		"updating":  ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
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

func parseRemoteLoginPortPublicAccess(input string) (*RemoteLoginPortPublicAccess, error) {
	vals := map[string]RemoteLoginPortPublicAccess{
		"disabled":     RemoteLoginPortPublicAccessDisabled,
		"enabled":      RemoteLoginPortPublicAccessEnabled,
		"notspecified": RemoteLoginPortPublicAccessNotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteLoginPortPublicAccess(input)
	return &out, nil
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

func parseSshPublicAccess(input string) (*SshPublicAccess, error) {
	vals := map[string]SshPublicAccess{
		"disabled": SshPublicAccessDisabled,
		"enabled":  SshPublicAccessEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SshPublicAccess(input)
	return &out, nil
}

type Status string

const (
	StatusAuto     Status = "Auto"
	StatusDisabled Status = "Disabled"
	StatusEnabled  Status = "Enabled"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusAuto),
		string(StatusDisabled),
		string(StatusEnabled),
	}
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"auto":     StatusAuto,
		"disabled": StatusDisabled,
		"enabled":  StatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
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

func parseUnderlyingResourceAction(input string) (*UnderlyingResourceAction, error) {
	vals := map[string]UnderlyingResourceAction{
		"delete": UnderlyingResourceActionDelete,
		"detach": UnderlyingResourceActionDetach,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnderlyingResourceAction(input)
	return &out, nil
}

type VmPriority string

const (
	VmPriorityDedicated   VmPriority = "Dedicated"
	VmPriorityLowPriority VmPriority = "LowPriority"
)

func PossibleValuesForVmPriority() []string {
	return []string{
		string(VmPriorityDedicated),
		string(VmPriorityLowPriority),
	}
}

func parseVmPriority(input string) (*VmPriority, error) {
	vals := map[string]VmPriority{
		"dedicated":   VmPriorityDedicated,
		"lowpriority": VmPriorityLowPriority,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VmPriority(input)
	return &out, nil
}
