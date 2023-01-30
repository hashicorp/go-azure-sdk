package roles

import "strings"

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

func parseEncryptionAlgorithm(input string) (*EncryptionAlgorithm, error) {
	vals := map[string]EncryptionAlgorithm{
		"aes256":            EncryptionAlgorithmAESTwoFiveSix,
		"none":              EncryptionAlgorithmNone,
		"rsaes_pkcs1_v_1_5": EncryptionAlgorithmRSAESPKCSOneVOneFive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionAlgorithm(input)
	return &out, nil
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

func parseHostPlatformType(input string) (*HostPlatformType, error) {
	vals := map[string]HostPlatformType{
		"kubernetescluster": HostPlatformTypeKubernetesCluster,
		"linuxvm":           HostPlatformTypeLinuxVM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostPlatformType(input)
	return &out, nil
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

func parseKubernetesNodeType(input string) (*KubernetesNodeType, error) {
	vals := map[string]KubernetesNodeType{
		"invalid": KubernetesNodeTypeInvalid,
		"master":  KubernetesNodeTypeMaster,
		"worker":  KubernetesNodeTypeWorker,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesNodeType(input)
	return &out, nil
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

func parseKubernetesState(input string) (*KubernetesState, error) {
	vals := map[string]KubernetesState{
		"created":       KubernetesStateCreated,
		"creating":      KubernetesStateCreating,
		"deleting":      KubernetesStateDeleting,
		"failed":        KubernetesStateFailed,
		"invalid":       KubernetesStateInvalid,
		"reconfiguring": KubernetesStateReconfiguring,
		"updating":      KubernetesStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesState(input)
	return &out, nil
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

func parseMountType(input string) (*MountType, error) {
	vals := map[string]MountType{
		"hostpath": MountTypeHostPath,
		"volume":   MountTypeVolume,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MountType(input)
	return &out, nil
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

func parsePlatformType(input string) (*PlatformType, error) {
	vals := map[string]PlatformType{
		"linux":   PlatformTypeLinux,
		"windows": PlatformTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlatformType(input)
	return &out, nil
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

func parsePosixComplianceStatus(input string) (*PosixComplianceStatus, error) {
	vals := map[string]PosixComplianceStatus{
		"disabled": PosixComplianceStatusDisabled,
		"enabled":  PosixComplianceStatusEnabled,
		"invalid":  PosixComplianceStatusInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PosixComplianceStatus(input)
	return &out, nil
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

func parseRoleStatus(input string) (*RoleStatus, error) {
	vals := map[string]RoleStatus{
		"disabled": RoleStatusDisabled,
		"enabled":  RoleStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleStatus(input)
	return &out, nil
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

func parseRoleTypes(input string) (*RoleTypes, error) {
	vals := map[string]RoleTypes{
		"asa":                 RoleTypesASA,
		"cloudedgemanagement": RoleTypesCloudEdgeManagement,
		"cognitive":           RoleTypesCognitive,
		"functions":           RoleTypesFunctions,
		"iot":                 RoleTypesIOT,
		"kubernetes":          RoleTypesKubernetes,
		"mec":                 RoleTypesMEC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleTypes(input)
	return &out, nil
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

func parseSubscriptionState(input string) (*SubscriptionState, error) {
	vals := map[string]SubscriptionState{
		"deleted":      SubscriptionStateDeleted,
		"registered":   SubscriptionStateRegistered,
		"suspended":    SubscriptionStateSuspended,
		"unregistered": SubscriptionStateUnregistered,
		"warned":       SubscriptionStateWarned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscriptionState(input)
	return &out, nil
}
