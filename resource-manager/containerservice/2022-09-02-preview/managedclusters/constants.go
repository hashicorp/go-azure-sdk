package managedclusters

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

type BackendPoolType string

const (
	BackendPoolTypeNodeIP              BackendPoolType = "NodeIP"
	BackendPoolTypeNodeIPConfiguration BackendPoolType = "NodeIPConfiguration"
)

func PossibleValuesForBackendPoolType() []string {
	return []string{
		string(BackendPoolTypeNodeIP),
		string(BackendPoolTypeNodeIPConfiguration),
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

type ControlledValues string

const (
	ControlledValuesRequestsAndLimits ControlledValues = "RequestsAndLimits"
	ControlledValuesRequestsOnly      ControlledValues = "RequestsOnly"
)

func PossibleValuesForControlledValues() []string {
	return []string{
		string(ControlledValuesRequestsAndLimits),
		string(ControlledValuesRequestsOnly),
	}
}

type EbpfDataplane string

const (
	EbpfDataplaneCilium EbpfDataplane = "cilium"
)

func PossibleValuesForEbpfDataplane() []string {
	return []string{
		string(EbpfDataplaneCilium),
	}
}

type Expander string

const (
	ExpanderLeastNegativewaste Expander = "least-waste"
	ExpanderMostNegativepods   Expander = "most-pods"
	ExpanderPriority           Expander = "priority"
	ExpanderRandom             Expander = "random"
)

func PossibleValuesForExpander() []string {
	return []string{
		string(ExpanderLeastNegativewaste),
		string(ExpanderMostNegativepods),
		string(ExpanderPriority),
		string(ExpanderRandom),
	}
}

type Format string

const (
	FormatAzure Format = "azure"
	FormatExec  Format = "exec"
)

func PossibleValuesForFormat() []string {
	return []string{
		string(FormatAzure),
		string(FormatExec),
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

type IPFamily string

const (
	IPFamilyIPvFour IPFamily = "IPv4"
	IPFamilyIPvSix  IPFamily = "IPv6"
)

func PossibleValuesForIPFamily() []string {
	return []string{
		string(IPFamilyIPvFour),
		string(IPFamilyIPvSix),
	}
}

type IPvsScheduler string

const (
	IPvsSchedulerLeastConnection IPvsScheduler = "LeastConnection"
	IPvsSchedulerRoundRobin      IPvsScheduler = "RoundRobin"
)

func PossibleValuesForIPvsScheduler() []string {
	return []string{
		string(IPvsSchedulerLeastConnection),
		string(IPvsSchedulerRoundRobin),
	}
}

type KeyVaultNetworkAccessTypes string

const (
	KeyVaultNetworkAccessTypesPrivate KeyVaultNetworkAccessTypes = "Private"
	KeyVaultNetworkAccessTypesPublic  KeyVaultNetworkAccessTypes = "Public"
)

func PossibleValuesForKeyVaultNetworkAccessTypes() []string {
	return []string{
		string(KeyVaultNetworkAccessTypesPrivate),
		string(KeyVaultNetworkAccessTypesPublic),
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

type Level string

const (
	LevelEnforcement Level = "Enforcement"
	LevelOff         Level = "Off"
	LevelWarning     Level = "Warning"
)

func PossibleValuesForLevel() []string {
	return []string{
		string(LevelEnforcement),
		string(LevelOff),
		string(LevelWarning),
	}
}

type LicenseType string

const (
	LicenseTypeNone          LicenseType = "None"
	LicenseTypeWindowsServer LicenseType = "Windows_Server"
)

func PossibleValuesForLicenseType() []string {
	return []string{
		string(LicenseTypeNone),
		string(LicenseTypeWindowsServer),
	}
}

type LoadBalancerSku string

const (
	LoadBalancerSkuBasic    LoadBalancerSku = "basic"
	LoadBalancerSkuStandard LoadBalancerSku = "standard"
)

func PossibleValuesForLoadBalancerSku() []string {
	return []string{
		string(LoadBalancerSkuBasic),
		string(LoadBalancerSkuStandard),
	}
}

type ManagedClusterPodIdentityProvisioningState string

const (
	ManagedClusterPodIdentityProvisioningStateAssigned ManagedClusterPodIdentityProvisioningState = "Assigned"
	ManagedClusterPodIdentityProvisioningStateDeleting ManagedClusterPodIdentityProvisioningState = "Deleting"
	ManagedClusterPodIdentityProvisioningStateFailed   ManagedClusterPodIdentityProvisioningState = "Failed"
	ManagedClusterPodIdentityProvisioningStateUpdating ManagedClusterPodIdentityProvisioningState = "Updating"
)

func PossibleValuesForManagedClusterPodIdentityProvisioningState() []string {
	return []string{
		string(ManagedClusterPodIdentityProvisioningStateAssigned),
		string(ManagedClusterPodIdentityProvisioningStateDeleting),
		string(ManagedClusterPodIdentityProvisioningStateFailed),
		string(ManagedClusterPodIdentityProvisioningStateUpdating),
	}
}

type ManagedClusterSKUName string

const (
	ManagedClusterSKUNameBasic ManagedClusterSKUName = "Basic"
)

func PossibleValuesForManagedClusterSKUName() []string {
	return []string{
		string(ManagedClusterSKUNameBasic),
	}
}

type ManagedClusterSKUTier string

const (
	ManagedClusterSKUTierFree ManagedClusterSKUTier = "Free"
	ManagedClusterSKUTierPaid ManagedClusterSKUTier = "Paid"
)

func PossibleValuesForManagedClusterSKUTier() []string {
	return []string{
		string(ManagedClusterSKUTierFree),
		string(ManagedClusterSKUTierPaid),
	}
}

type Mode string

const (
	ModeIPTABLES Mode = "IPTABLES"
	ModeIPVS     Mode = "IPVS"
)

func PossibleValuesForMode() []string {
	return []string{
		string(ModeIPTABLES),
		string(ModeIPVS),
	}
}

type NetworkMode string

const (
	NetworkModeBridge      NetworkMode = "bridge"
	NetworkModeTransparent NetworkMode = "transparent"
)

func PossibleValuesForNetworkMode() []string {
	return []string{
		string(NetworkModeBridge),
		string(NetworkModeTransparent),
	}
}

type NetworkPlugin string

const (
	NetworkPluginAzure   NetworkPlugin = "azure"
	NetworkPluginKubenet NetworkPlugin = "kubenet"
	NetworkPluginNone    NetworkPlugin = "none"
)

func PossibleValuesForNetworkPlugin() []string {
	return []string{
		string(NetworkPluginAzure),
		string(NetworkPluginKubenet),
		string(NetworkPluginNone),
	}
}

type NetworkPluginMode string

const (
	NetworkPluginModeOverlay NetworkPluginMode = "Overlay"
)

func PossibleValuesForNetworkPluginMode() []string {
	return []string{
		string(NetworkPluginModeOverlay),
	}
}

type NetworkPolicy string

const (
	NetworkPolicyAzure  NetworkPolicy = "azure"
	NetworkPolicyCalico NetworkPolicy = "calico"
)

func PossibleValuesForNetworkPolicy() []string {
	return []string{
		string(NetworkPolicyAzure),
		string(NetworkPolicyCalico),
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

type OutboundType string

const (
	OutboundTypeLoadBalancer           OutboundType = "loadBalancer"
	OutboundTypeManagedNATGateway      OutboundType = "managedNATGateway"
	OutboundTypeUserAssignedNATGateway OutboundType = "userAssignedNATGateway"
	OutboundTypeUserDefinedRouting     OutboundType = "userDefinedRouting"
)

func PossibleValuesForOutboundType() []string {
	return []string{
		string(OutboundTypeLoadBalancer),
		string(OutboundTypeManagedNATGateway),
		string(OutboundTypeUserAssignedNATGateway),
		string(OutboundTypeUserDefinedRouting),
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

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled           PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled            PublicNetworkAccess = "Enabled"
	PublicNetworkAccessSecuredByPerimeter PublicNetworkAccess = "SecuredByPerimeter"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
		string(PublicNetworkAccessSecuredByPerimeter),
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

type UpdateMode string

const (
	UpdateModeAuto     UpdateMode = "Auto"
	UpdateModeInitial  UpdateMode = "Initial"
	UpdateModeOff      UpdateMode = "Off"
	UpdateModeRecreate UpdateMode = "Recreate"
)

func PossibleValuesForUpdateMode() []string {
	return []string{
		string(UpdateModeAuto),
		string(UpdateModeInitial),
		string(UpdateModeOff),
		string(UpdateModeRecreate),
	}
}

type UpgradeChannel string

const (
	UpgradeChannelNodeNegativeimage UpgradeChannel = "node-image"
	UpgradeChannelNone              UpgradeChannel = "none"
	UpgradeChannelPatch             UpgradeChannel = "patch"
	UpgradeChannelRapid             UpgradeChannel = "rapid"
	UpgradeChannelStable            UpgradeChannel = "stable"
)

func PossibleValuesForUpgradeChannel() []string {
	return []string{
		string(UpgradeChannelNodeNegativeimage),
		string(UpgradeChannelNone),
		string(UpgradeChannelPatch),
		string(UpgradeChannelRapid),
		string(UpgradeChannelStable),
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
