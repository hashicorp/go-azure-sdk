package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentTroubleshootingEventEnrollmentType string

const (
	EnrollmentTroubleshootingEventEnrollmentType_AndroidEnterpriseCorporateWorkProfile EnrollmentTroubleshootingEventEnrollmentType = "androidEnterpriseCorporateWorkProfile"
	EnrollmentTroubleshootingEventEnrollmentType_AndroidEnterpriseDedicatedDevice      EnrollmentTroubleshootingEventEnrollmentType = "androidEnterpriseDedicatedDevice"
	EnrollmentTroubleshootingEventEnrollmentType_AndroidEnterpriseFullyManaged         EnrollmentTroubleshootingEventEnrollmentType = "androidEnterpriseFullyManaged"
	EnrollmentTroubleshootingEventEnrollmentType_AppleBulkWithUser                     EnrollmentTroubleshootingEventEnrollmentType = "appleBulkWithUser"
	EnrollmentTroubleshootingEventEnrollmentType_AppleBulkWithoutUser                  EnrollmentTroubleshootingEventEnrollmentType = "appleBulkWithoutUser"
	EnrollmentTroubleshootingEventEnrollmentType_AppleUserEnrollment                   EnrollmentTroubleshootingEventEnrollmentType = "appleUserEnrollment"
	EnrollmentTroubleshootingEventEnrollmentType_AppleUserEnrollmentWithServiceAccount EnrollmentTroubleshootingEventEnrollmentType = "appleUserEnrollmentWithServiceAccount"
	EnrollmentTroubleshootingEventEnrollmentType_AzureAdJoinUsingAzureVmExtension      EnrollmentTroubleshootingEventEnrollmentType = "azureAdJoinUsingAzureVmExtension"
	EnrollmentTroubleshootingEventEnrollmentType_DeviceEnrollmentManager               EnrollmentTroubleshootingEventEnrollmentType = "deviceEnrollmentManager"
	EnrollmentTroubleshootingEventEnrollmentType_Unknown                               EnrollmentTroubleshootingEventEnrollmentType = "unknown"
	EnrollmentTroubleshootingEventEnrollmentType_UserEnrollment                        EnrollmentTroubleshootingEventEnrollmentType = "userEnrollment"
	EnrollmentTroubleshootingEventEnrollmentType_WindowsAutoEnrollment                 EnrollmentTroubleshootingEventEnrollmentType = "windowsAutoEnrollment"
	EnrollmentTroubleshootingEventEnrollmentType_WindowsAzureADJoin                    EnrollmentTroubleshootingEventEnrollmentType = "windowsAzureADJoin"
	EnrollmentTroubleshootingEventEnrollmentType_WindowsAzureADJoinUsingDeviceAuth     EnrollmentTroubleshootingEventEnrollmentType = "windowsAzureADJoinUsingDeviceAuth"
	EnrollmentTroubleshootingEventEnrollmentType_WindowsBulkAzureDomainJoin            EnrollmentTroubleshootingEventEnrollmentType = "windowsBulkAzureDomainJoin"
	EnrollmentTroubleshootingEventEnrollmentType_WindowsBulkUserless                   EnrollmentTroubleshootingEventEnrollmentType = "windowsBulkUserless"
	EnrollmentTroubleshootingEventEnrollmentType_WindowsCoManagement                   EnrollmentTroubleshootingEventEnrollmentType = "windowsCoManagement"
)

func PossibleValuesForEnrollmentTroubleshootingEventEnrollmentType() []string {
	return []string{
		string(EnrollmentTroubleshootingEventEnrollmentType_AndroidEnterpriseCorporateWorkProfile),
		string(EnrollmentTroubleshootingEventEnrollmentType_AndroidEnterpriseDedicatedDevice),
		string(EnrollmentTroubleshootingEventEnrollmentType_AndroidEnterpriseFullyManaged),
		string(EnrollmentTroubleshootingEventEnrollmentType_AppleBulkWithUser),
		string(EnrollmentTroubleshootingEventEnrollmentType_AppleBulkWithoutUser),
		string(EnrollmentTroubleshootingEventEnrollmentType_AppleUserEnrollment),
		string(EnrollmentTroubleshootingEventEnrollmentType_AppleUserEnrollmentWithServiceAccount),
		string(EnrollmentTroubleshootingEventEnrollmentType_AzureAdJoinUsingAzureVmExtension),
		string(EnrollmentTroubleshootingEventEnrollmentType_DeviceEnrollmentManager),
		string(EnrollmentTroubleshootingEventEnrollmentType_Unknown),
		string(EnrollmentTroubleshootingEventEnrollmentType_UserEnrollment),
		string(EnrollmentTroubleshootingEventEnrollmentType_WindowsAutoEnrollment),
		string(EnrollmentTroubleshootingEventEnrollmentType_WindowsAzureADJoin),
		string(EnrollmentTroubleshootingEventEnrollmentType_WindowsAzureADJoinUsingDeviceAuth),
		string(EnrollmentTroubleshootingEventEnrollmentType_WindowsBulkAzureDomainJoin),
		string(EnrollmentTroubleshootingEventEnrollmentType_WindowsBulkUserless),
		string(EnrollmentTroubleshootingEventEnrollmentType_WindowsCoManagement),
	}
}

func (s *EnrollmentTroubleshootingEventEnrollmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentTroubleshootingEventEnrollmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentTroubleshootingEventEnrollmentType(input string) (*EnrollmentTroubleshootingEventEnrollmentType, error) {
	vals := map[string]EnrollmentTroubleshootingEventEnrollmentType{
		"androidenterprisecorporateworkprofile": EnrollmentTroubleshootingEventEnrollmentType_AndroidEnterpriseCorporateWorkProfile,
		"androidenterprisededicateddevice":      EnrollmentTroubleshootingEventEnrollmentType_AndroidEnterpriseDedicatedDevice,
		"androidenterprisefullymanaged":         EnrollmentTroubleshootingEventEnrollmentType_AndroidEnterpriseFullyManaged,
		"applebulkwithuser":                     EnrollmentTroubleshootingEventEnrollmentType_AppleBulkWithUser,
		"applebulkwithoutuser":                  EnrollmentTroubleshootingEventEnrollmentType_AppleBulkWithoutUser,
		"appleuserenrollment":                   EnrollmentTroubleshootingEventEnrollmentType_AppleUserEnrollment,
		"appleuserenrollmentwithserviceaccount": EnrollmentTroubleshootingEventEnrollmentType_AppleUserEnrollmentWithServiceAccount,
		"azureadjoinusingazurevmextension":      EnrollmentTroubleshootingEventEnrollmentType_AzureAdJoinUsingAzureVmExtension,
		"deviceenrollmentmanager":               EnrollmentTroubleshootingEventEnrollmentType_DeviceEnrollmentManager,
		"unknown":                               EnrollmentTroubleshootingEventEnrollmentType_Unknown,
		"userenrollment":                        EnrollmentTroubleshootingEventEnrollmentType_UserEnrollment,
		"windowsautoenrollment":                 EnrollmentTroubleshootingEventEnrollmentType_WindowsAutoEnrollment,
		"windowsazureadjoin":                    EnrollmentTroubleshootingEventEnrollmentType_WindowsAzureADJoin,
		"windowsazureadjoinusingdeviceauth":     EnrollmentTroubleshootingEventEnrollmentType_WindowsAzureADJoinUsingDeviceAuth,
		"windowsbulkazuredomainjoin":            EnrollmentTroubleshootingEventEnrollmentType_WindowsBulkAzureDomainJoin,
		"windowsbulkuserless":                   EnrollmentTroubleshootingEventEnrollmentType_WindowsBulkUserless,
		"windowscomanagement":                   EnrollmentTroubleshootingEventEnrollmentType_WindowsCoManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentTroubleshootingEventEnrollmentType(input)
	return &out, nil
}
