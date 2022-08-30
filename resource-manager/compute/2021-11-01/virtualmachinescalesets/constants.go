package virtualmachinescalesets

import "strings"

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

func parseCachingTypes(input string) (*CachingTypes, error) {
	vals := map[string]CachingTypes{
		"none":      CachingTypesNone,
		"readonly":  CachingTypesReadOnly,
		"readwrite": CachingTypesReadWrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CachingTypes(input)
	return &out, nil
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

func parseComponentNames(input string) (*ComponentNames, error) {
	vals := map[string]ComponentNames{
		"microsoft-windows-shell-setup": ComponentNamesMicrosoftNegativeWindowsNegativeShellNegativeSetup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComponentNames(input)
	return &out, nil
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

func parseDeleteOptions(input string) (*DeleteOptions, error) {
	vals := map[string]DeleteOptions{
		"delete": DeleteOptionsDelete,
		"detach": DeleteOptionsDetach,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeleteOptions(input)
	return &out, nil
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

func parseDiffDiskOptions(input string) (*DiffDiskOptions, error) {
	vals := map[string]DiffDiskOptions{
		"local": DiffDiskOptionsLocal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiffDiskOptions(input)
	return &out, nil
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

func parseDiffDiskPlacement(input string) (*DiffDiskPlacement, error) {
	vals := map[string]DiffDiskPlacement{
		"cachedisk":    DiffDiskPlacementCacheDisk,
		"resourcedisk": DiffDiskPlacementResourceDisk,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiffDiskPlacement(input)
	return &out, nil
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

func parseDiskCreateOptionTypes(input string) (*DiskCreateOptionTypes, error) {
	vals := map[string]DiskCreateOptionTypes{
		"attach":    DiskCreateOptionTypesAttach,
		"empty":     DiskCreateOptionTypesEmpty,
		"fromimage": DiskCreateOptionTypesFromImage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiskCreateOptionTypes(input)
	return &out, nil
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

func parseExpandTypesForGetVMScaleSets(input string) (*ExpandTypesForGetVMScaleSets, error) {
	vals := map[string]ExpandTypesForGetVMScaleSets{
		"userdata": ExpandTypesForGetVMScaleSetsUserData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExpandTypesForGetVMScaleSets(input)
	return &out, nil
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

func parseIPVersion(input string) (*IPVersion, error) {
	vals := map[string]IPVersion{
		"ipv4": IPVersionIPvFour,
		"ipv6": IPVersionIPvSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPVersion(input)
	return &out, nil
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

func parseLinuxPatchAssessmentMode(input string) (*LinuxPatchAssessmentMode, error) {
	vals := map[string]LinuxPatchAssessmentMode{
		"automaticbyplatform": LinuxPatchAssessmentModeAutomaticByPlatform,
		"imagedefault":        LinuxPatchAssessmentModeImageDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LinuxPatchAssessmentMode(input)
	return &out, nil
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

func parseLinuxVMGuestPatchMode(input string) (*LinuxVMGuestPatchMode, error) {
	vals := map[string]LinuxVMGuestPatchMode{
		"automaticbyplatform": LinuxVMGuestPatchModeAutomaticByPlatform,
		"imagedefault":        LinuxVMGuestPatchModeImageDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LinuxVMGuestPatchMode(input)
	return &out, nil
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

func parseNetworkApiVersion(input string) (*NetworkApiVersion, error) {
	vals := map[string]NetworkApiVersion{
		"2020-11-01": NetworkApiVersionTwoZeroTwoZeroNegativeOneOneNegativeZeroOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkApiVersion(input)
	return &out, nil
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

func parseOperatingSystemTypes(input string) (*OperatingSystemTypes, error) {
	vals := map[string]OperatingSystemTypes{
		"linux":   OperatingSystemTypesLinux,
		"windows": OperatingSystemTypesWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperatingSystemTypes(input)
	return &out, nil
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

func parseOrchestrationMode(input string) (*OrchestrationMode, error) {
	vals := map[string]OrchestrationMode{
		"flexible": OrchestrationModeFlexible,
		"uniform":  OrchestrationModeUniform,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrchestrationMode(input)
	return &out, nil
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

func parseOrchestrationServiceNames(input string) (*OrchestrationServiceNames, error) {
	vals := map[string]OrchestrationServiceNames{
		"automaticrepairs": OrchestrationServiceNamesAutomaticRepairs,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrchestrationServiceNames(input)
	return &out, nil
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

func parseOrchestrationServiceState(input string) (*OrchestrationServiceState, error) {
	vals := map[string]OrchestrationServiceState{
		"notrunning": OrchestrationServiceStateNotRunning,
		"running":    OrchestrationServiceStateRunning,
		"suspended":  OrchestrationServiceStateSuspended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrchestrationServiceState(input)
	return &out, nil
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

func parseOrchestrationServiceStateAction(input string) (*OrchestrationServiceStateAction, error) {
	vals := map[string]OrchestrationServiceStateAction{
		"resume":  OrchestrationServiceStateActionResume,
		"suspend": OrchestrationServiceStateActionSuspend,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrchestrationServiceStateAction(input)
	return &out, nil
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

func parsePassNames(input string) (*PassNames, error) {
	vals := map[string]PassNames{
		"oobesystem": PassNamesOobeSystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PassNames(input)
	return &out, nil
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

func parseProtocolTypes(input string) (*ProtocolTypes, error) {
	vals := map[string]ProtocolTypes{
		"http":  ProtocolTypesHTTP,
		"https": ProtocolTypesHTTPS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtocolTypes(input)
	return &out, nil
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

func parsePublicIPAddressSkuName(input string) (*PublicIPAddressSkuName, error) {
	vals := map[string]PublicIPAddressSkuName{
		"basic":    PublicIPAddressSkuNameBasic,
		"standard": PublicIPAddressSkuNameStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicIPAddressSkuName(input)
	return &out, nil
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

func parsePublicIPAddressSkuTier(input string) (*PublicIPAddressSkuTier, error) {
	vals := map[string]PublicIPAddressSkuTier{
		"global":   PublicIPAddressSkuTierGlobal,
		"regional": PublicIPAddressSkuTierRegional,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicIPAddressSkuTier(input)
	return &out, nil
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

func parseRepairAction(input string) (*RepairAction, error) {
	vals := map[string]RepairAction{
		"reimage": RepairActionReimage,
		"replace": RepairActionReplace,
		"restart": RepairActionRestart,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RepairAction(input)
	return &out, nil
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

func parseSecurityEncryptionTypes(input string) (*SecurityEncryptionTypes, error) {
	vals := map[string]SecurityEncryptionTypes{
		"diskwithvmgueststate": SecurityEncryptionTypesDiskWithVMGuestState,
		"vmgueststateonly":     SecurityEncryptionTypesVMGuestStateOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEncryptionTypes(input)
	return &out, nil
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

func parseSecurityTypes(input string) (*SecurityTypes, error) {
	vals := map[string]SecurityTypes{
		"confidentialvm": SecurityTypesConfidentialVM,
		"trustedlaunch":  SecurityTypesTrustedLaunch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityTypes(input)
	return &out, nil
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

func parseSettingNames(input string) (*SettingNames, error) {
	vals := map[string]SettingNames{
		"autologon":          SettingNamesAutoLogon,
		"firstlogoncommands": SettingNamesFirstLogonCommands,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SettingNames(input)
	return &out, nil
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

func parseStatusLevelTypes(input string) (*StatusLevelTypes, error) {
	vals := map[string]StatusLevelTypes{
		"error":   StatusLevelTypesError,
		"info":    StatusLevelTypesInfo,
		"warning": StatusLevelTypesWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusLevelTypes(input)
	return &out, nil
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

func parseStorageAccountTypes(input string) (*StorageAccountTypes, error) {
	vals := map[string]StorageAccountTypes{
		"premium_lrs":     StorageAccountTypesPremiumLRS,
		"premium_zrs":     StorageAccountTypesPremiumZRS,
		"standard_lrs":    StorageAccountTypesStandardLRS,
		"standardssd_lrs": StorageAccountTypesStandardSSDLRS,
		"standardssd_zrs": StorageAccountTypesStandardSSDZRS,
		"ultrassd_lrs":    StorageAccountTypesUltraSSDLRS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageAccountTypes(input)
	return &out, nil
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

func parseUpgradeMode(input string) (*UpgradeMode, error) {
	vals := map[string]UpgradeMode{
		"automatic": UpgradeModeAutomatic,
		"manual":    UpgradeModeManual,
		"rolling":   UpgradeModeRolling,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpgradeMode(input)
	return &out, nil
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

func parseUpgradeOperationInvoker(input string) (*UpgradeOperationInvoker, error) {
	vals := map[string]UpgradeOperationInvoker{
		"platform": UpgradeOperationInvokerPlatform,
		"unknown":  UpgradeOperationInvokerUnknown,
		"user":     UpgradeOperationInvokerUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpgradeOperationInvoker(input)
	return &out, nil
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

func parseUpgradeState(input string) (*UpgradeState, error) {
	vals := map[string]UpgradeState{
		"cancelled":      UpgradeStateCancelled,
		"completed":      UpgradeStateCompleted,
		"faulted":        UpgradeStateFaulted,
		"rollingforward": UpgradeStateRollingForward,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpgradeState(input)
	return &out, nil
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

func parseVirtualMachineEvictionPolicyTypes(input string) (*VirtualMachineEvictionPolicyTypes, error) {
	vals := map[string]VirtualMachineEvictionPolicyTypes{
		"deallocate": VirtualMachineEvictionPolicyTypesDeallocate,
		"delete":     VirtualMachineEvictionPolicyTypesDelete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineEvictionPolicyTypes(input)
	return &out, nil
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

func parseVirtualMachinePriorityTypes(input string) (*VirtualMachinePriorityTypes, error) {
	vals := map[string]VirtualMachinePriorityTypes{
		"low":     VirtualMachinePriorityTypesLow,
		"regular": VirtualMachinePriorityTypesRegular,
		"spot":    VirtualMachinePriorityTypesSpot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachinePriorityTypes(input)
	return &out, nil
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

func parseVirtualMachineScaleSetScaleInRules(input string) (*VirtualMachineScaleSetScaleInRules, error) {
	vals := map[string]VirtualMachineScaleSetScaleInRules{
		"default":  VirtualMachineScaleSetScaleInRulesDefault,
		"newestvm": VirtualMachineScaleSetScaleInRulesNewestVM,
		"oldestvm": VirtualMachineScaleSetScaleInRulesOldestVM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineScaleSetScaleInRules(input)
	return &out, nil
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

func parseVirtualMachineScaleSetSkuScaleType(input string) (*VirtualMachineScaleSetSkuScaleType, error) {
	vals := map[string]VirtualMachineScaleSetSkuScaleType{
		"automatic": VirtualMachineScaleSetSkuScaleTypeAutomatic,
		"none":      VirtualMachineScaleSetSkuScaleTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineScaleSetSkuScaleType(input)
	return &out, nil
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

func parseWindowsPatchAssessmentMode(input string) (*WindowsPatchAssessmentMode, error) {
	vals := map[string]WindowsPatchAssessmentMode{
		"automaticbyplatform": WindowsPatchAssessmentModeAutomaticByPlatform,
		"imagedefault":        WindowsPatchAssessmentModeImageDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPatchAssessmentMode(input)
	return &out, nil
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

func parseWindowsVMGuestPatchMode(input string) (*WindowsVMGuestPatchMode, error) {
	vals := map[string]WindowsVMGuestPatchMode{
		"automaticbyos":       WindowsVMGuestPatchModeAutomaticByOS,
		"automaticbyplatform": WindowsVMGuestPatchModeAutomaticByPlatform,
		"manual":              WindowsVMGuestPatchModeManual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsVMGuestPatchMode(input)
	return &out, nil
}
