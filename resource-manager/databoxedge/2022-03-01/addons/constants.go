package addons

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddonState string

const (
	AddonStateCreated       AddonState = "Created"
	AddonStateCreating      AddonState = "Creating"
	AddonStateDeleting      AddonState = "Deleting"
	AddonStateFailed        AddonState = "Failed"
	AddonStateInvalid       AddonState = "Invalid"
	AddonStateReconfiguring AddonState = "Reconfiguring"
	AddonStateUpdating      AddonState = "Updating"
)

func PossibleValuesForAddonState() []string {
	return []string{
		string(AddonStateCreated),
		string(AddonStateCreating),
		string(AddonStateDeleting),
		string(AddonStateFailed),
		string(AddonStateInvalid),
		string(AddonStateReconfiguring),
		string(AddonStateUpdating),
	}
}

func parseAddonState(input string) (*AddonState, error) {
	vals := map[string]AddonState{
		"created":       AddonStateCreated,
		"creating":      AddonStateCreating,
		"deleting":      AddonStateDeleting,
		"failed":        AddonStateFailed,
		"invalid":       AddonStateInvalid,
		"reconfiguring": AddonStateReconfiguring,
		"updating":      AddonStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddonState(input)
	return &out, nil
}

type AddonType string

const (
	AddonTypeArcForKubernetes AddonType = "ArcForKubernetes"
	AddonTypeIotEdge          AddonType = "IotEdge"
)

func PossibleValuesForAddonType() []string {
	return []string{
		string(AddonTypeArcForKubernetes),
		string(AddonTypeIotEdge),
	}
}

func parseAddonType(input string) (*AddonType, error) {
	vals := map[string]AddonType{
		"arcforkubernetes": AddonTypeArcForKubernetes,
		"iotedge":          AddonTypeIotEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddonType(input)
	return &out, nil
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
