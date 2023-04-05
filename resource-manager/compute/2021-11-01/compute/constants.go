package compute

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

type ConsistencyModeTypes string

const (
	ConsistencyModeTypesApplicationConsistent ConsistencyModeTypes = "ApplicationConsistent"
	ConsistencyModeTypesCrashConsistent       ConsistencyModeTypes = "CrashConsistent"
	ConsistencyModeTypesFileSystemConsistent  ConsistencyModeTypes = "FileSystemConsistent"
)

func PossibleValuesForConsistencyModeTypes() []string {
	return []string{
		string(ConsistencyModeTypesApplicationConsistent),
		string(ConsistencyModeTypesCrashConsistent),
		string(ConsistencyModeTypesFileSystemConsistent),
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

type OperatingSystemType string

const (
	OperatingSystemTypeLinux   OperatingSystemType = "Linux"
	OperatingSystemTypeWindows OperatingSystemType = "Windows"
)

func PossibleValuesForOperatingSystemType() []string {
	return []string{
		string(OperatingSystemTypeLinux),
		string(OperatingSystemTypeWindows),
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

type RestorePointExpandOptions string

const (
	RestorePointExpandOptionsInstanceView RestorePointExpandOptions = "instanceView"
)

func PossibleValuesForRestorePointExpandOptions() []string {
	return []string{
		string(RestorePointExpandOptionsInstanceView),
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

type VirtualMachineSizeTypes string

const (
	VirtualMachineSizeTypesBasicAFour                              VirtualMachineSizeTypes = "Basic_A4"
	VirtualMachineSizeTypesBasicAOne                               VirtualMachineSizeTypes = "Basic_A1"
	VirtualMachineSizeTypesBasicAThree                             VirtualMachineSizeTypes = "Basic_A3"
	VirtualMachineSizeTypesBasicATwo                               VirtualMachineSizeTypes = "Basic_A2"
	VirtualMachineSizeTypesBasicAZero                              VirtualMachineSizeTypes = "Basic_A0"
	VirtualMachineSizeTypesStandardAEight                          VirtualMachineSizeTypes = "Standard_A8"
	VirtualMachineSizeTypesStandardAEightVTwo                      VirtualMachineSizeTypes = "Standard_A8_v2"
	VirtualMachineSizeTypesStandardAEightmVTwo                     VirtualMachineSizeTypes = "Standard_A8m_v2"
	VirtualMachineSizeTypesStandardAFive                           VirtualMachineSizeTypes = "Standard_A5"
	VirtualMachineSizeTypesStandardAFour                           VirtualMachineSizeTypes = "Standard_A4"
	VirtualMachineSizeTypesStandardAFourVTwo                       VirtualMachineSizeTypes = "Standard_A4_v2"
	VirtualMachineSizeTypesStandardAFourmVTwo                      VirtualMachineSizeTypes = "Standard_A4m_v2"
	VirtualMachineSizeTypesStandardANine                           VirtualMachineSizeTypes = "Standard_A9"
	VirtualMachineSizeTypesStandardAOne                            VirtualMachineSizeTypes = "Standard_A1"
	VirtualMachineSizeTypesStandardAOneOne                         VirtualMachineSizeTypes = "Standard_A11"
	VirtualMachineSizeTypesStandardAOneVTwo                        VirtualMachineSizeTypes = "Standard_A1_v2"
	VirtualMachineSizeTypesStandardAOneZero                        VirtualMachineSizeTypes = "Standard_A10"
	VirtualMachineSizeTypesStandardASeven                          VirtualMachineSizeTypes = "Standard_A7"
	VirtualMachineSizeTypesStandardASix                            VirtualMachineSizeTypes = "Standard_A6"
	VirtualMachineSizeTypesStandardAThree                          VirtualMachineSizeTypes = "Standard_A3"
	VirtualMachineSizeTypesStandardATwo                            VirtualMachineSizeTypes = "Standard_A2"
	VirtualMachineSizeTypesStandardATwoVTwo                        VirtualMachineSizeTypes = "Standard_A2_v2"
	VirtualMachineSizeTypesStandardATwomVTwo                       VirtualMachineSizeTypes = "Standard_A2m_v2"
	VirtualMachineSizeTypesStandardAZero                           VirtualMachineSizeTypes = "Standard_A0"
	VirtualMachineSizeTypesStandardBEightms                        VirtualMachineSizeTypes = "Standard_B8ms"
	VirtualMachineSizeTypesStandardBFourms                         VirtualMachineSizeTypes = "Standard_B4ms"
	VirtualMachineSizeTypesStandardBOnems                          VirtualMachineSizeTypes = "Standard_B1ms"
	VirtualMachineSizeTypesStandardBOnes                           VirtualMachineSizeTypes = "Standard_B1s"
	VirtualMachineSizeTypesStandardBTwoms                          VirtualMachineSizeTypes = "Standard_B2ms"
	VirtualMachineSizeTypesStandardBTwos                           VirtualMachineSizeTypes = "Standard_B2s"
	VirtualMachineSizeTypesStandardDEightVThree                    VirtualMachineSizeTypes = "Standard_D8_v3"
	VirtualMachineSizeTypesStandardDEightsVThree                   VirtualMachineSizeTypes = "Standard_D8s_v3"
	VirtualMachineSizeTypesStandardDFiveVTwo                       VirtualMachineSizeTypes = "Standard_D5_v2"
	VirtualMachineSizeTypesStandardDFour                           VirtualMachineSizeTypes = "Standard_D4"
	VirtualMachineSizeTypesStandardDFourVThree                     VirtualMachineSizeTypes = "Standard_D4_v3"
	VirtualMachineSizeTypesStandardDFourVTwo                       VirtualMachineSizeTypes = "Standard_D4_v2"
	VirtualMachineSizeTypesStandardDFoursVThree                    VirtualMachineSizeTypes = "Standard_D4s_v3"
	VirtualMachineSizeTypesStandardDOne                            VirtualMachineSizeTypes = "Standard_D1"
	VirtualMachineSizeTypesStandardDOneFiveVTwo                    VirtualMachineSizeTypes = "Standard_D15_v2"
	VirtualMachineSizeTypesStandardDOneFour                        VirtualMachineSizeTypes = "Standard_D14"
	VirtualMachineSizeTypesStandardDOneFourVTwo                    VirtualMachineSizeTypes = "Standard_D14_v2"
	VirtualMachineSizeTypesStandardDOneOne                         VirtualMachineSizeTypes = "Standard_D11"
	VirtualMachineSizeTypesStandardDOneOneVTwo                     VirtualMachineSizeTypes = "Standard_D11_v2"
	VirtualMachineSizeTypesStandardDOneSixVThree                   VirtualMachineSizeTypes = "Standard_D16_v3"
	VirtualMachineSizeTypesStandardDOneSixsVThree                  VirtualMachineSizeTypes = "Standard_D16s_v3"
	VirtualMachineSizeTypesStandardDOneThree                       VirtualMachineSizeTypes = "Standard_D13"
	VirtualMachineSizeTypesStandardDOneThreeVTwo                   VirtualMachineSizeTypes = "Standard_D13_v2"
	VirtualMachineSizeTypesStandardDOneTwo                         VirtualMachineSizeTypes = "Standard_D12"
	VirtualMachineSizeTypesStandardDOneTwoVTwo                     VirtualMachineSizeTypes = "Standard_D12_v2"
	VirtualMachineSizeTypesStandardDOneVTwo                        VirtualMachineSizeTypes = "Standard_D1_v2"
	VirtualMachineSizeTypesStandardDSFiveVTwo                      VirtualMachineSizeTypes = "Standard_DS5_v2"
	VirtualMachineSizeTypesStandardDSFour                          VirtualMachineSizeTypes = "Standard_DS4"
	VirtualMachineSizeTypesStandardDSFourVTwo                      VirtualMachineSizeTypes = "Standard_DS4_v2"
	VirtualMachineSizeTypesStandardDSOne                           VirtualMachineSizeTypes = "Standard_DS1"
	VirtualMachineSizeTypesStandardDSOneFiveVTwo                   VirtualMachineSizeTypes = "Standard_DS15_v2"
	VirtualMachineSizeTypesStandardDSOneFour                       VirtualMachineSizeTypes = "Standard_DS14"
	VirtualMachineSizeTypesStandardDSOneFourNegativeEightVTwo      VirtualMachineSizeTypes = "Standard_DS14-8_v2"
	VirtualMachineSizeTypesStandardDSOneFourNegativeFourVTwo       VirtualMachineSizeTypes = "Standard_DS14-4_v2"
	VirtualMachineSizeTypesStandardDSOneFourVTwo                   VirtualMachineSizeTypes = "Standard_DS14_v2"
	VirtualMachineSizeTypesStandardDSOneOne                        VirtualMachineSizeTypes = "Standard_DS11"
	VirtualMachineSizeTypesStandardDSOneOneVTwo                    VirtualMachineSizeTypes = "Standard_DS11_v2"
	VirtualMachineSizeTypesStandardDSOneThree                      VirtualMachineSizeTypes = "Standard_DS13"
	VirtualMachineSizeTypesStandardDSOneThreeNegativeFourVTwo      VirtualMachineSizeTypes = "Standard_DS13-4_v2"
	VirtualMachineSizeTypesStandardDSOneThreeNegativeTwoVTwo       VirtualMachineSizeTypes = "Standard_DS13-2_v2"
	VirtualMachineSizeTypesStandardDSOneThreeVTwo                  VirtualMachineSizeTypes = "Standard_DS13_v2"
	VirtualMachineSizeTypesStandardDSOneTwo                        VirtualMachineSizeTypes = "Standard_DS12"
	VirtualMachineSizeTypesStandardDSOneTwoVTwo                    VirtualMachineSizeTypes = "Standard_DS12_v2"
	VirtualMachineSizeTypesStandardDSOneVTwo                       VirtualMachineSizeTypes = "Standard_DS1_v2"
	VirtualMachineSizeTypesStandardDSThree                         VirtualMachineSizeTypes = "Standard_DS3"
	VirtualMachineSizeTypesStandardDSThreeVTwo                     VirtualMachineSizeTypes = "Standard_DS3_v2"
	VirtualMachineSizeTypesStandardDSTwo                           VirtualMachineSizeTypes = "Standard_DS2"
	VirtualMachineSizeTypesStandardDSTwoVTwo                       VirtualMachineSizeTypes = "Standard_DS2_v2"
	VirtualMachineSizeTypesStandardDSixFourVThree                  VirtualMachineSizeTypes = "Standard_D64_v3"
	VirtualMachineSizeTypesStandardDSixFoursVThree                 VirtualMachineSizeTypes = "Standard_D64s_v3"
	VirtualMachineSizeTypesStandardDThree                          VirtualMachineSizeTypes = "Standard_D3"
	VirtualMachineSizeTypesStandardDThreeTwoVThree                 VirtualMachineSizeTypes = "Standard_D32_v3"
	VirtualMachineSizeTypesStandardDThreeTwosVThree                VirtualMachineSizeTypes = "Standard_D32s_v3"
	VirtualMachineSizeTypesStandardDThreeVTwo                      VirtualMachineSizeTypes = "Standard_D3_v2"
	VirtualMachineSizeTypesStandardDTwo                            VirtualMachineSizeTypes = "Standard_D2"
	VirtualMachineSizeTypesStandardDTwoVThree                      VirtualMachineSizeTypes = "Standard_D2_v3"
	VirtualMachineSizeTypesStandardDTwoVTwo                        VirtualMachineSizeTypes = "Standard_D2_v2"
	VirtualMachineSizeTypesStandardDTwosVThree                     VirtualMachineSizeTypes = "Standard_D2s_v3"
	VirtualMachineSizeTypesStandardEEightVThree                    VirtualMachineSizeTypes = "Standard_E8_v3"
	VirtualMachineSizeTypesStandardEEightsVThree                   VirtualMachineSizeTypes = "Standard_E8s_v3"
	VirtualMachineSizeTypesStandardEFourVThree                     VirtualMachineSizeTypes = "Standard_E4_v3"
	VirtualMachineSizeTypesStandardEFoursVThree                    VirtualMachineSizeTypes = "Standard_E4s_v3"
	VirtualMachineSizeTypesStandardEOneSixVThree                   VirtualMachineSizeTypes = "Standard_E16_v3"
	VirtualMachineSizeTypesStandardEOneSixsVThree                  VirtualMachineSizeTypes = "Standard_E16s_v3"
	VirtualMachineSizeTypesStandardESixFourNegativeOneSixsVThree   VirtualMachineSizeTypes = "Standard_E64-16s_v3"
	VirtualMachineSizeTypesStandardESixFourNegativeThreeTwosVThree VirtualMachineSizeTypes = "Standard_E64-32s_v3"
	VirtualMachineSizeTypesStandardESixFourVThree                  VirtualMachineSizeTypes = "Standard_E64_v3"
	VirtualMachineSizeTypesStandardESixFoursVThree                 VirtualMachineSizeTypes = "Standard_E64s_v3"
	VirtualMachineSizeTypesStandardEThreeTwoNegativeEightsVThree   VirtualMachineSizeTypes = "Standard_E32-8s_v3"
	VirtualMachineSizeTypesStandardEThreeTwoNegativeOneSixVThree   VirtualMachineSizeTypes = "Standard_E32-16_v3"
	VirtualMachineSizeTypesStandardEThreeTwoVThree                 VirtualMachineSizeTypes = "Standard_E32_v3"
	VirtualMachineSizeTypesStandardEThreeTwosVThree                VirtualMachineSizeTypes = "Standard_E32s_v3"
	VirtualMachineSizeTypesStandardETwoVThree                      VirtualMachineSizeTypes = "Standard_E2_v3"
	VirtualMachineSizeTypesStandardETwosVThree                     VirtualMachineSizeTypes = "Standard_E2s_v3"
	VirtualMachineSizeTypesStandardFEight                          VirtualMachineSizeTypes = "Standard_F8"
	VirtualMachineSizeTypesStandardFEights                         VirtualMachineSizeTypes = "Standard_F8s"
	VirtualMachineSizeTypesStandardFEightsVTwo                     VirtualMachineSizeTypes = "Standard_F8s_v2"
	VirtualMachineSizeTypesStandardFFour                           VirtualMachineSizeTypes = "Standard_F4"
	VirtualMachineSizeTypesStandardFFours                          VirtualMachineSizeTypes = "Standard_F4s"
	VirtualMachineSizeTypesStandardFFoursVTwo                      VirtualMachineSizeTypes = "Standard_F4s_v2"
	VirtualMachineSizeTypesStandardFOne                            VirtualMachineSizeTypes = "Standard_F1"
	VirtualMachineSizeTypesStandardFOneSix                         VirtualMachineSizeTypes = "Standard_F16"
	VirtualMachineSizeTypesStandardFOneSixs                        VirtualMachineSizeTypes = "Standard_F16s"
	VirtualMachineSizeTypesStandardFOneSixsVTwo                    VirtualMachineSizeTypes = "Standard_F16s_v2"
	VirtualMachineSizeTypesStandardFOnes                           VirtualMachineSizeTypes = "Standard_F1s"
	VirtualMachineSizeTypesStandardFSevenTwosVTwo                  VirtualMachineSizeTypes = "Standard_F72s_v2"
	VirtualMachineSizeTypesStandardFSixFoursVTwo                   VirtualMachineSizeTypes = "Standard_F64s_v2"
	VirtualMachineSizeTypesStandardFThreeTwosVTwo                  VirtualMachineSizeTypes = "Standard_F32s_v2"
	VirtualMachineSizeTypesStandardFTwo                            VirtualMachineSizeTypes = "Standard_F2"
	VirtualMachineSizeTypesStandardFTwos                           VirtualMachineSizeTypes = "Standard_F2s"
	VirtualMachineSizeTypesStandardFTwosVTwo                       VirtualMachineSizeTypes = "Standard_F2s_v2"
	VirtualMachineSizeTypesStandardGFive                           VirtualMachineSizeTypes = "Standard_G5"
	VirtualMachineSizeTypesStandardGFour                           VirtualMachineSizeTypes = "Standard_G4"
	VirtualMachineSizeTypesStandardGOne                            VirtualMachineSizeTypes = "Standard_G1"
	VirtualMachineSizeTypesStandardGSFive                          VirtualMachineSizeTypes = "Standard_GS5"
	VirtualMachineSizeTypesStandardGSFiveNegativeEight             VirtualMachineSizeTypes = "Standard_GS5-8"
	VirtualMachineSizeTypesStandardGSFiveNegativeOneSix            VirtualMachineSizeTypes = "Standard_GS5-16"
	VirtualMachineSizeTypesStandardGSFour                          VirtualMachineSizeTypes = "Standard_GS4"
	VirtualMachineSizeTypesStandardGSFourNegativeEight             VirtualMachineSizeTypes = "Standard_GS4-8"
	VirtualMachineSizeTypesStandardGSFourNegativeFour              VirtualMachineSizeTypes = "Standard_GS4-4"
	VirtualMachineSizeTypesStandardGSOne                           VirtualMachineSizeTypes = "Standard_GS1"
	VirtualMachineSizeTypesStandardGSThree                         VirtualMachineSizeTypes = "Standard_GS3"
	VirtualMachineSizeTypesStandardGSTwo                           VirtualMachineSizeTypes = "Standard_GS2"
	VirtualMachineSizeTypesStandardGThree                          VirtualMachineSizeTypes = "Standard_G3"
	VirtualMachineSizeTypesStandardGTwo                            VirtualMachineSizeTypes = "Standard_G2"
	VirtualMachineSizeTypesStandardHEight                          VirtualMachineSizeTypes = "Standard_H8"
	VirtualMachineSizeTypesStandardHEightm                         VirtualMachineSizeTypes = "Standard_H8m"
	VirtualMachineSizeTypesStandardHOneSix                         VirtualMachineSizeTypes = "Standard_H16"
	VirtualMachineSizeTypesStandardHOneSixm                        VirtualMachineSizeTypes = "Standard_H16m"
	VirtualMachineSizeTypesStandardHOneSixmr                       VirtualMachineSizeTypes = "Standard_H16mr"
	VirtualMachineSizeTypesStandardHOneSixr                        VirtualMachineSizeTypes = "Standard_H16r"
	VirtualMachineSizeTypesStandardLEights                         VirtualMachineSizeTypes = "Standard_L8s"
	VirtualMachineSizeTypesStandardLFours                          VirtualMachineSizeTypes = "Standard_L4s"
	VirtualMachineSizeTypesStandardLOneSixs                        VirtualMachineSizeTypes = "Standard_L16s"
	VirtualMachineSizeTypesStandardLThreeTwos                      VirtualMachineSizeTypes = "Standard_L32s"
	VirtualMachineSizeTypesStandardMOneTwoEightNegativeSixFourms   VirtualMachineSizeTypes = "Standard_M128-64ms"
	VirtualMachineSizeTypesStandardMOneTwoEightNegativeThreeTwoms  VirtualMachineSizeTypes = "Standard_M128-32ms"
	VirtualMachineSizeTypesStandardMOneTwoEightms                  VirtualMachineSizeTypes = "Standard_M128ms"
	VirtualMachineSizeTypesStandardMOneTwoEights                   VirtualMachineSizeTypes = "Standard_M128s"
	VirtualMachineSizeTypesStandardMSixFourNegativeOneSixms        VirtualMachineSizeTypes = "Standard_M64-16ms"
	VirtualMachineSizeTypesStandardMSixFourNegativeThreeTwoms      VirtualMachineSizeTypes = "Standard_M64-32ms"
	VirtualMachineSizeTypesStandardMSixFourms                      VirtualMachineSizeTypes = "Standard_M64ms"
	VirtualMachineSizeTypesStandardMSixFours                       VirtualMachineSizeTypes = "Standard_M64s"
	VirtualMachineSizeTypesStandardNCOneTwo                        VirtualMachineSizeTypes = "Standard_NC12"
	VirtualMachineSizeTypesStandardNCOneTwosVThree                 VirtualMachineSizeTypes = "Standard_NC12s_v3"
	VirtualMachineSizeTypesStandardNCOneTwosVTwo                   VirtualMachineSizeTypes = "Standard_NC12s_v2"
	VirtualMachineSizeTypesStandardNCSix                           VirtualMachineSizeTypes = "Standard_NC6"
	VirtualMachineSizeTypesStandardNCSixsVThree                    VirtualMachineSizeTypes = "Standard_NC6s_v3"
	VirtualMachineSizeTypesStandardNCSixsVTwo                      VirtualMachineSizeTypes = "Standard_NC6s_v2"
	VirtualMachineSizeTypesStandardNCTwoFour                       VirtualMachineSizeTypes = "Standard_NC24"
	VirtualMachineSizeTypesStandardNCTwoFourr                      VirtualMachineSizeTypes = "Standard_NC24r"
	VirtualMachineSizeTypesStandardNCTwoFourrsVThree               VirtualMachineSizeTypes = "Standard_NC24rs_v3"
	VirtualMachineSizeTypesStandardNCTwoFourrsVTwo                 VirtualMachineSizeTypes = "Standard_NC24rs_v2"
	VirtualMachineSizeTypesStandardNCTwoFoursVThree                VirtualMachineSizeTypes = "Standard_NC24s_v3"
	VirtualMachineSizeTypesStandardNCTwoFoursVTwo                  VirtualMachineSizeTypes = "Standard_NC24s_v2"
	VirtualMachineSizeTypesStandardNDOneTwos                       VirtualMachineSizeTypes = "Standard_ND12s"
	VirtualMachineSizeTypesStandardNDSixs                          VirtualMachineSizeTypes = "Standard_ND6s"
	VirtualMachineSizeTypesStandardNDTwoFourrs                     VirtualMachineSizeTypes = "Standard_ND24rs"
	VirtualMachineSizeTypesStandardNDTwoFours                      VirtualMachineSizeTypes = "Standard_ND24s"
	VirtualMachineSizeTypesStandardNVOneTwo                        VirtualMachineSizeTypes = "Standard_NV12"
	VirtualMachineSizeTypesStandardNVSix                           VirtualMachineSizeTypes = "Standard_NV6"
	VirtualMachineSizeTypesStandardNVTwoFour                       VirtualMachineSizeTypes = "Standard_NV24"
)

func PossibleValuesForVirtualMachineSizeTypes() []string {
	return []string{
		string(VirtualMachineSizeTypesBasicAFour),
		string(VirtualMachineSizeTypesBasicAOne),
		string(VirtualMachineSizeTypesBasicAThree),
		string(VirtualMachineSizeTypesBasicATwo),
		string(VirtualMachineSizeTypesBasicAZero),
		string(VirtualMachineSizeTypesStandardAEight),
		string(VirtualMachineSizeTypesStandardAEightVTwo),
		string(VirtualMachineSizeTypesStandardAEightmVTwo),
		string(VirtualMachineSizeTypesStandardAFive),
		string(VirtualMachineSizeTypesStandardAFour),
		string(VirtualMachineSizeTypesStandardAFourVTwo),
		string(VirtualMachineSizeTypesStandardAFourmVTwo),
		string(VirtualMachineSizeTypesStandardANine),
		string(VirtualMachineSizeTypesStandardAOne),
		string(VirtualMachineSizeTypesStandardAOneOne),
		string(VirtualMachineSizeTypesStandardAOneVTwo),
		string(VirtualMachineSizeTypesStandardAOneZero),
		string(VirtualMachineSizeTypesStandardASeven),
		string(VirtualMachineSizeTypesStandardASix),
		string(VirtualMachineSizeTypesStandardAThree),
		string(VirtualMachineSizeTypesStandardATwo),
		string(VirtualMachineSizeTypesStandardATwoVTwo),
		string(VirtualMachineSizeTypesStandardATwomVTwo),
		string(VirtualMachineSizeTypesStandardAZero),
		string(VirtualMachineSizeTypesStandardBEightms),
		string(VirtualMachineSizeTypesStandardBFourms),
		string(VirtualMachineSizeTypesStandardBOnems),
		string(VirtualMachineSizeTypesStandardBOnes),
		string(VirtualMachineSizeTypesStandardBTwoms),
		string(VirtualMachineSizeTypesStandardBTwos),
		string(VirtualMachineSizeTypesStandardDEightVThree),
		string(VirtualMachineSizeTypesStandardDEightsVThree),
		string(VirtualMachineSizeTypesStandardDFiveVTwo),
		string(VirtualMachineSizeTypesStandardDFour),
		string(VirtualMachineSizeTypesStandardDFourVThree),
		string(VirtualMachineSizeTypesStandardDFourVTwo),
		string(VirtualMachineSizeTypesStandardDFoursVThree),
		string(VirtualMachineSizeTypesStandardDOne),
		string(VirtualMachineSizeTypesStandardDOneFiveVTwo),
		string(VirtualMachineSizeTypesStandardDOneFour),
		string(VirtualMachineSizeTypesStandardDOneFourVTwo),
		string(VirtualMachineSizeTypesStandardDOneOne),
		string(VirtualMachineSizeTypesStandardDOneOneVTwo),
		string(VirtualMachineSizeTypesStandardDOneSixVThree),
		string(VirtualMachineSizeTypesStandardDOneSixsVThree),
		string(VirtualMachineSizeTypesStandardDOneThree),
		string(VirtualMachineSizeTypesStandardDOneThreeVTwo),
		string(VirtualMachineSizeTypesStandardDOneTwo),
		string(VirtualMachineSizeTypesStandardDOneTwoVTwo),
		string(VirtualMachineSizeTypesStandardDOneVTwo),
		string(VirtualMachineSizeTypesStandardDSFiveVTwo),
		string(VirtualMachineSizeTypesStandardDSFour),
		string(VirtualMachineSizeTypesStandardDSFourVTwo),
		string(VirtualMachineSizeTypesStandardDSOne),
		string(VirtualMachineSizeTypesStandardDSOneFiveVTwo),
		string(VirtualMachineSizeTypesStandardDSOneFour),
		string(VirtualMachineSizeTypesStandardDSOneFourNegativeEightVTwo),
		string(VirtualMachineSizeTypesStandardDSOneFourNegativeFourVTwo),
		string(VirtualMachineSizeTypesStandardDSOneFourVTwo),
		string(VirtualMachineSizeTypesStandardDSOneOne),
		string(VirtualMachineSizeTypesStandardDSOneOneVTwo),
		string(VirtualMachineSizeTypesStandardDSOneThree),
		string(VirtualMachineSizeTypesStandardDSOneThreeNegativeFourVTwo),
		string(VirtualMachineSizeTypesStandardDSOneThreeNegativeTwoVTwo),
		string(VirtualMachineSizeTypesStandardDSOneThreeVTwo),
		string(VirtualMachineSizeTypesStandardDSOneTwo),
		string(VirtualMachineSizeTypesStandardDSOneTwoVTwo),
		string(VirtualMachineSizeTypesStandardDSOneVTwo),
		string(VirtualMachineSizeTypesStandardDSThree),
		string(VirtualMachineSizeTypesStandardDSThreeVTwo),
		string(VirtualMachineSizeTypesStandardDSTwo),
		string(VirtualMachineSizeTypesStandardDSTwoVTwo),
		string(VirtualMachineSizeTypesStandardDSixFourVThree),
		string(VirtualMachineSizeTypesStandardDSixFoursVThree),
		string(VirtualMachineSizeTypesStandardDThree),
		string(VirtualMachineSizeTypesStandardDThreeTwoVThree),
		string(VirtualMachineSizeTypesStandardDThreeTwosVThree),
		string(VirtualMachineSizeTypesStandardDThreeVTwo),
		string(VirtualMachineSizeTypesStandardDTwo),
		string(VirtualMachineSizeTypesStandardDTwoVThree),
		string(VirtualMachineSizeTypesStandardDTwoVTwo),
		string(VirtualMachineSizeTypesStandardDTwosVThree),
		string(VirtualMachineSizeTypesStandardEEightVThree),
		string(VirtualMachineSizeTypesStandardEEightsVThree),
		string(VirtualMachineSizeTypesStandardEFourVThree),
		string(VirtualMachineSizeTypesStandardEFoursVThree),
		string(VirtualMachineSizeTypesStandardEOneSixVThree),
		string(VirtualMachineSizeTypesStandardEOneSixsVThree),
		string(VirtualMachineSizeTypesStandardESixFourNegativeOneSixsVThree),
		string(VirtualMachineSizeTypesStandardESixFourNegativeThreeTwosVThree),
		string(VirtualMachineSizeTypesStandardESixFourVThree),
		string(VirtualMachineSizeTypesStandardESixFoursVThree),
		string(VirtualMachineSizeTypesStandardEThreeTwoNegativeEightsVThree),
		string(VirtualMachineSizeTypesStandardEThreeTwoNegativeOneSixVThree),
		string(VirtualMachineSizeTypesStandardEThreeTwoVThree),
		string(VirtualMachineSizeTypesStandardEThreeTwosVThree),
		string(VirtualMachineSizeTypesStandardETwoVThree),
		string(VirtualMachineSizeTypesStandardETwosVThree),
		string(VirtualMachineSizeTypesStandardFEight),
		string(VirtualMachineSizeTypesStandardFEights),
		string(VirtualMachineSizeTypesStandardFEightsVTwo),
		string(VirtualMachineSizeTypesStandardFFour),
		string(VirtualMachineSizeTypesStandardFFours),
		string(VirtualMachineSizeTypesStandardFFoursVTwo),
		string(VirtualMachineSizeTypesStandardFOne),
		string(VirtualMachineSizeTypesStandardFOneSix),
		string(VirtualMachineSizeTypesStandardFOneSixs),
		string(VirtualMachineSizeTypesStandardFOneSixsVTwo),
		string(VirtualMachineSizeTypesStandardFOnes),
		string(VirtualMachineSizeTypesStandardFSevenTwosVTwo),
		string(VirtualMachineSizeTypesStandardFSixFoursVTwo),
		string(VirtualMachineSizeTypesStandardFThreeTwosVTwo),
		string(VirtualMachineSizeTypesStandardFTwo),
		string(VirtualMachineSizeTypesStandardFTwos),
		string(VirtualMachineSizeTypesStandardFTwosVTwo),
		string(VirtualMachineSizeTypesStandardGFive),
		string(VirtualMachineSizeTypesStandardGFour),
		string(VirtualMachineSizeTypesStandardGOne),
		string(VirtualMachineSizeTypesStandardGSFive),
		string(VirtualMachineSizeTypesStandardGSFiveNegativeEight),
		string(VirtualMachineSizeTypesStandardGSFiveNegativeOneSix),
		string(VirtualMachineSizeTypesStandardGSFour),
		string(VirtualMachineSizeTypesStandardGSFourNegativeEight),
		string(VirtualMachineSizeTypesStandardGSFourNegativeFour),
		string(VirtualMachineSizeTypesStandardGSOne),
		string(VirtualMachineSizeTypesStandardGSThree),
		string(VirtualMachineSizeTypesStandardGSTwo),
		string(VirtualMachineSizeTypesStandardGThree),
		string(VirtualMachineSizeTypesStandardGTwo),
		string(VirtualMachineSizeTypesStandardHEight),
		string(VirtualMachineSizeTypesStandardHEightm),
		string(VirtualMachineSizeTypesStandardHOneSix),
		string(VirtualMachineSizeTypesStandardHOneSixm),
		string(VirtualMachineSizeTypesStandardHOneSixmr),
		string(VirtualMachineSizeTypesStandardHOneSixr),
		string(VirtualMachineSizeTypesStandardLEights),
		string(VirtualMachineSizeTypesStandardLFours),
		string(VirtualMachineSizeTypesStandardLOneSixs),
		string(VirtualMachineSizeTypesStandardLThreeTwos),
		string(VirtualMachineSizeTypesStandardMOneTwoEightNegativeSixFourms),
		string(VirtualMachineSizeTypesStandardMOneTwoEightNegativeThreeTwoms),
		string(VirtualMachineSizeTypesStandardMOneTwoEightms),
		string(VirtualMachineSizeTypesStandardMOneTwoEights),
		string(VirtualMachineSizeTypesStandardMSixFourNegativeOneSixms),
		string(VirtualMachineSizeTypesStandardMSixFourNegativeThreeTwoms),
		string(VirtualMachineSizeTypesStandardMSixFourms),
		string(VirtualMachineSizeTypesStandardMSixFours),
		string(VirtualMachineSizeTypesStandardNCOneTwo),
		string(VirtualMachineSizeTypesStandardNCOneTwosVThree),
		string(VirtualMachineSizeTypesStandardNCOneTwosVTwo),
		string(VirtualMachineSizeTypesStandardNCSix),
		string(VirtualMachineSizeTypesStandardNCSixsVThree),
		string(VirtualMachineSizeTypesStandardNCSixsVTwo),
		string(VirtualMachineSizeTypesStandardNCTwoFour),
		string(VirtualMachineSizeTypesStandardNCTwoFourr),
		string(VirtualMachineSizeTypesStandardNCTwoFourrsVThree),
		string(VirtualMachineSizeTypesStandardNCTwoFourrsVTwo),
		string(VirtualMachineSizeTypesStandardNCTwoFoursVThree),
		string(VirtualMachineSizeTypesStandardNCTwoFoursVTwo),
		string(VirtualMachineSizeTypesStandardNDOneTwos),
		string(VirtualMachineSizeTypesStandardNDSixs),
		string(VirtualMachineSizeTypesStandardNDTwoFourrs),
		string(VirtualMachineSizeTypesStandardNDTwoFours),
		string(VirtualMachineSizeTypesStandardNVOneTwo),
		string(VirtualMachineSizeTypesStandardNVSix),
		string(VirtualMachineSizeTypesStandardNVTwoFour),
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
