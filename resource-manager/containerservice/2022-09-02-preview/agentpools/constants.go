package agentpools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentPoolMode string

const (
	AgentPoolModeSystem AgentPoolMode = "System"
	AgentPoolModeUser   AgentPoolMode = "User"
)

func PossibleValuesForAgentPoolMode() []string {
	return []string{
		string(AgentPoolModeSystem),
		string(AgentPoolModeUser),
	}
}

type AgentPoolType string

const (
	AgentPoolTypeAvailabilitySet         AgentPoolType = "AvailabilitySet"
	AgentPoolTypeVirtualMachineScaleSets AgentPoolType = "VirtualMachineScaleSets"
)

func PossibleValuesForAgentPoolType() []string {
	return []string{
		string(AgentPoolTypeAvailabilitySet),
		string(AgentPoolTypeVirtualMachineScaleSets),
	}
}

type Code string

const (
	CodeRunning Code = "Running"
	CodeStopped Code = "Stopped"
)

func PossibleValuesForCode() []string {
	return []string{
		string(CodeRunning),
		string(CodeStopped),
	}
}

type GPUInstanceProfile string

const (
	GPUInstanceProfileMIGFourg  GPUInstanceProfile = "MIG4g"
	GPUInstanceProfileMIGOneg   GPUInstanceProfile = "MIG1g"
	GPUInstanceProfileMIGSeveng GPUInstanceProfile = "MIG7g"
	GPUInstanceProfileMIGThreeg GPUInstanceProfile = "MIG3g"
	GPUInstanceProfileMIGTwog   GPUInstanceProfile = "MIG2g"
)

func PossibleValuesForGPUInstanceProfile() []string {
	return []string{
		string(GPUInstanceProfileMIGFourg),
		string(GPUInstanceProfileMIGOneg),
		string(GPUInstanceProfileMIGSeveng),
		string(GPUInstanceProfileMIGThreeg),
		string(GPUInstanceProfileMIGTwog),
	}
}

type KubeletDiskType string

const (
	KubeletDiskTypeOS        KubeletDiskType = "OS"
	KubeletDiskTypeTemporary KubeletDiskType = "Temporary"
)

func PossibleValuesForKubeletDiskType() []string {
	return []string{
		string(KubeletDiskTypeOS),
		string(KubeletDiskTypeTemporary),
	}
}

type OSDiskType string

const (
	OSDiskTypeEphemeral OSDiskType = "Ephemeral"
	OSDiskTypeManaged   OSDiskType = "Managed"
)

func PossibleValuesForOSDiskType() []string {
	return []string{
		string(OSDiskTypeEphemeral),
		string(OSDiskTypeManaged),
	}
}

type OSSKU string

const (
	OSSKUCBLMariner            OSSKU = "CBLMariner"
	OSSKUMariner               OSSKU = "Mariner"
	OSSKUUbuntu                OSSKU = "Ubuntu"
	OSSKUWindowsTwoZeroOneNine OSSKU = "Windows2019"
	OSSKUWindowsTwoZeroTwoTwo  OSSKU = "Windows2022"
)

func PossibleValuesForOSSKU() []string {
	return []string{
		string(OSSKUCBLMariner),
		string(OSSKUMariner),
		string(OSSKUUbuntu),
		string(OSSKUWindowsTwoZeroOneNine),
		string(OSSKUWindowsTwoZeroTwoTwo),
	}
}

type OSType string

const (
	OSTypeLinux   OSType = "Linux"
	OSTypeWindows OSType = "Windows"
)

func PossibleValuesForOSType() []string {
	return []string{
		string(OSTypeLinux),
		string(OSTypeWindows),
	}
}

type Protocol string

const (
	ProtocolTCP Protocol = "TCP"
	ProtocolUDP Protocol = "UDP"
)

func PossibleValuesForProtocol() []string {
	return []string{
		string(ProtocolTCP),
		string(ProtocolUDP),
	}
}

type ScaleDownMode string

const (
	ScaleDownModeDeallocate ScaleDownMode = "Deallocate"
	ScaleDownModeDelete     ScaleDownMode = "Delete"
)

func PossibleValuesForScaleDownMode() []string {
	return []string{
		string(ScaleDownModeDeallocate),
		string(ScaleDownModeDelete),
	}
}

type ScaleSetEvictionPolicy string

const (
	ScaleSetEvictionPolicyDeallocate ScaleSetEvictionPolicy = "Deallocate"
	ScaleSetEvictionPolicyDelete     ScaleSetEvictionPolicy = "Delete"
)

func PossibleValuesForScaleSetEvictionPolicy() []string {
	return []string{
		string(ScaleSetEvictionPolicyDeallocate),
		string(ScaleSetEvictionPolicyDelete),
	}
}

type ScaleSetPriority string

const (
	ScaleSetPriorityRegular ScaleSetPriority = "Regular"
	ScaleSetPrioritySpot    ScaleSetPriority = "Spot"
)

func PossibleValuesForScaleSetPriority() []string {
	return []string{
		string(ScaleSetPriorityRegular),
		string(ScaleSetPrioritySpot),
	}
}

type WorkloadRuntime string

const (
	WorkloadRuntimeOCIContainer WorkloadRuntime = "OCIContainer"
	WorkloadRuntimeWasmWasi     WorkloadRuntime = "WasmWasi"
)

func PossibleValuesForWorkloadRuntime() []string {
	return []string{
		string(WorkloadRuntimeOCIContainer),
		string(WorkloadRuntimeWasmWasi),
	}
}
