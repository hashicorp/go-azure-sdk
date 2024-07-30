package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosUpdateDeviceStatusInstallStatus string

const (
	IosUpdateDeviceStatusInstallStatus_Available                          IosUpdateDeviceStatusInstallStatus = "available"
	IosUpdateDeviceStatusInstallStatus_DeviceOsHigherThanDesiredOsVersion IosUpdateDeviceStatusInstallStatus = "deviceOsHigherThanDesiredOsVersion"
	IosUpdateDeviceStatusInstallStatus_DownloadFailed                     IosUpdateDeviceStatusInstallStatus = "downloadFailed"
	IosUpdateDeviceStatusInstallStatus_DownloadInsufficientNetwork        IosUpdateDeviceStatusInstallStatus = "downloadInsufficientNetwork"
	IosUpdateDeviceStatusInstallStatus_DownloadInsufficientPower          IosUpdateDeviceStatusInstallStatus = "downloadInsufficientPower"
	IosUpdateDeviceStatusInstallStatus_DownloadInsufficientSpace          IosUpdateDeviceStatusInstallStatus = "downloadInsufficientSpace"
	IosUpdateDeviceStatusInstallStatus_DownloadRequiresComputer           IosUpdateDeviceStatusInstallStatus = "downloadRequiresComputer"
	IosUpdateDeviceStatusInstallStatus_Downloading                        IosUpdateDeviceStatusInstallStatus = "downloading"
	IosUpdateDeviceStatusInstallStatus_Idle                               IosUpdateDeviceStatusInstallStatus = "idle"
	IosUpdateDeviceStatusInstallStatus_InstallFailed                      IosUpdateDeviceStatusInstallStatus = "installFailed"
	IosUpdateDeviceStatusInstallStatus_InstallInsufficientPower           IosUpdateDeviceStatusInstallStatus = "installInsufficientPower"
	IosUpdateDeviceStatusInstallStatus_InstallInsufficientSpace           IosUpdateDeviceStatusInstallStatus = "installInsufficientSpace"
	IosUpdateDeviceStatusInstallStatus_InstallPhoneCallInProgress         IosUpdateDeviceStatusInstallStatus = "installPhoneCallInProgress"
	IosUpdateDeviceStatusInstallStatus_Installing                         IosUpdateDeviceStatusInstallStatus = "installing"
	IosUpdateDeviceStatusInstallStatus_NotSupportedOperation              IosUpdateDeviceStatusInstallStatus = "notSupportedOperation"
	IosUpdateDeviceStatusInstallStatus_SharedDeviceUserLoggedInError      IosUpdateDeviceStatusInstallStatus = "sharedDeviceUserLoggedInError"
	IosUpdateDeviceStatusInstallStatus_Success                            IosUpdateDeviceStatusInstallStatus = "success"
	IosUpdateDeviceStatusInstallStatus_Unknown                            IosUpdateDeviceStatusInstallStatus = "unknown"
)

func PossibleValuesForIosUpdateDeviceStatusInstallStatus() []string {
	return []string{
		string(IosUpdateDeviceStatusInstallStatus_Available),
		string(IosUpdateDeviceStatusInstallStatus_DeviceOsHigherThanDesiredOsVersion),
		string(IosUpdateDeviceStatusInstallStatus_DownloadFailed),
		string(IosUpdateDeviceStatusInstallStatus_DownloadInsufficientNetwork),
		string(IosUpdateDeviceStatusInstallStatus_DownloadInsufficientPower),
		string(IosUpdateDeviceStatusInstallStatus_DownloadInsufficientSpace),
		string(IosUpdateDeviceStatusInstallStatus_DownloadRequiresComputer),
		string(IosUpdateDeviceStatusInstallStatus_Downloading),
		string(IosUpdateDeviceStatusInstallStatus_Idle),
		string(IosUpdateDeviceStatusInstallStatus_InstallFailed),
		string(IosUpdateDeviceStatusInstallStatus_InstallInsufficientPower),
		string(IosUpdateDeviceStatusInstallStatus_InstallInsufficientSpace),
		string(IosUpdateDeviceStatusInstallStatus_InstallPhoneCallInProgress),
		string(IosUpdateDeviceStatusInstallStatus_Installing),
		string(IosUpdateDeviceStatusInstallStatus_NotSupportedOperation),
		string(IosUpdateDeviceStatusInstallStatus_SharedDeviceUserLoggedInError),
		string(IosUpdateDeviceStatusInstallStatus_Success),
		string(IosUpdateDeviceStatusInstallStatus_Unknown),
	}
}

func (s *IosUpdateDeviceStatusInstallStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosUpdateDeviceStatusInstallStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosUpdateDeviceStatusInstallStatus(input string) (*IosUpdateDeviceStatusInstallStatus, error) {
	vals := map[string]IosUpdateDeviceStatusInstallStatus{
		"available":                          IosUpdateDeviceStatusInstallStatus_Available,
		"deviceoshigherthandesiredosversion": IosUpdateDeviceStatusInstallStatus_DeviceOsHigherThanDesiredOsVersion,
		"downloadfailed":                     IosUpdateDeviceStatusInstallStatus_DownloadFailed,
		"downloadinsufficientnetwork":        IosUpdateDeviceStatusInstallStatus_DownloadInsufficientNetwork,
		"downloadinsufficientpower":          IosUpdateDeviceStatusInstallStatus_DownloadInsufficientPower,
		"downloadinsufficientspace":          IosUpdateDeviceStatusInstallStatus_DownloadInsufficientSpace,
		"downloadrequirescomputer":           IosUpdateDeviceStatusInstallStatus_DownloadRequiresComputer,
		"downloading":                        IosUpdateDeviceStatusInstallStatus_Downloading,
		"idle":                               IosUpdateDeviceStatusInstallStatus_Idle,
		"installfailed":                      IosUpdateDeviceStatusInstallStatus_InstallFailed,
		"installinsufficientpower":           IosUpdateDeviceStatusInstallStatus_InstallInsufficientPower,
		"installinsufficientspace":           IosUpdateDeviceStatusInstallStatus_InstallInsufficientSpace,
		"installphonecallinprogress":         IosUpdateDeviceStatusInstallStatus_InstallPhoneCallInProgress,
		"installing":                         IosUpdateDeviceStatusInstallStatus_Installing,
		"notsupportedoperation":              IosUpdateDeviceStatusInstallStatus_NotSupportedOperation,
		"shareddeviceuserloggedinerror":      IosUpdateDeviceStatusInstallStatus_SharedDeviceUserLoggedInError,
		"success":                            IosUpdateDeviceStatusInstallStatus_Success,
		"unknown":                            IosUpdateDeviceStatusInstallStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosUpdateDeviceStatusInstallStatus(input)
	return &out, nil
}
