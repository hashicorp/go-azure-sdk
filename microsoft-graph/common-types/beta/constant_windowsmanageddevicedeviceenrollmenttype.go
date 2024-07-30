package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceDeviceEnrollmentType string

const (
	WindowsManagedDeviceDeviceEnrollmentType_AndroidEnterpriseCorporateWorkProfile WindowsManagedDeviceDeviceEnrollmentType = "androidEnterpriseCorporateWorkProfile"
	WindowsManagedDeviceDeviceEnrollmentType_AndroidEnterpriseDedicatedDevice      WindowsManagedDeviceDeviceEnrollmentType = "androidEnterpriseDedicatedDevice"
	WindowsManagedDeviceDeviceEnrollmentType_AndroidEnterpriseFullyManaged         WindowsManagedDeviceDeviceEnrollmentType = "androidEnterpriseFullyManaged"
	WindowsManagedDeviceDeviceEnrollmentType_AppleBulkWithUser                     WindowsManagedDeviceDeviceEnrollmentType = "appleBulkWithUser"
	WindowsManagedDeviceDeviceEnrollmentType_AppleBulkWithoutUser                  WindowsManagedDeviceDeviceEnrollmentType = "appleBulkWithoutUser"
	WindowsManagedDeviceDeviceEnrollmentType_AppleUserEnrollment                   WindowsManagedDeviceDeviceEnrollmentType = "appleUserEnrollment"
	WindowsManagedDeviceDeviceEnrollmentType_AppleUserEnrollmentWithServiceAccount WindowsManagedDeviceDeviceEnrollmentType = "appleUserEnrollmentWithServiceAccount"
	WindowsManagedDeviceDeviceEnrollmentType_AzureAdJoinUsingAzureVmExtension      WindowsManagedDeviceDeviceEnrollmentType = "azureAdJoinUsingAzureVmExtension"
	WindowsManagedDeviceDeviceEnrollmentType_DeviceEnrollmentManager               WindowsManagedDeviceDeviceEnrollmentType = "deviceEnrollmentManager"
	WindowsManagedDeviceDeviceEnrollmentType_Unknown                               WindowsManagedDeviceDeviceEnrollmentType = "unknown"
	WindowsManagedDeviceDeviceEnrollmentType_UserEnrollment                        WindowsManagedDeviceDeviceEnrollmentType = "userEnrollment"
	WindowsManagedDeviceDeviceEnrollmentType_WindowsAutoEnrollment                 WindowsManagedDeviceDeviceEnrollmentType = "windowsAutoEnrollment"
	WindowsManagedDeviceDeviceEnrollmentType_WindowsAzureADJoin                    WindowsManagedDeviceDeviceEnrollmentType = "windowsAzureADJoin"
	WindowsManagedDeviceDeviceEnrollmentType_WindowsAzureADJoinUsingDeviceAuth     WindowsManagedDeviceDeviceEnrollmentType = "windowsAzureADJoinUsingDeviceAuth"
	WindowsManagedDeviceDeviceEnrollmentType_WindowsBulkAzureDomainJoin            WindowsManagedDeviceDeviceEnrollmentType = "windowsBulkAzureDomainJoin"
	WindowsManagedDeviceDeviceEnrollmentType_WindowsBulkUserless                   WindowsManagedDeviceDeviceEnrollmentType = "windowsBulkUserless"
	WindowsManagedDeviceDeviceEnrollmentType_WindowsCoManagement                   WindowsManagedDeviceDeviceEnrollmentType = "windowsCoManagement"
)

func PossibleValuesForWindowsManagedDeviceDeviceEnrollmentType() []string {
	return []string{
		string(WindowsManagedDeviceDeviceEnrollmentType_AndroidEnterpriseCorporateWorkProfile),
		string(WindowsManagedDeviceDeviceEnrollmentType_AndroidEnterpriseDedicatedDevice),
		string(WindowsManagedDeviceDeviceEnrollmentType_AndroidEnterpriseFullyManaged),
		string(WindowsManagedDeviceDeviceEnrollmentType_AppleBulkWithUser),
		string(WindowsManagedDeviceDeviceEnrollmentType_AppleBulkWithoutUser),
		string(WindowsManagedDeviceDeviceEnrollmentType_AppleUserEnrollment),
		string(WindowsManagedDeviceDeviceEnrollmentType_AppleUserEnrollmentWithServiceAccount),
		string(WindowsManagedDeviceDeviceEnrollmentType_AzureAdJoinUsingAzureVmExtension),
		string(WindowsManagedDeviceDeviceEnrollmentType_DeviceEnrollmentManager),
		string(WindowsManagedDeviceDeviceEnrollmentType_Unknown),
		string(WindowsManagedDeviceDeviceEnrollmentType_UserEnrollment),
		string(WindowsManagedDeviceDeviceEnrollmentType_WindowsAutoEnrollment),
		string(WindowsManagedDeviceDeviceEnrollmentType_WindowsAzureADJoin),
		string(WindowsManagedDeviceDeviceEnrollmentType_WindowsAzureADJoinUsingDeviceAuth),
		string(WindowsManagedDeviceDeviceEnrollmentType_WindowsBulkAzureDomainJoin),
		string(WindowsManagedDeviceDeviceEnrollmentType_WindowsBulkUserless),
		string(WindowsManagedDeviceDeviceEnrollmentType_WindowsCoManagement),
	}
}

func (s *WindowsManagedDeviceDeviceEnrollmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceDeviceEnrollmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceDeviceEnrollmentType(input string) (*WindowsManagedDeviceDeviceEnrollmentType, error) {
	vals := map[string]WindowsManagedDeviceDeviceEnrollmentType{
		"androidenterprisecorporateworkprofile": WindowsManagedDeviceDeviceEnrollmentType_AndroidEnterpriseCorporateWorkProfile,
		"androidenterprisededicateddevice":      WindowsManagedDeviceDeviceEnrollmentType_AndroidEnterpriseDedicatedDevice,
		"androidenterprisefullymanaged":         WindowsManagedDeviceDeviceEnrollmentType_AndroidEnterpriseFullyManaged,
		"applebulkwithuser":                     WindowsManagedDeviceDeviceEnrollmentType_AppleBulkWithUser,
		"applebulkwithoutuser":                  WindowsManagedDeviceDeviceEnrollmentType_AppleBulkWithoutUser,
		"appleuserenrollment":                   WindowsManagedDeviceDeviceEnrollmentType_AppleUserEnrollment,
		"appleuserenrollmentwithserviceaccount": WindowsManagedDeviceDeviceEnrollmentType_AppleUserEnrollmentWithServiceAccount,
		"azureadjoinusingazurevmextension":      WindowsManagedDeviceDeviceEnrollmentType_AzureAdJoinUsingAzureVmExtension,
		"deviceenrollmentmanager":               WindowsManagedDeviceDeviceEnrollmentType_DeviceEnrollmentManager,
		"unknown":                               WindowsManagedDeviceDeviceEnrollmentType_Unknown,
		"userenrollment":                        WindowsManagedDeviceDeviceEnrollmentType_UserEnrollment,
		"windowsautoenrollment":                 WindowsManagedDeviceDeviceEnrollmentType_WindowsAutoEnrollment,
		"windowsazureadjoin":                    WindowsManagedDeviceDeviceEnrollmentType_WindowsAzureADJoin,
		"windowsazureadjoinusingdeviceauth":     WindowsManagedDeviceDeviceEnrollmentType_WindowsAzureADJoinUsingDeviceAuth,
		"windowsbulkazuredomainjoin":            WindowsManagedDeviceDeviceEnrollmentType_WindowsBulkAzureDomainJoin,
		"windowsbulkuserless":                   WindowsManagedDeviceDeviceEnrollmentType_WindowsBulkUserless,
		"windowscomanagement":                   WindowsManagedDeviceDeviceEnrollmentType_WindowsCoManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceDeviceEnrollmentType(input)
	return &out, nil
}
