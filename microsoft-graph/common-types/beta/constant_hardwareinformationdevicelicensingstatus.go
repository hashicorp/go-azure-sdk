package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareInformationDeviceLicensingStatus string

const (
	HardwareInformationDeviceLicensingStatus_AcquiringDeviceLicense                HardwareInformationDeviceLicensingStatus = "acquiringDeviceLicense"
	HardwareInformationDeviceLicensingStatus_DeviceIdentityVerificationFailed      HardwareInformationDeviceLicensingStatus = "deviceIdentityVerificationFailed"
	HardwareInformationDeviceLicensingStatus_DeviceIsNotAzureActiveDirectoryJoined HardwareInformationDeviceLicensingStatus = "deviceIsNotAzureActiveDirectoryJoined"
	HardwareInformationDeviceLicensingStatus_DeviceLicenseRefreshFailed            HardwareInformationDeviceLicensingStatus = "deviceLicenseRefreshFailed"
	HardwareInformationDeviceLicensingStatus_DeviceLicenseRefreshSucceed           HardwareInformationDeviceLicensingStatus = "deviceLicenseRefreshSucceed"
	HardwareInformationDeviceLicensingStatus_DeviceLicenseRemoveFailed             HardwareInformationDeviceLicensingStatus = "deviceLicenseRemoveFailed"
	HardwareInformationDeviceLicensingStatus_DeviceLicenseRemoveSucceed            HardwareInformationDeviceLicensingStatus = "deviceLicenseRemoveSucceed"
	HardwareInformationDeviceLicensingStatus_LicenseRefreshPending                 HardwareInformationDeviceLicensingStatus = "licenseRefreshPending"
	HardwareInformationDeviceLicensingStatus_LicenseRefreshStarted                 HardwareInformationDeviceLicensingStatus = "licenseRefreshStarted"
	HardwareInformationDeviceLicensingStatus_MicrosoftAccountVerificationFailed    HardwareInformationDeviceLicensingStatus = "microsoftAccountVerificationFailed"
	HardwareInformationDeviceLicensingStatus_RefreshingDeviceLicense               HardwareInformationDeviceLicensingStatus = "refreshingDeviceLicense"
	HardwareInformationDeviceLicensingStatus_RemovingDeviceLicense                 HardwareInformationDeviceLicensingStatus = "removingDeviceLicense"
	HardwareInformationDeviceLicensingStatus_Unknown                               HardwareInformationDeviceLicensingStatus = "unknown"
	HardwareInformationDeviceLicensingStatus_VerifyingMicrosoftAccountIdentity     HardwareInformationDeviceLicensingStatus = "verifyingMicrosoftAccountIdentity"
	HardwareInformationDeviceLicensingStatus_VerifyingMicrosoftDeviceIdentity      HardwareInformationDeviceLicensingStatus = "verifyingMicrosoftDeviceIdentity"
)

func PossibleValuesForHardwareInformationDeviceLicensingStatus() []string {
	return []string{
		string(HardwareInformationDeviceLicensingStatus_AcquiringDeviceLicense),
		string(HardwareInformationDeviceLicensingStatus_DeviceIdentityVerificationFailed),
		string(HardwareInformationDeviceLicensingStatus_DeviceIsNotAzureActiveDirectoryJoined),
		string(HardwareInformationDeviceLicensingStatus_DeviceLicenseRefreshFailed),
		string(HardwareInformationDeviceLicensingStatus_DeviceLicenseRefreshSucceed),
		string(HardwareInformationDeviceLicensingStatus_DeviceLicenseRemoveFailed),
		string(HardwareInformationDeviceLicensingStatus_DeviceLicenseRemoveSucceed),
		string(HardwareInformationDeviceLicensingStatus_LicenseRefreshPending),
		string(HardwareInformationDeviceLicensingStatus_LicenseRefreshStarted),
		string(HardwareInformationDeviceLicensingStatus_MicrosoftAccountVerificationFailed),
		string(HardwareInformationDeviceLicensingStatus_RefreshingDeviceLicense),
		string(HardwareInformationDeviceLicensingStatus_RemovingDeviceLicense),
		string(HardwareInformationDeviceLicensingStatus_Unknown),
		string(HardwareInformationDeviceLicensingStatus_VerifyingMicrosoftAccountIdentity),
		string(HardwareInformationDeviceLicensingStatus_VerifyingMicrosoftDeviceIdentity),
	}
}

func (s *HardwareInformationDeviceLicensingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHardwareInformationDeviceLicensingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHardwareInformationDeviceLicensingStatus(input string) (*HardwareInformationDeviceLicensingStatus, error) {
	vals := map[string]HardwareInformationDeviceLicensingStatus{
		"acquiringdevicelicense":                HardwareInformationDeviceLicensingStatus_AcquiringDeviceLicense,
		"deviceidentityverificationfailed":      HardwareInformationDeviceLicensingStatus_DeviceIdentityVerificationFailed,
		"deviceisnotazureactivedirectoryjoined": HardwareInformationDeviceLicensingStatus_DeviceIsNotAzureActiveDirectoryJoined,
		"devicelicenserefreshfailed":            HardwareInformationDeviceLicensingStatus_DeviceLicenseRefreshFailed,
		"devicelicenserefreshsucceed":           HardwareInformationDeviceLicensingStatus_DeviceLicenseRefreshSucceed,
		"devicelicenseremovefailed":             HardwareInformationDeviceLicensingStatus_DeviceLicenseRemoveFailed,
		"devicelicenseremovesucceed":            HardwareInformationDeviceLicensingStatus_DeviceLicenseRemoveSucceed,
		"licenserefreshpending":                 HardwareInformationDeviceLicensingStatus_LicenseRefreshPending,
		"licenserefreshstarted":                 HardwareInformationDeviceLicensingStatus_LicenseRefreshStarted,
		"microsoftaccountverificationfailed":    HardwareInformationDeviceLicensingStatus_MicrosoftAccountVerificationFailed,
		"refreshingdevicelicense":               HardwareInformationDeviceLicensingStatus_RefreshingDeviceLicense,
		"removingdevicelicense":                 HardwareInformationDeviceLicensingStatus_RemovingDeviceLicense,
		"unknown":                               HardwareInformationDeviceLicensingStatus_Unknown,
		"verifyingmicrosoftaccountidentity":     HardwareInformationDeviceLicensingStatus_VerifyingMicrosoftAccountIdentity,
		"verifyingmicrosoftdeviceidentity":      HardwareInformationDeviceLicensingStatus_VerifyingMicrosoftDeviceIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareInformationDeviceLicensingStatus(input)
	return &out, nil
}
