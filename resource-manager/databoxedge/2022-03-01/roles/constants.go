package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type HostPlatformType string

const (
	HostPlatformTypeKubernetesCluster HostPlatformType = "KubernetesCluster"
	HostPlatformTypeLinuxVM           HostPlatformType = "LinuxVM"
)

func PossibleValuesForHostPlatformType() []string {
	return []string{
		string(HostPlatformTypeKubernetesCluster),
		string(HostPlatformTypeLinuxVM),
	}
}

type KubernetesNodeType string

const (
	KubernetesNodeTypeInvalid KubernetesNodeType = "Invalid"
	KubernetesNodeTypeMaster  KubernetesNodeType = "Master"
	KubernetesNodeTypeWorker  KubernetesNodeType = "Worker"
)

func PossibleValuesForKubernetesNodeType() []string {
	return []string{
		string(KubernetesNodeTypeInvalid),
		string(KubernetesNodeTypeMaster),
		string(KubernetesNodeTypeWorker),
	}
}

type KubernetesState string

const (
	KubernetesStateCreated       KubernetesState = "Created"
	KubernetesStateCreating      KubernetesState = "Creating"
	KubernetesStateDeleting      KubernetesState = "Deleting"
	KubernetesStateFailed        KubernetesState = "Failed"
	KubernetesStateInvalid       KubernetesState = "Invalid"
	KubernetesStateReconfiguring KubernetesState = "Reconfiguring"
	KubernetesStateUpdating      KubernetesState = "Updating"
)

func PossibleValuesForKubernetesState() []string {
	return []string{
		string(KubernetesStateCreated),
		string(KubernetesStateCreating),
		string(KubernetesStateDeleting),
		string(KubernetesStateFailed),
		string(KubernetesStateInvalid),
		string(KubernetesStateReconfiguring),
		string(KubernetesStateUpdating),
	}
}

type MountType string

const (
	MountTypeHostPath MountType = "HostPath"
	MountTypeVolume   MountType = "Volume"
)

func PossibleValuesForMountType() []string {
	return []string{
		string(MountTypeHostPath),
		string(MountTypeVolume),
	}
}

type PlatformType string

const (
	PlatformTypeLinux   PlatformType = "Linux"
	PlatformTypeWindows PlatformType = "Windows"
)

func PossibleValuesForPlatformType() []string {
	return []string{
		string(PlatformTypeLinux),
		string(PlatformTypeWindows),
	}
}

type PosixComplianceStatus string

const (
	PosixComplianceStatusDisabled PosixComplianceStatus = "Disabled"
	PosixComplianceStatusEnabled  PosixComplianceStatus = "Enabled"
	PosixComplianceStatusInvalid  PosixComplianceStatus = "Invalid"
)

func PossibleValuesForPosixComplianceStatus() []string {
	return []string{
		string(PosixComplianceStatusDisabled),
		string(PosixComplianceStatusEnabled),
		string(PosixComplianceStatusInvalid),
	}
}

type RoleStatus string

const (
	RoleStatusDisabled RoleStatus = "Disabled"
	RoleStatusEnabled  RoleStatus = "Enabled"
)

func PossibleValuesForRoleStatus() []string {
	return []string{
		string(RoleStatusDisabled),
		string(RoleStatusEnabled),
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
