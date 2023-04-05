package addons

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
