package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteActionAuditAction string

const (
	RemoteActionAuditAction_ActivateDeviceEsim                                   RemoteActionAuditAction = "activateDeviceEsim"
	RemoteActionAuditAction_AutomaticRedeployment                                RemoteActionAuditAction = "automaticRedeployment"
	RemoteActionAuditAction_CleanWindowsDevice                                   RemoteActionAuditAction = "cleanWindowsDevice"
	RemoteActionAuditAction_Deprovision                                          RemoteActionAuditAction = "deprovision"
	RemoteActionAuditAction_Disable                                              RemoteActionAuditAction = "disable"
	RemoteActionAuditAction_DisableLostMode                                      RemoteActionAuditAction = "disableLostMode"
	RemoteActionAuditAction_EnableLostMode                                       RemoteActionAuditAction = "enableLostMode"
	RemoteActionAuditAction_FactoryReset                                         RemoteActionAuditAction = "factoryReset"
	RemoteActionAuditAction_FactoryResetKeepEnrollmentData                       RemoteActionAuditAction = "factoryResetKeepEnrollmentData"
	RemoteActionAuditAction_FullScan                                             RemoteActionAuditAction = "fullScan"
	RemoteActionAuditAction_GetFileVaultKey                                      RemoteActionAuditAction = "getFileVaultKey"
	RemoteActionAuditAction_InitiateMobileDeviceManagementKeyRecovery            RemoteActionAuditAction = "initiateMobileDeviceManagementKeyRecovery"
	RemoteActionAuditAction_InitiateOnDemandProactiveRemediation                 RemoteActionAuditAction = "initiateOnDemandProactiveRemediation"
	RemoteActionAuditAction_LaunchRemoteHelp                                     RemoteActionAuditAction = "launchRemoteHelp"
	RemoteActionAuditAction_LocateDevice                                         RemoteActionAuditAction = "locateDevice"
	RemoteActionAuditAction_LogoutSharedAppleDeviceActiveUser                    RemoteActionAuditAction = "logoutSharedAppleDeviceActiveUser"
	RemoteActionAuditAction_MoveDeviceToOrganizationalUnit                       RemoteActionAuditAction = "moveDeviceToOrganizationalUnit"
	RemoteActionAuditAction_PauseConfigurationRefresh                            RemoteActionAuditAction = "pauseConfigurationRefresh"
	RemoteActionAuditAction_QuickScan                                            RemoteActionAuditAction = "quickScan"
	RemoteActionAuditAction_RebootNow                                            RemoteActionAuditAction = "rebootNow"
	RemoteActionAuditAction_RecoverPasscode                                      RemoteActionAuditAction = "recoverPasscode"
	RemoteActionAuditAction_Reenable                                             RemoteActionAuditAction = "reenable"
	RemoteActionAuditAction_RemoteLock                                           RemoteActionAuditAction = "remoteLock"
	RemoteActionAuditAction_RemoveCompanyData                                    RemoteActionAuditAction = "removeCompanyData"
	RemoteActionAuditAction_RemoveDeviceFirmwareConfigurationInterfaceManagement RemoteActionAuditAction = "removeDeviceFirmwareConfigurationInterfaceManagement"
	RemoteActionAuditAction_ResetPasscode                                        RemoteActionAuditAction = "resetPasscode"
	RemoteActionAuditAction_RevokeAppleVppLicenses                               RemoteActionAuditAction = "revokeAppleVppLicenses"
	RemoteActionAuditAction_RotateBitLockerKeys                                  RemoteActionAuditAction = "rotateBitLockerKeys"
	RemoteActionAuditAction_RotateFileVaultKey                                   RemoteActionAuditAction = "rotateFileVaultKey"
	RemoteActionAuditAction_RotateLocalAdminPassword                             RemoteActionAuditAction = "rotateLocalAdminPassword"
	RemoteActionAuditAction_SetDeviceName                                        RemoteActionAuditAction = "setDeviceName"
	RemoteActionAuditAction_ShutDown                                             RemoteActionAuditAction = "shutDown"
	RemoteActionAuditAction_Unknown                                              RemoteActionAuditAction = "unknown"
	RemoteActionAuditAction_UpdateDeviceAccount                                  RemoteActionAuditAction = "updateDeviceAccount"
	RemoteActionAuditAction_WindowsDefenderUpdateSignatures                      RemoteActionAuditAction = "windowsDefenderUpdateSignatures"
)

func PossibleValuesForRemoteActionAuditAction() []string {
	return []string{
		string(RemoteActionAuditAction_ActivateDeviceEsim),
		string(RemoteActionAuditAction_AutomaticRedeployment),
		string(RemoteActionAuditAction_CleanWindowsDevice),
		string(RemoteActionAuditAction_Deprovision),
		string(RemoteActionAuditAction_Disable),
		string(RemoteActionAuditAction_DisableLostMode),
		string(RemoteActionAuditAction_EnableLostMode),
		string(RemoteActionAuditAction_FactoryReset),
		string(RemoteActionAuditAction_FactoryResetKeepEnrollmentData),
		string(RemoteActionAuditAction_FullScan),
		string(RemoteActionAuditAction_GetFileVaultKey),
		string(RemoteActionAuditAction_InitiateMobileDeviceManagementKeyRecovery),
		string(RemoteActionAuditAction_InitiateOnDemandProactiveRemediation),
		string(RemoteActionAuditAction_LaunchRemoteHelp),
		string(RemoteActionAuditAction_LocateDevice),
		string(RemoteActionAuditAction_LogoutSharedAppleDeviceActiveUser),
		string(RemoteActionAuditAction_MoveDeviceToOrganizationalUnit),
		string(RemoteActionAuditAction_PauseConfigurationRefresh),
		string(RemoteActionAuditAction_QuickScan),
		string(RemoteActionAuditAction_RebootNow),
		string(RemoteActionAuditAction_RecoverPasscode),
		string(RemoteActionAuditAction_Reenable),
		string(RemoteActionAuditAction_RemoteLock),
		string(RemoteActionAuditAction_RemoveCompanyData),
		string(RemoteActionAuditAction_RemoveDeviceFirmwareConfigurationInterfaceManagement),
		string(RemoteActionAuditAction_ResetPasscode),
		string(RemoteActionAuditAction_RevokeAppleVppLicenses),
		string(RemoteActionAuditAction_RotateBitLockerKeys),
		string(RemoteActionAuditAction_RotateFileVaultKey),
		string(RemoteActionAuditAction_RotateLocalAdminPassword),
		string(RemoteActionAuditAction_SetDeviceName),
		string(RemoteActionAuditAction_ShutDown),
		string(RemoteActionAuditAction_Unknown),
		string(RemoteActionAuditAction_UpdateDeviceAccount),
		string(RemoteActionAuditAction_WindowsDefenderUpdateSignatures),
	}
}

func (s *RemoteActionAuditAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemoteActionAuditAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemoteActionAuditAction(input string) (*RemoteActionAuditAction, error) {
	vals := map[string]RemoteActionAuditAction{
		"activatedeviceesim":             RemoteActionAuditAction_ActivateDeviceEsim,
		"automaticredeployment":          RemoteActionAuditAction_AutomaticRedeployment,
		"cleanwindowsdevice":             RemoteActionAuditAction_CleanWindowsDevice,
		"deprovision":                    RemoteActionAuditAction_Deprovision,
		"disable":                        RemoteActionAuditAction_Disable,
		"disablelostmode":                RemoteActionAuditAction_DisableLostMode,
		"enablelostmode":                 RemoteActionAuditAction_EnableLostMode,
		"factoryreset":                   RemoteActionAuditAction_FactoryReset,
		"factoryresetkeepenrollmentdata": RemoteActionAuditAction_FactoryResetKeepEnrollmentData,
		"fullscan":                       RemoteActionAuditAction_FullScan,
		"getfilevaultkey":                RemoteActionAuditAction_GetFileVaultKey,
		"initiatemobiledevicemanagementkeyrecovery": RemoteActionAuditAction_InitiateMobileDeviceManagementKeyRecovery,
		"initiateondemandproactiveremediation":      RemoteActionAuditAction_InitiateOnDemandProactiveRemediation,
		"launchremotehelp":                          RemoteActionAuditAction_LaunchRemoteHelp,
		"locatedevice":                              RemoteActionAuditAction_LocateDevice,
		"logoutsharedappledeviceactiveuser":         RemoteActionAuditAction_LogoutSharedAppleDeviceActiveUser,
		"movedevicetoorganizationalunit":            RemoteActionAuditAction_MoveDeviceToOrganizationalUnit,
		"pauseconfigurationrefresh":                 RemoteActionAuditAction_PauseConfigurationRefresh,
		"quickscan":                                 RemoteActionAuditAction_QuickScan,
		"rebootnow":                                 RemoteActionAuditAction_RebootNow,
		"recoverpasscode":                           RemoteActionAuditAction_RecoverPasscode,
		"reenable":                                  RemoteActionAuditAction_Reenable,
		"remotelock":                                RemoteActionAuditAction_RemoteLock,
		"removecompanydata":                         RemoteActionAuditAction_RemoveCompanyData,
		"removedevicefirmwareconfigurationinterfacemanagement": RemoteActionAuditAction_RemoveDeviceFirmwareConfigurationInterfaceManagement,
		"resetpasscode":                   RemoteActionAuditAction_ResetPasscode,
		"revokeapplevpplicenses":          RemoteActionAuditAction_RevokeAppleVppLicenses,
		"rotatebitlockerkeys":             RemoteActionAuditAction_RotateBitLockerKeys,
		"rotatefilevaultkey":              RemoteActionAuditAction_RotateFileVaultKey,
		"rotatelocaladminpassword":        RemoteActionAuditAction_RotateLocalAdminPassword,
		"setdevicename":                   RemoteActionAuditAction_SetDeviceName,
		"shutdown":                        RemoteActionAuditAction_ShutDown,
		"unknown":                         RemoteActionAuditAction_Unknown,
		"updatedeviceaccount":             RemoteActionAuditAction_UpdateDeviceAccount,
		"windowsdefenderupdatesignatures": RemoteActionAuditAction_WindowsDefenderUpdateSignatures,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteActionAuditAction(input)
	return &out, nil
}
