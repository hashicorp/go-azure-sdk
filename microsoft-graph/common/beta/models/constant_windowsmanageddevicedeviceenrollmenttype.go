package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceDeviceEnrollmentType string

const (
	WindowsManagedDeviceDeviceEnrollmentTypeandroidEnterpriseCorporateWorkProfile WindowsManagedDeviceDeviceEnrollmentType = "AndroidEnterpriseCorporateWorkProfile"
	WindowsManagedDeviceDeviceEnrollmentTypeandroidEnterpriseDedicatedDevice      WindowsManagedDeviceDeviceEnrollmentType = "AndroidEnterpriseDedicatedDevice"
	WindowsManagedDeviceDeviceEnrollmentTypeandroidEnterpriseFullyManaged         WindowsManagedDeviceDeviceEnrollmentType = "AndroidEnterpriseFullyManaged"
	WindowsManagedDeviceDeviceEnrollmentTypeappleBulkWithUser                     WindowsManagedDeviceDeviceEnrollmentType = "AppleBulkWithUser"
	WindowsManagedDeviceDeviceEnrollmentTypeappleBulkWithoutUser                  WindowsManagedDeviceDeviceEnrollmentType = "AppleBulkWithoutUser"
	WindowsManagedDeviceDeviceEnrollmentTypeappleUserEnrollment                   WindowsManagedDeviceDeviceEnrollmentType = "AppleUserEnrollment"
	WindowsManagedDeviceDeviceEnrollmentTypeappleUserEnrollmentWithServiceAccount WindowsManagedDeviceDeviceEnrollmentType = "AppleUserEnrollmentWithServiceAccount"
	WindowsManagedDeviceDeviceEnrollmentTypeazureAdJoinUsingAzureVmExtension      WindowsManagedDeviceDeviceEnrollmentType = "AzureAdJoinUsingAzureVmExtension"
	WindowsManagedDeviceDeviceEnrollmentTypedeviceEnrollmentManager               WindowsManagedDeviceDeviceEnrollmentType = "DeviceEnrollmentManager"
	WindowsManagedDeviceDeviceEnrollmentTypeunknown                               WindowsManagedDeviceDeviceEnrollmentType = "Unknown"
	WindowsManagedDeviceDeviceEnrollmentTypeuserEnrollment                        WindowsManagedDeviceDeviceEnrollmentType = "UserEnrollment"
	WindowsManagedDeviceDeviceEnrollmentTypewindowsAutoEnrollment                 WindowsManagedDeviceDeviceEnrollmentType = "WindowsAutoEnrollment"
	WindowsManagedDeviceDeviceEnrollmentTypewindowsAzureADJoin                    WindowsManagedDeviceDeviceEnrollmentType = "WindowsAzureADJoin"
	WindowsManagedDeviceDeviceEnrollmentTypewindowsAzureADJoinUsingDeviceAuth     WindowsManagedDeviceDeviceEnrollmentType = "WindowsAzureADJoinUsingDeviceAuth"
	WindowsManagedDeviceDeviceEnrollmentTypewindowsBulkAzureDomainJoin            WindowsManagedDeviceDeviceEnrollmentType = "WindowsBulkAzureDomainJoin"
	WindowsManagedDeviceDeviceEnrollmentTypewindowsBulkUserless                   WindowsManagedDeviceDeviceEnrollmentType = "WindowsBulkUserless"
	WindowsManagedDeviceDeviceEnrollmentTypewindowsCoManagement                   WindowsManagedDeviceDeviceEnrollmentType = "WindowsCoManagement"
)

func PossibleValuesForWindowsManagedDeviceDeviceEnrollmentType() []string {
	return []string{
		string(WindowsManagedDeviceDeviceEnrollmentTypeandroidEnterpriseCorporateWorkProfile),
		string(WindowsManagedDeviceDeviceEnrollmentTypeandroidEnterpriseDedicatedDevice),
		string(WindowsManagedDeviceDeviceEnrollmentTypeandroidEnterpriseFullyManaged),
		string(WindowsManagedDeviceDeviceEnrollmentTypeappleBulkWithUser),
		string(WindowsManagedDeviceDeviceEnrollmentTypeappleBulkWithoutUser),
		string(WindowsManagedDeviceDeviceEnrollmentTypeappleUserEnrollment),
		string(WindowsManagedDeviceDeviceEnrollmentTypeappleUserEnrollmentWithServiceAccount),
		string(WindowsManagedDeviceDeviceEnrollmentTypeazureAdJoinUsingAzureVmExtension),
		string(WindowsManagedDeviceDeviceEnrollmentTypedeviceEnrollmentManager),
		string(WindowsManagedDeviceDeviceEnrollmentTypeunknown),
		string(WindowsManagedDeviceDeviceEnrollmentTypeuserEnrollment),
		string(WindowsManagedDeviceDeviceEnrollmentTypewindowsAutoEnrollment),
		string(WindowsManagedDeviceDeviceEnrollmentTypewindowsAzureADJoin),
		string(WindowsManagedDeviceDeviceEnrollmentTypewindowsAzureADJoinUsingDeviceAuth),
		string(WindowsManagedDeviceDeviceEnrollmentTypewindowsBulkAzureDomainJoin),
		string(WindowsManagedDeviceDeviceEnrollmentTypewindowsBulkUserless),
		string(WindowsManagedDeviceDeviceEnrollmentTypewindowsCoManagement),
	}
}

func parseWindowsManagedDeviceDeviceEnrollmentType(input string) (*WindowsManagedDeviceDeviceEnrollmentType, error) {
	vals := map[string]WindowsManagedDeviceDeviceEnrollmentType{
		"androidenterprisecorporateworkprofile": WindowsManagedDeviceDeviceEnrollmentTypeandroidEnterpriseCorporateWorkProfile,
		"androidenterprisededicateddevice":      WindowsManagedDeviceDeviceEnrollmentTypeandroidEnterpriseDedicatedDevice,
		"androidenterprisefullymanaged":         WindowsManagedDeviceDeviceEnrollmentTypeandroidEnterpriseFullyManaged,
		"applebulkwithuser":                     WindowsManagedDeviceDeviceEnrollmentTypeappleBulkWithUser,
		"applebulkwithoutuser":                  WindowsManagedDeviceDeviceEnrollmentTypeappleBulkWithoutUser,
		"appleuserenrollment":                   WindowsManagedDeviceDeviceEnrollmentTypeappleUserEnrollment,
		"appleuserenrollmentwithserviceaccount": WindowsManagedDeviceDeviceEnrollmentTypeappleUserEnrollmentWithServiceAccount,
		"azureadjoinusingazurevmextension":      WindowsManagedDeviceDeviceEnrollmentTypeazureAdJoinUsingAzureVmExtension,
		"deviceenrollmentmanager":               WindowsManagedDeviceDeviceEnrollmentTypedeviceEnrollmentManager,
		"unknown":                               WindowsManagedDeviceDeviceEnrollmentTypeunknown,
		"userenrollment":                        WindowsManagedDeviceDeviceEnrollmentTypeuserEnrollment,
		"windowsautoenrollment":                 WindowsManagedDeviceDeviceEnrollmentTypewindowsAutoEnrollment,
		"windowsazureadjoin":                    WindowsManagedDeviceDeviceEnrollmentTypewindowsAzureADJoin,
		"windowsazureadjoinusingdeviceauth":     WindowsManagedDeviceDeviceEnrollmentTypewindowsAzureADJoinUsingDeviceAuth,
		"windowsbulkazuredomainjoin":            WindowsManagedDeviceDeviceEnrollmentTypewindowsBulkAzureDomainJoin,
		"windowsbulkuserless":                   WindowsManagedDeviceDeviceEnrollmentTypewindowsBulkUserless,
		"windowscomanagement":                   WindowsManagedDeviceDeviceEnrollmentTypewindowsCoManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceDeviceEnrollmentType(input)
	return &out, nil
}
