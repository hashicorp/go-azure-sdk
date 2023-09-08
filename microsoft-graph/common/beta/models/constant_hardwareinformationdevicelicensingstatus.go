package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareInformationDeviceLicensingStatus string

const (
	HardwareInformationDeviceLicensingStatusacquiringDeviceLicense                HardwareInformationDeviceLicensingStatus = "AcquiringDeviceLicense"
	HardwareInformationDeviceLicensingStatusdeviceIdentityVerificationFailed      HardwareInformationDeviceLicensingStatus = "DeviceIdentityVerificationFailed"
	HardwareInformationDeviceLicensingStatusdeviceIsNotAzureActiveDirectoryJoined HardwareInformationDeviceLicensingStatus = "DeviceIsNotAzureActiveDirectoryJoined"
	HardwareInformationDeviceLicensingStatusdeviceLicenseRefreshFailed            HardwareInformationDeviceLicensingStatus = "DeviceLicenseRefreshFailed"
	HardwareInformationDeviceLicensingStatusdeviceLicenseRefreshSucceed           HardwareInformationDeviceLicensingStatus = "DeviceLicenseRefreshSucceed"
	HardwareInformationDeviceLicensingStatusdeviceLicenseRemoveFailed             HardwareInformationDeviceLicensingStatus = "DeviceLicenseRemoveFailed"
	HardwareInformationDeviceLicensingStatusdeviceLicenseRemoveSucceed            HardwareInformationDeviceLicensingStatus = "DeviceLicenseRemoveSucceed"
	HardwareInformationDeviceLicensingStatuslicenseRefreshPending                 HardwareInformationDeviceLicensingStatus = "LicenseRefreshPending"
	HardwareInformationDeviceLicensingStatuslicenseRefreshStarted                 HardwareInformationDeviceLicensingStatus = "LicenseRefreshStarted"
	HardwareInformationDeviceLicensingStatusmicrosoftAccountVerificationFailed    HardwareInformationDeviceLicensingStatus = "MicrosoftAccountVerificationFailed"
	HardwareInformationDeviceLicensingStatusrefreshingDeviceLicense               HardwareInformationDeviceLicensingStatus = "RefreshingDeviceLicense"
	HardwareInformationDeviceLicensingStatusremovingDeviceLicense                 HardwareInformationDeviceLicensingStatus = "RemovingDeviceLicense"
	HardwareInformationDeviceLicensingStatusunknown                               HardwareInformationDeviceLicensingStatus = "Unknown"
	HardwareInformationDeviceLicensingStatusverifyingMicrosoftAccountIdentity     HardwareInformationDeviceLicensingStatus = "VerifyingMicrosoftAccountIdentity"
	HardwareInformationDeviceLicensingStatusverifyingMicrosoftDeviceIdentity      HardwareInformationDeviceLicensingStatus = "VerifyingMicrosoftDeviceIdentity"
)

func PossibleValuesForHardwareInformationDeviceLicensingStatus() []string {
	return []string{
		string(HardwareInformationDeviceLicensingStatusacquiringDeviceLicense),
		string(HardwareInformationDeviceLicensingStatusdeviceIdentityVerificationFailed),
		string(HardwareInformationDeviceLicensingStatusdeviceIsNotAzureActiveDirectoryJoined),
		string(HardwareInformationDeviceLicensingStatusdeviceLicenseRefreshFailed),
		string(HardwareInformationDeviceLicensingStatusdeviceLicenseRefreshSucceed),
		string(HardwareInformationDeviceLicensingStatusdeviceLicenseRemoveFailed),
		string(HardwareInformationDeviceLicensingStatusdeviceLicenseRemoveSucceed),
		string(HardwareInformationDeviceLicensingStatuslicenseRefreshPending),
		string(HardwareInformationDeviceLicensingStatuslicenseRefreshStarted),
		string(HardwareInformationDeviceLicensingStatusmicrosoftAccountVerificationFailed),
		string(HardwareInformationDeviceLicensingStatusrefreshingDeviceLicense),
		string(HardwareInformationDeviceLicensingStatusremovingDeviceLicense),
		string(HardwareInformationDeviceLicensingStatusunknown),
		string(HardwareInformationDeviceLicensingStatusverifyingMicrosoftAccountIdentity),
		string(HardwareInformationDeviceLicensingStatusverifyingMicrosoftDeviceIdentity),
	}
}

func parseHardwareInformationDeviceLicensingStatus(input string) (*HardwareInformationDeviceLicensingStatus, error) {
	vals := map[string]HardwareInformationDeviceLicensingStatus{
		"acquiringdevicelicense":                HardwareInformationDeviceLicensingStatusacquiringDeviceLicense,
		"deviceidentityverificationfailed":      HardwareInformationDeviceLicensingStatusdeviceIdentityVerificationFailed,
		"deviceisnotazureactivedirectoryjoined": HardwareInformationDeviceLicensingStatusdeviceIsNotAzureActiveDirectoryJoined,
		"devicelicenserefreshfailed":            HardwareInformationDeviceLicensingStatusdeviceLicenseRefreshFailed,
		"devicelicenserefreshsucceed":           HardwareInformationDeviceLicensingStatusdeviceLicenseRefreshSucceed,
		"devicelicenseremovefailed":             HardwareInformationDeviceLicensingStatusdeviceLicenseRemoveFailed,
		"devicelicenseremovesucceed":            HardwareInformationDeviceLicensingStatusdeviceLicenseRemoveSucceed,
		"licenserefreshpending":                 HardwareInformationDeviceLicensingStatuslicenseRefreshPending,
		"licenserefreshstarted":                 HardwareInformationDeviceLicensingStatuslicenseRefreshStarted,
		"microsoftaccountverificationfailed":    HardwareInformationDeviceLicensingStatusmicrosoftAccountVerificationFailed,
		"refreshingdevicelicense":               HardwareInformationDeviceLicensingStatusrefreshingDeviceLicense,
		"removingdevicelicense":                 HardwareInformationDeviceLicensingStatusremovingDeviceLicense,
		"unknown":                               HardwareInformationDeviceLicensingStatusunknown,
		"verifyingmicrosoftaccountidentity":     HardwareInformationDeviceLicensingStatusverifyingMicrosoftAccountIdentity,
		"verifyingmicrosoftdeviceidentity":      HardwareInformationDeviceLicensingStatusverifyingMicrosoftDeviceIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareInformationDeviceLicensingStatus(input)
	return &out, nil
}
