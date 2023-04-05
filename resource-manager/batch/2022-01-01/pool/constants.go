package pool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllocationState string

const (
	AllocationStateResizing AllocationState = "Resizing"
	AllocationStateSteady   AllocationState = "Steady"
	AllocationStateStopping AllocationState = "Stopping"
)

func PossibleValuesForAllocationState() []string {
	return []string{
		string(AllocationStateResizing),
		string(AllocationStateSteady),
		string(AllocationStateStopping),
	}
}

type AutoUserScope string

const (
	AutoUserScopePool AutoUserScope = "Pool"
	AutoUserScopeTask AutoUserScope = "Task"
)

func PossibleValuesForAutoUserScope() []string {
	return []string{
		string(AutoUserScopePool),
		string(AutoUserScopeTask),
	}
}

type CachingType string

const (
	CachingTypeNone      CachingType = "None"
	CachingTypeReadOnly  CachingType = "ReadOnly"
	CachingTypeReadWrite CachingType = "ReadWrite"
)

func PossibleValuesForCachingType() []string {
	return []string{
		string(CachingTypeNone),
		string(CachingTypeReadOnly),
		string(CachingTypeReadWrite),
	}
}

type CertificateStoreLocation string

const (
	CertificateStoreLocationCurrentUser  CertificateStoreLocation = "CurrentUser"
	CertificateStoreLocationLocalMachine CertificateStoreLocation = "LocalMachine"
)

func PossibleValuesForCertificateStoreLocation() []string {
	return []string{
		string(CertificateStoreLocationCurrentUser),
		string(CertificateStoreLocationLocalMachine),
	}
}

type CertificateVisibility string

const (
	CertificateVisibilityRemoteUser CertificateVisibility = "RemoteUser"
	CertificateVisibilityStartTask  CertificateVisibility = "StartTask"
	CertificateVisibilityTask       CertificateVisibility = "Task"
)

func PossibleValuesForCertificateVisibility() []string {
	return []string{
		string(CertificateVisibilityRemoteUser),
		string(CertificateVisibilityStartTask),
		string(CertificateVisibilityTask),
	}
}

type ComputeNodeDeallocationOption string

const (
	ComputeNodeDeallocationOptionRequeue        ComputeNodeDeallocationOption = "Requeue"
	ComputeNodeDeallocationOptionRetainedData   ComputeNodeDeallocationOption = "RetainedData"
	ComputeNodeDeallocationOptionTaskCompletion ComputeNodeDeallocationOption = "TaskCompletion"
	ComputeNodeDeallocationOptionTerminate      ComputeNodeDeallocationOption = "Terminate"
)

func PossibleValuesForComputeNodeDeallocationOption() []string {
	return []string{
		string(ComputeNodeDeallocationOptionRequeue),
		string(ComputeNodeDeallocationOptionRetainedData),
		string(ComputeNodeDeallocationOptionTaskCompletion),
		string(ComputeNodeDeallocationOptionTerminate),
	}
}

type ComputeNodeFillType string

const (
	ComputeNodeFillTypePack   ComputeNodeFillType = "Pack"
	ComputeNodeFillTypeSpread ComputeNodeFillType = "Spread"
)

func PossibleValuesForComputeNodeFillType() []string {
	return []string{
		string(ComputeNodeFillTypePack),
		string(ComputeNodeFillTypeSpread),
	}
}

type ContainerType string

const (
	ContainerTypeDockerCompatible ContainerType = "DockerCompatible"
)

func PossibleValuesForContainerType() []string {
	return []string{
		string(ContainerTypeDockerCompatible),
	}
}

type ContainerWorkingDirectory string

const (
	ContainerWorkingDirectoryContainerImageDefault ContainerWorkingDirectory = "ContainerImageDefault"
	ContainerWorkingDirectoryTaskWorkingDirectory  ContainerWorkingDirectory = "TaskWorkingDirectory"
)

func PossibleValuesForContainerWorkingDirectory() []string {
	return []string{
		string(ContainerWorkingDirectoryContainerImageDefault),
		string(ContainerWorkingDirectoryTaskWorkingDirectory),
	}
}

type DiffDiskPlacement string

const (
	DiffDiskPlacementCacheDisk DiffDiskPlacement = "CacheDisk"
)

func PossibleValuesForDiffDiskPlacement() []string {
	return []string{
		string(DiffDiskPlacementCacheDisk),
	}
}

type DiskEncryptionTarget string

const (
	DiskEncryptionTargetOsDisk        DiskEncryptionTarget = "OsDisk"
	DiskEncryptionTargetTemporaryDisk DiskEncryptionTarget = "TemporaryDisk"
)

func PossibleValuesForDiskEncryptionTarget() []string {
	return []string{
		string(DiskEncryptionTargetOsDisk),
		string(DiskEncryptionTargetTemporaryDisk),
	}
}

type DynamicVNetAssignmentScope string

const (
	DynamicVNetAssignmentScopeJob  DynamicVNetAssignmentScope = "job"
	DynamicVNetAssignmentScopeNone DynamicVNetAssignmentScope = "none"
)

func PossibleValuesForDynamicVNetAssignmentScope() []string {
	return []string{
		string(DynamicVNetAssignmentScopeJob),
		string(DynamicVNetAssignmentScopeNone),
	}
}

type ElevationLevel string

const (
	ElevationLevelAdmin    ElevationLevel = "Admin"
	ElevationLevelNonAdmin ElevationLevel = "NonAdmin"
)

func PossibleValuesForElevationLevel() []string {
	return []string{
		string(ElevationLevelAdmin),
		string(ElevationLevelNonAdmin),
	}
}

type IPAddressProvisioningType string

const (
	IPAddressProvisioningTypeBatchManaged        IPAddressProvisioningType = "BatchManaged"
	IPAddressProvisioningTypeNoPublicIPAddresses IPAddressProvisioningType = "NoPublicIPAddresses"
	IPAddressProvisioningTypeUserManaged         IPAddressProvisioningType = "UserManaged"
)

func PossibleValuesForIPAddressProvisioningType() []string {
	return []string{
		string(IPAddressProvisioningTypeBatchManaged),
		string(IPAddressProvisioningTypeNoPublicIPAddresses),
		string(IPAddressProvisioningTypeUserManaged),
	}
}

type InboundEndpointProtocol string

const (
	InboundEndpointProtocolTCP InboundEndpointProtocol = "TCP"
	InboundEndpointProtocolUDP InboundEndpointProtocol = "UDP"
)

func PossibleValuesForInboundEndpointProtocol() []string {
	return []string{
		string(InboundEndpointProtocolTCP),
		string(InboundEndpointProtocolUDP),
	}
}

type InterNodeCommunicationState string

const (
	InterNodeCommunicationStateDisabled InterNodeCommunicationState = "Disabled"
	InterNodeCommunicationStateEnabled  InterNodeCommunicationState = "Enabled"
)

func PossibleValuesForInterNodeCommunicationState() []string {
	return []string{
		string(InterNodeCommunicationStateDisabled),
		string(InterNodeCommunicationStateEnabled),
	}
}

type LoginMode string

const (
	LoginModeBatch       LoginMode = "Batch"
	LoginModeInteractive LoginMode = "Interactive"
)

func PossibleValuesForLoginMode() []string {
	return []string{
		string(LoginModeBatch),
		string(LoginModeInteractive),
	}
}

type NetworkSecurityGroupRuleAccess string

const (
	NetworkSecurityGroupRuleAccessAllow NetworkSecurityGroupRuleAccess = "Allow"
	NetworkSecurityGroupRuleAccessDeny  NetworkSecurityGroupRuleAccess = "Deny"
)

func PossibleValuesForNetworkSecurityGroupRuleAccess() []string {
	return []string{
		string(NetworkSecurityGroupRuleAccessAllow),
		string(NetworkSecurityGroupRuleAccessDeny),
	}
}

type NodePlacementPolicyType string

const (
	NodePlacementPolicyTypeRegional NodePlacementPolicyType = "Regional"
	NodePlacementPolicyTypeZonal    NodePlacementPolicyType = "Zonal"
)

func PossibleValuesForNodePlacementPolicyType() []string {
	return []string{
		string(NodePlacementPolicyTypeRegional),
		string(NodePlacementPolicyTypeZonal),
	}
}

type PoolProvisioningState string

const (
	PoolProvisioningStateDeleting  PoolProvisioningState = "Deleting"
	PoolProvisioningStateSucceeded PoolProvisioningState = "Succeeded"
)

func PossibleValuesForPoolProvisioningState() []string {
	return []string{
		string(PoolProvisioningStateDeleting),
		string(PoolProvisioningStateSucceeded),
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
