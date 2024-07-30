package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceDeviceEnrollmentType string

const (
	ManagedDeviceDeviceEnrollmentType_AndroidEnterpriseCorporateWorkProfile ManagedDeviceDeviceEnrollmentType = "androidEnterpriseCorporateWorkProfile"
	ManagedDeviceDeviceEnrollmentType_AndroidEnterpriseDedicatedDevice      ManagedDeviceDeviceEnrollmentType = "androidEnterpriseDedicatedDevice"
	ManagedDeviceDeviceEnrollmentType_AndroidEnterpriseFullyManaged         ManagedDeviceDeviceEnrollmentType = "androidEnterpriseFullyManaged"
	ManagedDeviceDeviceEnrollmentType_AppleBulkWithUser                     ManagedDeviceDeviceEnrollmentType = "appleBulkWithUser"
	ManagedDeviceDeviceEnrollmentType_AppleBulkWithoutUser                  ManagedDeviceDeviceEnrollmentType = "appleBulkWithoutUser"
	ManagedDeviceDeviceEnrollmentType_AppleUserEnrollment                   ManagedDeviceDeviceEnrollmentType = "appleUserEnrollment"
	ManagedDeviceDeviceEnrollmentType_AppleUserEnrollmentWithServiceAccount ManagedDeviceDeviceEnrollmentType = "appleUserEnrollmentWithServiceAccount"
	ManagedDeviceDeviceEnrollmentType_AzureAdJoinUsingAzureVmExtension      ManagedDeviceDeviceEnrollmentType = "azureAdJoinUsingAzureVmExtension"
	ManagedDeviceDeviceEnrollmentType_DeviceEnrollmentManager               ManagedDeviceDeviceEnrollmentType = "deviceEnrollmentManager"
	ManagedDeviceDeviceEnrollmentType_Unknown                               ManagedDeviceDeviceEnrollmentType = "unknown"
	ManagedDeviceDeviceEnrollmentType_UserEnrollment                        ManagedDeviceDeviceEnrollmentType = "userEnrollment"
	ManagedDeviceDeviceEnrollmentType_WindowsAutoEnrollment                 ManagedDeviceDeviceEnrollmentType = "windowsAutoEnrollment"
	ManagedDeviceDeviceEnrollmentType_WindowsAzureADJoin                    ManagedDeviceDeviceEnrollmentType = "windowsAzureADJoin"
	ManagedDeviceDeviceEnrollmentType_WindowsAzureADJoinUsingDeviceAuth     ManagedDeviceDeviceEnrollmentType = "windowsAzureADJoinUsingDeviceAuth"
	ManagedDeviceDeviceEnrollmentType_WindowsBulkAzureDomainJoin            ManagedDeviceDeviceEnrollmentType = "windowsBulkAzureDomainJoin"
	ManagedDeviceDeviceEnrollmentType_WindowsBulkUserless                   ManagedDeviceDeviceEnrollmentType = "windowsBulkUserless"
	ManagedDeviceDeviceEnrollmentType_WindowsCoManagement                   ManagedDeviceDeviceEnrollmentType = "windowsCoManagement"
)

func PossibleValuesForManagedDeviceDeviceEnrollmentType() []string {
	return []string{
		string(ManagedDeviceDeviceEnrollmentType_AndroidEnterpriseCorporateWorkProfile),
		string(ManagedDeviceDeviceEnrollmentType_AndroidEnterpriseDedicatedDevice),
		string(ManagedDeviceDeviceEnrollmentType_AndroidEnterpriseFullyManaged),
		string(ManagedDeviceDeviceEnrollmentType_AppleBulkWithUser),
		string(ManagedDeviceDeviceEnrollmentType_AppleBulkWithoutUser),
		string(ManagedDeviceDeviceEnrollmentType_AppleUserEnrollment),
		string(ManagedDeviceDeviceEnrollmentType_AppleUserEnrollmentWithServiceAccount),
		string(ManagedDeviceDeviceEnrollmentType_AzureAdJoinUsingAzureVmExtension),
		string(ManagedDeviceDeviceEnrollmentType_DeviceEnrollmentManager),
		string(ManagedDeviceDeviceEnrollmentType_Unknown),
		string(ManagedDeviceDeviceEnrollmentType_UserEnrollment),
		string(ManagedDeviceDeviceEnrollmentType_WindowsAutoEnrollment),
		string(ManagedDeviceDeviceEnrollmentType_WindowsAzureADJoin),
		string(ManagedDeviceDeviceEnrollmentType_WindowsAzureADJoinUsingDeviceAuth),
		string(ManagedDeviceDeviceEnrollmentType_WindowsBulkAzureDomainJoin),
		string(ManagedDeviceDeviceEnrollmentType_WindowsBulkUserless),
		string(ManagedDeviceDeviceEnrollmentType_WindowsCoManagement),
	}
}

func (s *ManagedDeviceDeviceEnrollmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceDeviceEnrollmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceDeviceEnrollmentType(input string) (*ManagedDeviceDeviceEnrollmentType, error) {
	vals := map[string]ManagedDeviceDeviceEnrollmentType{
		"androidenterprisecorporateworkprofile": ManagedDeviceDeviceEnrollmentType_AndroidEnterpriseCorporateWorkProfile,
		"androidenterprisededicateddevice":      ManagedDeviceDeviceEnrollmentType_AndroidEnterpriseDedicatedDevice,
		"androidenterprisefullymanaged":         ManagedDeviceDeviceEnrollmentType_AndroidEnterpriseFullyManaged,
		"applebulkwithuser":                     ManagedDeviceDeviceEnrollmentType_AppleBulkWithUser,
		"applebulkwithoutuser":                  ManagedDeviceDeviceEnrollmentType_AppleBulkWithoutUser,
		"appleuserenrollment":                   ManagedDeviceDeviceEnrollmentType_AppleUserEnrollment,
		"appleuserenrollmentwithserviceaccount": ManagedDeviceDeviceEnrollmentType_AppleUserEnrollmentWithServiceAccount,
		"azureadjoinusingazurevmextension":      ManagedDeviceDeviceEnrollmentType_AzureAdJoinUsingAzureVmExtension,
		"deviceenrollmentmanager":               ManagedDeviceDeviceEnrollmentType_DeviceEnrollmentManager,
		"unknown":                               ManagedDeviceDeviceEnrollmentType_Unknown,
		"userenrollment":                        ManagedDeviceDeviceEnrollmentType_UserEnrollment,
		"windowsautoenrollment":                 ManagedDeviceDeviceEnrollmentType_WindowsAutoEnrollment,
		"windowsazureadjoin":                    ManagedDeviceDeviceEnrollmentType_WindowsAzureADJoin,
		"windowsazureadjoinusingdeviceauth":     ManagedDeviceDeviceEnrollmentType_WindowsAzureADJoinUsingDeviceAuth,
		"windowsbulkazuredomainjoin":            ManagedDeviceDeviceEnrollmentType_WindowsBulkAzureDomainJoin,
		"windowsbulkuserless":                   ManagedDeviceDeviceEnrollmentType_WindowsBulkUserless,
		"windowscomanagement":                   ManagedDeviceDeviceEnrollmentType_WindowsCoManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceDeviceEnrollmentType(input)
	return &out, nil
}
