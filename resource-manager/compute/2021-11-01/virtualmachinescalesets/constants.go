package virtualmachinescalesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CachingTypes string

const (
	CachingTypesNone      CachingTypes = "None"
	CachingTypesReadOnly  CachingTypes = "ReadOnly"
	CachingTypesReadWrite CachingTypes = "ReadWrite"
)

func PossibleValuesForCachingTypes() []string {
	return []string{
		string(CachingTypesNone),
		string(CachingTypesReadOnly),
		string(CachingTypesReadWrite),
	}
}

type ComponentNames string

const (
	ComponentNamesMicrosoftNegativeWindowsNegativeShellNegativeSetup ComponentNames = "Microsoft-Windows-Shell-Setup"
)

func PossibleValuesForComponentNames() []string {
	return []string{
		string(ComponentNamesMicrosoftNegativeWindowsNegativeShellNegativeSetup),
	}
}

type DeleteOptions string

const (
	DeleteOptionsDelete DeleteOptions = "Delete"
	DeleteOptionsDetach DeleteOptions = "Detach"
)

func PossibleValuesForDeleteOptions() []string {
	return []string{
		string(DeleteOptionsDelete),
		string(DeleteOptionsDetach),
	}
}

type DiffDiskOptions string

const (
	DiffDiskOptionsLocal DiffDiskOptions = "Local"
)

func PossibleValuesForDiffDiskOptions() []string {
	return []string{
		string(DiffDiskOptionsLocal),
	}
}

type DiffDiskPlacement string

const (
	DiffDiskPlacementCacheDisk    DiffDiskPlacement = "CacheDisk"
	DiffDiskPlacementResourceDisk DiffDiskPlacement = "ResourceDisk"
)

func PossibleValuesForDiffDiskPlacement() []string {
	return []string{
		string(DiffDiskPlacementCacheDisk),
		string(DiffDiskPlacementResourceDisk),
	}
}

type DiskCreateOptionTypes string

const (
	DiskCreateOptionTypesAttach    DiskCreateOptionTypes = "Attach"
	DiskCreateOptionTypesEmpty     DiskCreateOptionTypes = "Empty"
	DiskCreateOptionTypesFromImage DiskCreateOptionTypes = "FromImage"
)

func PossibleValuesForDiskCreateOptionTypes() []string {
	return []string{
		string(DiskCreateOptionTypesAttach),
		string(DiskCreateOptionTypesEmpty),
		string(DiskCreateOptionTypesFromImage),
	}
}

type ExpandTypesForGetVMScaleSets string

const (
	ExpandTypesForGetVMScaleSetsUserData ExpandTypesForGetVMScaleSets = "userData"
)

func PossibleValuesForExpandTypesForGetVMScaleSets() []string {
	return []string{
		string(ExpandTypesForGetVMScaleSetsUserData),
	}
}

type IPVersion string

const (
	IPVersionIPvFour IPVersion = "IPv4"
	IPVersionIPvSix  IPVersion = "IPv6"
)

func PossibleValuesForIPVersion() []string {
	return []string{
		string(IPVersionIPvFour),
		string(IPVersionIPvSix),
	}
}

type LinuxPatchAssessmentMode string

const (
	LinuxPatchAssessmentModeAutomaticByPlatform LinuxPatchAssessmentMode = "AutomaticByPlatform"
	LinuxPatchAssessmentModeImageDefault        LinuxPatchAssessmentMode = "ImageDefault"
)

func PossibleValuesForLinuxPatchAssessmentMode() []string {
	return []string{
		string(LinuxPatchAssessmentModeAutomaticByPlatform),
		string(LinuxPatchAssessmentModeImageDefault),
	}
}

type LinuxVMGuestPatchMode string

const (
	LinuxVMGuestPatchModeAutomaticByPlatform LinuxVMGuestPatchMode = "AutomaticByPlatform"
	LinuxVMGuestPatchModeImageDefault        LinuxVMGuestPatchMode = "ImageDefault"
)

func PossibleValuesForLinuxVMGuestPatchMode() []string {
	return []string{
		string(LinuxVMGuestPatchModeAutomaticByPlatform),
		string(LinuxVMGuestPatchModeImageDefault),
	}
}

type NetworkApiVersion string

const (
	NetworkApiVersionTwoZeroTwoZeroNegativeOneOneNegativeZeroOne NetworkApiVersion = "2020-11-01"
)

func PossibleValuesForNetworkApiVersion() []string {
	return []string{
		string(NetworkApiVersionTwoZeroTwoZeroNegativeOneOneNegativeZeroOne),
	}
}

type OperatingSystemTypes string

const (
	OperatingSystemTypesLinux   OperatingSystemTypes = "Linux"
	OperatingSystemTypesWindows OperatingSystemTypes = "Windows"
)

func PossibleValuesForOperatingSystemTypes() []string {
	return []string{
		string(OperatingSystemTypesLinux),
		string(OperatingSystemTypesWindows),
	}
}

type OrchestrationMode string

const (
	OrchestrationModeFlexible OrchestrationMode = "Flexible"
	OrchestrationModeUniform  OrchestrationMode = "Uniform"
)

func PossibleValuesForOrchestrationMode() []string {
	return []string{
		string(OrchestrationModeFlexible),
		string(OrchestrationModeUniform),
	}
}

type OrchestrationServiceNames string

const (
	OrchestrationServiceNamesAutomaticRepairs OrchestrationServiceNames = "AutomaticRepairs"
)

func PossibleValuesForOrchestrationServiceNames() []string {
	return []string{
		string(OrchestrationServiceNamesAutomaticRepairs),
	}
}

type OrchestrationServiceState string

const (
	OrchestrationServiceStateNotRunning OrchestrationServiceState = "NotRunning"
	OrchestrationServiceStateRunning    OrchestrationServiceState = "Running"
	OrchestrationServiceStateSuspended  OrchestrationServiceState = "Suspended"
)

func PossibleValuesForOrchestrationServiceState() []string {
	return []string{
		string(OrchestrationServiceStateNotRunning),
		string(OrchestrationServiceStateRunning),
		string(OrchestrationServiceStateSuspended),
	}
}

type OrchestrationServiceStateAction string

const (
	OrchestrationServiceStateActionResume  OrchestrationServiceStateAction = "Resume"
	OrchestrationServiceStateActionSuspend OrchestrationServiceStateAction = "Suspend"
)

func PossibleValuesForOrchestrationServiceStateAction() []string {
	return []string{
		string(OrchestrationServiceStateActionResume),
		string(OrchestrationServiceStateActionSuspend),
	}
}

type PassNames string

const (
	PassNamesOobeSystem PassNames = "OobeSystem"
)

func PossibleValuesForPassNames() []string {
	return []string{
		string(PassNamesOobeSystem),
	}
}

type ProtocolTypes string

const (
	ProtocolTypesHTTP  ProtocolTypes = "Http"
	ProtocolTypesHTTPS ProtocolTypes = "Https"
)

func PossibleValuesForProtocolTypes() []string {
	return []string{
		string(ProtocolTypesHTTP),
		string(ProtocolTypesHTTPS),
	}
}

type PublicIPAddressSkuName string

const (
	PublicIPAddressSkuNameBasic    PublicIPAddressSkuName = "Basic"
	PublicIPAddressSkuNameStandard PublicIPAddressSkuName = "Standard"
)

func PossibleValuesForPublicIPAddressSkuName() []string {
	return []string{
		string(PublicIPAddressSkuNameBasic),
		string(PublicIPAddressSkuNameStandard),
	}
}

type PublicIPAddressSkuTier string

const (
	PublicIPAddressSkuTierGlobal   PublicIPAddressSkuTier = "Global"
	PublicIPAddressSkuTierRegional PublicIPAddressSkuTier = "Regional"
)

func PossibleValuesForPublicIPAddressSkuTier() []string {
	return []string{
		string(PublicIPAddressSkuTierGlobal),
		string(PublicIPAddressSkuTierRegional),
	}
}

type RepairAction string

const (
	RepairActionReimage RepairAction = "Reimage"
	RepairActionReplace RepairAction = "Replace"
	RepairActionRestart RepairAction = "Restart"
)

func PossibleValuesForRepairAction() []string {
	return []string{
		string(RepairActionReimage),
		string(RepairActionReplace),
		string(RepairActionRestart),
	}
}

type SecurityEncryptionTypes string

const (
	SecurityEncryptionTypesDiskWithVMGuestState SecurityEncryptionTypes = "DiskWithVMGuestState"
	SecurityEncryptionTypesVMGuestStateOnly     SecurityEncryptionTypes = "VMGuestStateOnly"
)

func PossibleValuesForSecurityEncryptionTypes() []string {
	return []string{
		string(SecurityEncryptionTypesDiskWithVMGuestState),
		string(SecurityEncryptionTypesVMGuestStateOnly),
	}
}

type SecurityTypes string

const (
	SecurityTypesConfidentialVM SecurityTypes = "ConfidentialVM"
	SecurityTypesTrustedLaunch  SecurityTypes = "TrustedLaunch"
)

func PossibleValuesForSecurityTypes() []string {
	return []string{
		string(SecurityTypesConfidentialVM),
		string(SecurityTypesTrustedLaunch),
	}
}

type SettingNames string

const (
	SettingNamesAutoLogon          SettingNames = "AutoLogon"
	SettingNamesFirstLogonCommands SettingNames = "FirstLogonCommands"
)

func PossibleValuesForSettingNames() []string {
	return []string{
		string(SettingNamesAutoLogon),
		string(SettingNamesFirstLogonCommands),
	}
}

type StatusLevelTypes string

const (
	StatusLevelTypesError   StatusLevelTypes = "Error"
	StatusLevelTypesInfo    StatusLevelTypes = "Info"
	StatusLevelTypesWarning StatusLevelTypes = "Warning"
)

func PossibleValuesForStatusLevelTypes() []string {
	return []string{
		string(StatusLevelTypesError),
		string(StatusLevelTypesInfo),
		string(StatusLevelTypesWarning),
	}
}

type StorageAccountTypes string

const (
	StorageAccountTypesPremiumLRS     StorageAccountTypes = "Premium_LRS"
	StorageAccountTypesPremiumZRS     StorageAccountTypes = "Premium_ZRS"
	StorageAccountTypesStandardLRS    StorageAccountTypes = "Standard_LRS"
	StorageAccountTypesStandardSSDLRS StorageAccountTypes = "StandardSSD_LRS"
	StorageAccountTypesStandardSSDZRS StorageAccountTypes = "StandardSSD_ZRS"
	StorageAccountTypesUltraSSDLRS    StorageAccountTypes = "UltraSSD_LRS"
)

func PossibleValuesForStorageAccountTypes() []string {
	return []string{
		string(StorageAccountTypesPremiumLRS),
		string(StorageAccountTypesPremiumZRS),
		string(StorageAccountTypesStandardLRS),
		string(StorageAccountTypesStandardSSDLRS),
		string(StorageAccountTypesStandardSSDZRS),
		string(StorageAccountTypesUltraSSDLRS),
	}
}

type UpgradeMode string

const (
	UpgradeModeAutomatic UpgradeMode = "Automatic"
	UpgradeModeManual    UpgradeMode = "Manual"
	UpgradeModeRolling   UpgradeMode = "Rolling"
)

func PossibleValuesForUpgradeMode() []string {
	return []string{
		string(UpgradeModeAutomatic),
		string(UpgradeModeManual),
		string(UpgradeModeRolling),
	}
}

type UpgradeOperationInvoker string

const (
	UpgradeOperationInvokerPlatform UpgradeOperationInvoker = "Platform"
	UpgradeOperationInvokerUnknown  UpgradeOperationInvoker = "Unknown"
	UpgradeOperationInvokerUser     UpgradeOperationInvoker = "User"
)

func PossibleValuesForUpgradeOperationInvoker() []string {
	return []string{
		string(UpgradeOperationInvokerPlatform),
		string(UpgradeOperationInvokerUnknown),
		string(UpgradeOperationInvokerUser),
	}
}

type UpgradeState string

const (
	UpgradeStateCancelled      UpgradeState = "Cancelled"
	UpgradeStateCompleted      UpgradeState = "Completed"
	UpgradeStateFaulted        UpgradeState = "Faulted"
	UpgradeStateRollingForward UpgradeState = "RollingForward"
)

func PossibleValuesForUpgradeState() []string {
	return []string{
		string(UpgradeStateCancelled),
		string(UpgradeStateCompleted),
		string(UpgradeStateFaulted),
		string(UpgradeStateRollingForward),
	}
}

type VirtualMachineEvictionPolicyTypes string

const (
	VirtualMachineEvictionPolicyTypesDeallocate VirtualMachineEvictionPolicyTypes = "Deallocate"
	VirtualMachineEvictionPolicyTypesDelete     VirtualMachineEvictionPolicyTypes = "Delete"
)

func PossibleValuesForVirtualMachineEvictionPolicyTypes() []string {
	return []string{
		string(VirtualMachineEvictionPolicyTypesDeallocate),
		string(VirtualMachineEvictionPolicyTypesDelete),
	}
}

type VirtualMachinePriorityTypes string

const (
	VirtualMachinePriorityTypesLow     VirtualMachinePriorityTypes = "Low"
	VirtualMachinePriorityTypesRegular VirtualMachinePriorityTypes = "Regular"
	VirtualMachinePriorityTypesSpot    VirtualMachinePriorityTypes = "Spot"
)

func PossibleValuesForVirtualMachinePriorityTypes() []string {
	return []string{
		string(VirtualMachinePriorityTypesLow),
		string(VirtualMachinePriorityTypesRegular),
		string(VirtualMachinePriorityTypesSpot),
	}
}

type VirtualMachineScaleSetScaleInRules string

const (
	VirtualMachineScaleSetScaleInRulesDefault  VirtualMachineScaleSetScaleInRules = "Default"
	VirtualMachineScaleSetScaleInRulesNewestVM VirtualMachineScaleSetScaleInRules = "NewestVM"
	VirtualMachineScaleSetScaleInRulesOldestVM VirtualMachineScaleSetScaleInRules = "OldestVM"
)

func PossibleValuesForVirtualMachineScaleSetScaleInRules() []string {
	return []string{
		string(VirtualMachineScaleSetScaleInRulesDefault),
		string(VirtualMachineScaleSetScaleInRulesNewestVM),
		string(VirtualMachineScaleSetScaleInRulesOldestVM),
	}
}

type VirtualMachineScaleSetSkuScaleType string

const (
	VirtualMachineScaleSetSkuScaleTypeAutomatic VirtualMachineScaleSetSkuScaleType = "Automatic"
	VirtualMachineScaleSetSkuScaleTypeNone      VirtualMachineScaleSetSkuScaleType = "None"
)

func PossibleValuesForVirtualMachineScaleSetSkuScaleType() []string {
	return []string{
		string(VirtualMachineScaleSetSkuScaleTypeAutomatic),
		string(VirtualMachineScaleSetSkuScaleTypeNone),
	}
}

type WindowsPatchAssessmentMode string

const (
	WindowsPatchAssessmentModeAutomaticByPlatform WindowsPatchAssessmentMode = "AutomaticByPlatform"
	WindowsPatchAssessmentModeImageDefault        WindowsPatchAssessmentMode = "ImageDefault"
)

func PossibleValuesForWindowsPatchAssessmentMode() []string {
	return []string{
		string(WindowsPatchAssessmentModeAutomaticByPlatform),
		string(WindowsPatchAssessmentModeImageDefault),
	}
}

type WindowsVMGuestPatchMode string

const (
	WindowsVMGuestPatchModeAutomaticByOS       WindowsVMGuestPatchMode = "AutomaticByOS"
	WindowsVMGuestPatchModeAutomaticByPlatform WindowsVMGuestPatchMode = "AutomaticByPlatform"
	WindowsVMGuestPatchModeManual              WindowsVMGuestPatchMode = "Manual"
)

func PossibleValuesForWindowsVMGuestPatchMode() []string {
	return []string{
		string(WindowsVMGuestPatchModeAutomaticByOS),
		string(WindowsVMGuestPatchModeAutomaticByPlatform),
		string(WindowsVMGuestPatchModeManual),
	}
}
