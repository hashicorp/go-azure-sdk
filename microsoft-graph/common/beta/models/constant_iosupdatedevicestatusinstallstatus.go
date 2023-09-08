package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosUpdateDeviceStatusInstallStatus string

const (
	IosUpdateDeviceStatusInstallStatusavailable                          IosUpdateDeviceStatusInstallStatus = "Available"
	IosUpdateDeviceStatusInstallStatusdeviceOsHigherThanDesiredOsVersion IosUpdateDeviceStatusInstallStatus = "DeviceOsHigherThanDesiredOsVersion"
	IosUpdateDeviceStatusInstallStatusdownloadFailed                     IosUpdateDeviceStatusInstallStatus = "DownloadFailed"
	IosUpdateDeviceStatusInstallStatusdownloadInsufficientNetwork        IosUpdateDeviceStatusInstallStatus = "DownloadInsufficientNetwork"
	IosUpdateDeviceStatusInstallStatusdownloadInsufficientPower          IosUpdateDeviceStatusInstallStatus = "DownloadInsufficientPower"
	IosUpdateDeviceStatusInstallStatusdownloadInsufficientSpace          IosUpdateDeviceStatusInstallStatus = "DownloadInsufficientSpace"
	IosUpdateDeviceStatusInstallStatusdownloadRequiresComputer           IosUpdateDeviceStatusInstallStatus = "DownloadRequiresComputer"
	IosUpdateDeviceStatusInstallStatusdownloading                        IosUpdateDeviceStatusInstallStatus = "Downloading"
	IosUpdateDeviceStatusInstallStatusidle                               IosUpdateDeviceStatusInstallStatus = "Idle"
	IosUpdateDeviceStatusInstallStatusinstallFailed                      IosUpdateDeviceStatusInstallStatus = "InstallFailed"
	IosUpdateDeviceStatusInstallStatusinstallInsufficientPower           IosUpdateDeviceStatusInstallStatus = "InstallInsufficientPower"
	IosUpdateDeviceStatusInstallStatusinstallInsufficientSpace           IosUpdateDeviceStatusInstallStatus = "InstallInsufficientSpace"
	IosUpdateDeviceStatusInstallStatusinstallPhoneCallInProgress         IosUpdateDeviceStatusInstallStatus = "InstallPhoneCallInProgress"
	IosUpdateDeviceStatusInstallStatusinstalling                         IosUpdateDeviceStatusInstallStatus = "Installing"
	IosUpdateDeviceStatusInstallStatusmdmClientCrashed                   IosUpdateDeviceStatusInstallStatus = "MdmClientCrashed"
	IosUpdateDeviceStatusInstallStatusnotSupportedOperation              IosUpdateDeviceStatusInstallStatus = "NotSupportedOperation"
	IosUpdateDeviceStatusInstallStatussharedDeviceUserLoggedInError      IosUpdateDeviceStatusInstallStatus = "SharedDeviceUserLoggedInError"
	IosUpdateDeviceStatusInstallStatussuccess                            IosUpdateDeviceStatusInstallStatus = "Success"
	IosUpdateDeviceStatusInstallStatustimeout                            IosUpdateDeviceStatusInstallStatus = "Timeout"
	IosUpdateDeviceStatusInstallStatusunknown                            IosUpdateDeviceStatusInstallStatus = "Unknown"
	IosUpdateDeviceStatusInstallStatusupdateError                        IosUpdateDeviceStatusInstallStatus = "UpdateError"
	IosUpdateDeviceStatusInstallStatusupdateScanFailed                   IosUpdateDeviceStatusInstallStatus = "UpdateScanFailed"
)

func PossibleValuesForIosUpdateDeviceStatusInstallStatus() []string {
	return []string{
		string(IosUpdateDeviceStatusInstallStatusavailable),
		string(IosUpdateDeviceStatusInstallStatusdeviceOsHigherThanDesiredOsVersion),
		string(IosUpdateDeviceStatusInstallStatusdownloadFailed),
		string(IosUpdateDeviceStatusInstallStatusdownloadInsufficientNetwork),
		string(IosUpdateDeviceStatusInstallStatusdownloadInsufficientPower),
		string(IosUpdateDeviceStatusInstallStatusdownloadInsufficientSpace),
		string(IosUpdateDeviceStatusInstallStatusdownloadRequiresComputer),
		string(IosUpdateDeviceStatusInstallStatusdownloading),
		string(IosUpdateDeviceStatusInstallStatusidle),
		string(IosUpdateDeviceStatusInstallStatusinstallFailed),
		string(IosUpdateDeviceStatusInstallStatusinstallInsufficientPower),
		string(IosUpdateDeviceStatusInstallStatusinstallInsufficientSpace),
		string(IosUpdateDeviceStatusInstallStatusinstallPhoneCallInProgress),
		string(IosUpdateDeviceStatusInstallStatusinstalling),
		string(IosUpdateDeviceStatusInstallStatusmdmClientCrashed),
		string(IosUpdateDeviceStatusInstallStatusnotSupportedOperation),
		string(IosUpdateDeviceStatusInstallStatussharedDeviceUserLoggedInError),
		string(IosUpdateDeviceStatusInstallStatussuccess),
		string(IosUpdateDeviceStatusInstallStatustimeout),
		string(IosUpdateDeviceStatusInstallStatusunknown),
		string(IosUpdateDeviceStatusInstallStatusupdateError),
		string(IosUpdateDeviceStatusInstallStatusupdateScanFailed),
	}
}

func parseIosUpdateDeviceStatusInstallStatus(input string) (*IosUpdateDeviceStatusInstallStatus, error) {
	vals := map[string]IosUpdateDeviceStatusInstallStatus{
		"available":                          IosUpdateDeviceStatusInstallStatusavailable,
		"deviceoshigherthandesiredosversion": IosUpdateDeviceStatusInstallStatusdeviceOsHigherThanDesiredOsVersion,
		"downloadfailed":                     IosUpdateDeviceStatusInstallStatusdownloadFailed,
		"downloadinsufficientnetwork":        IosUpdateDeviceStatusInstallStatusdownloadInsufficientNetwork,
		"downloadinsufficientpower":          IosUpdateDeviceStatusInstallStatusdownloadInsufficientPower,
		"downloadinsufficientspace":          IosUpdateDeviceStatusInstallStatusdownloadInsufficientSpace,
		"downloadrequirescomputer":           IosUpdateDeviceStatusInstallStatusdownloadRequiresComputer,
		"downloading":                        IosUpdateDeviceStatusInstallStatusdownloading,
		"idle":                               IosUpdateDeviceStatusInstallStatusidle,
		"installfailed":                      IosUpdateDeviceStatusInstallStatusinstallFailed,
		"installinsufficientpower":           IosUpdateDeviceStatusInstallStatusinstallInsufficientPower,
		"installinsufficientspace":           IosUpdateDeviceStatusInstallStatusinstallInsufficientSpace,
		"installphonecallinprogress":         IosUpdateDeviceStatusInstallStatusinstallPhoneCallInProgress,
		"installing":                         IosUpdateDeviceStatusInstallStatusinstalling,
		"mdmclientcrashed":                   IosUpdateDeviceStatusInstallStatusmdmClientCrashed,
		"notsupportedoperation":              IosUpdateDeviceStatusInstallStatusnotSupportedOperation,
		"shareddeviceuserloggedinerror":      IosUpdateDeviceStatusInstallStatussharedDeviceUserLoggedInError,
		"success":                            IosUpdateDeviceStatusInstallStatussuccess,
		"timeout":                            IosUpdateDeviceStatusInstallStatustimeout,
		"unknown":                            IosUpdateDeviceStatusInstallStatusunknown,
		"updateerror":                        IosUpdateDeviceStatusInstallStatusupdateError,
		"updatescanfailed":                   IosUpdateDeviceStatusInstallStatusupdateScanFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosUpdateDeviceStatusInstallStatus(input)
	return &out, nil
}
