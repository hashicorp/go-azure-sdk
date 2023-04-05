package disks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessLevel string

const (
	AccessLevelNone  AccessLevel = "None"
	AccessLevelRead  AccessLevel = "Read"
	AccessLevelWrite AccessLevel = "Write"
)

func PossibleValuesForAccessLevel() []string {
	return []string{
		string(AccessLevelNone),
		string(AccessLevelRead),
		string(AccessLevelWrite),
	}
}

type Architecture string

const (
	ArchitectureArmSixFour Architecture = "Arm64"
	ArchitectureXSixFour   Architecture = "x64"
)

func PossibleValuesForArchitecture() []string {
	return []string{
		string(ArchitectureArmSixFour),
		string(ArchitectureXSixFour),
	}
}

type DataAccessAuthMode string

const (
	DataAccessAuthModeAzureActiveDirectory DataAccessAuthMode = "AzureActiveDirectory"
	DataAccessAuthModeNone                 DataAccessAuthMode = "None"
)

func PossibleValuesForDataAccessAuthMode() []string {
	return []string{
		string(DataAccessAuthModeAzureActiveDirectory),
		string(DataAccessAuthModeNone),
	}
}

type DiskCreateOption string

const (
	DiskCreateOptionAttach               DiskCreateOption = "Attach"
	DiskCreateOptionCopy                 DiskCreateOption = "Copy"
	DiskCreateOptionCopyStart            DiskCreateOption = "CopyStart"
	DiskCreateOptionEmpty                DiskCreateOption = "Empty"
	DiskCreateOptionFromImage            DiskCreateOption = "FromImage"
	DiskCreateOptionImport               DiskCreateOption = "Import"
	DiskCreateOptionImportSecure         DiskCreateOption = "ImportSecure"
	DiskCreateOptionRestore              DiskCreateOption = "Restore"
	DiskCreateOptionUpload               DiskCreateOption = "Upload"
	DiskCreateOptionUploadPreparedSecure DiskCreateOption = "UploadPreparedSecure"
)

func PossibleValuesForDiskCreateOption() []string {
	return []string{
		string(DiskCreateOptionAttach),
		string(DiskCreateOptionCopy),
		string(DiskCreateOptionCopyStart),
		string(DiskCreateOptionEmpty),
		string(DiskCreateOptionFromImage),
		string(DiskCreateOptionImport),
		string(DiskCreateOptionImportSecure),
		string(DiskCreateOptionRestore),
		string(DiskCreateOptionUpload),
		string(DiskCreateOptionUploadPreparedSecure),
	}
}

type DiskSecurityTypes string

const (
	DiskSecurityTypesConfidentialVMDiskEncryptedWithCustomerKey             DiskSecurityTypes = "ConfidentialVM_DiskEncryptedWithCustomerKey"
	DiskSecurityTypesConfidentialVMDiskEncryptedWithPlatformKey             DiskSecurityTypes = "ConfidentialVM_DiskEncryptedWithPlatformKey"
	DiskSecurityTypesConfidentialVMVMGuestStateOnlyEncryptedWithPlatformKey DiskSecurityTypes = "ConfidentialVM_VMGuestStateOnlyEncryptedWithPlatformKey"
	DiskSecurityTypesTrustedLaunch                                          DiskSecurityTypes = "TrustedLaunch"
)

func PossibleValuesForDiskSecurityTypes() []string {
	return []string{
		string(DiskSecurityTypesConfidentialVMDiskEncryptedWithCustomerKey),
		string(DiskSecurityTypesConfidentialVMDiskEncryptedWithPlatformKey),
		string(DiskSecurityTypesConfidentialVMVMGuestStateOnlyEncryptedWithPlatformKey),
		string(DiskSecurityTypesTrustedLaunch),
	}
}

type DiskState string

const (
	DiskStateActiveSAS       DiskState = "ActiveSAS"
	DiskStateActiveSASFrozen DiskState = "ActiveSASFrozen"
	DiskStateActiveUpload    DiskState = "ActiveUpload"
	DiskStateAttached        DiskState = "Attached"
	DiskStateFrozen          DiskState = "Frozen"
	DiskStateReadyToUpload   DiskState = "ReadyToUpload"
	DiskStateReserved        DiskState = "Reserved"
	DiskStateUnattached      DiskState = "Unattached"
)

func PossibleValuesForDiskState() []string {
	return []string{
		string(DiskStateActiveSAS),
		string(DiskStateActiveSASFrozen),
		string(DiskStateActiveUpload),
		string(DiskStateAttached),
		string(DiskStateFrozen),
		string(DiskStateReadyToUpload),
		string(DiskStateReserved),
		string(DiskStateUnattached),
	}
}

type DiskStorageAccountTypes string

const (
	DiskStorageAccountTypesPremiumLRS     DiskStorageAccountTypes = "Premium_LRS"
	DiskStorageAccountTypesPremiumVTwoLRS DiskStorageAccountTypes = "PremiumV2_LRS"
	DiskStorageAccountTypesPremiumZRS     DiskStorageAccountTypes = "Premium_ZRS"
	DiskStorageAccountTypesStandardLRS    DiskStorageAccountTypes = "Standard_LRS"
	DiskStorageAccountTypesStandardSSDLRS DiskStorageAccountTypes = "StandardSSD_LRS"
	DiskStorageAccountTypesStandardSSDZRS DiskStorageAccountTypes = "StandardSSD_ZRS"
	DiskStorageAccountTypesUltraSSDLRS    DiskStorageAccountTypes = "UltraSSD_LRS"
)

func PossibleValuesForDiskStorageAccountTypes() []string {
	return []string{
		string(DiskStorageAccountTypesPremiumLRS),
		string(DiskStorageAccountTypesPremiumVTwoLRS),
		string(DiskStorageAccountTypesPremiumZRS),
		string(DiskStorageAccountTypesStandardLRS),
		string(DiskStorageAccountTypesStandardSSDLRS),
		string(DiskStorageAccountTypesStandardSSDZRS),
		string(DiskStorageAccountTypesUltraSSDLRS),
	}
}

type EncryptionType string

const (
	EncryptionTypeEncryptionAtRestWithCustomerKey             EncryptionType = "EncryptionAtRestWithCustomerKey"
	EncryptionTypeEncryptionAtRestWithPlatformAndCustomerKeys EncryptionType = "EncryptionAtRestWithPlatformAndCustomerKeys"
	EncryptionTypeEncryptionAtRestWithPlatformKey             EncryptionType = "EncryptionAtRestWithPlatformKey"
)

func PossibleValuesForEncryptionType() []string {
	return []string{
		string(EncryptionTypeEncryptionAtRestWithCustomerKey),
		string(EncryptionTypeEncryptionAtRestWithPlatformAndCustomerKeys),
		string(EncryptionTypeEncryptionAtRestWithPlatformKey),
	}
}

type HyperVGeneration string

const (
	HyperVGenerationVOne HyperVGeneration = "V1"
	HyperVGenerationVTwo HyperVGeneration = "V2"
)

func PossibleValuesForHyperVGeneration() []string {
	return []string{
		string(HyperVGenerationVOne),
		string(HyperVGenerationVTwo),
	}
}

type NetworkAccessPolicy string

const (
	NetworkAccessPolicyAllowAll     NetworkAccessPolicy = "AllowAll"
	NetworkAccessPolicyAllowPrivate NetworkAccessPolicy = "AllowPrivate"
	NetworkAccessPolicyDenyAll      NetworkAccessPolicy = "DenyAll"
)

func PossibleValuesForNetworkAccessPolicy() []string {
	return []string{
		string(NetworkAccessPolicyAllowAll),
		string(NetworkAccessPolicyAllowPrivate),
		string(NetworkAccessPolicyDenyAll),
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

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}
