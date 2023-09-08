package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceDeviceEnrollmentType string

const (
	ManagedDeviceDeviceEnrollmentTypeappleBulkWithUser                     ManagedDeviceDeviceEnrollmentType = "AppleBulkWithUser"
	ManagedDeviceDeviceEnrollmentTypeappleBulkWithoutUser                  ManagedDeviceDeviceEnrollmentType = "AppleBulkWithoutUser"
	ManagedDeviceDeviceEnrollmentTypeappleUserEnrollment                   ManagedDeviceDeviceEnrollmentType = "AppleUserEnrollment"
	ManagedDeviceDeviceEnrollmentTypeappleUserEnrollmentWithServiceAccount ManagedDeviceDeviceEnrollmentType = "AppleUserEnrollmentWithServiceAccount"
	ManagedDeviceDeviceEnrollmentTypedeviceEnrollmentManager               ManagedDeviceDeviceEnrollmentType = "DeviceEnrollmentManager"
	ManagedDeviceDeviceEnrollmentTypeunknown                               ManagedDeviceDeviceEnrollmentType = "Unknown"
	ManagedDeviceDeviceEnrollmentTypeuserEnrollment                        ManagedDeviceDeviceEnrollmentType = "UserEnrollment"
	ManagedDeviceDeviceEnrollmentTypewindowsAutoEnrollment                 ManagedDeviceDeviceEnrollmentType = "WindowsAutoEnrollment"
	ManagedDeviceDeviceEnrollmentTypewindowsAzureADJoin                    ManagedDeviceDeviceEnrollmentType = "WindowsAzureADJoin"
	ManagedDeviceDeviceEnrollmentTypewindowsAzureADJoinUsingDeviceAuth     ManagedDeviceDeviceEnrollmentType = "WindowsAzureADJoinUsingDeviceAuth"
	ManagedDeviceDeviceEnrollmentTypewindowsBulkAzureDomainJoin            ManagedDeviceDeviceEnrollmentType = "WindowsBulkAzureDomainJoin"
	ManagedDeviceDeviceEnrollmentTypewindowsBulkUserless                   ManagedDeviceDeviceEnrollmentType = "WindowsBulkUserless"
	ManagedDeviceDeviceEnrollmentTypewindowsCoManagement                   ManagedDeviceDeviceEnrollmentType = "WindowsCoManagement"
)

func PossibleValuesForManagedDeviceDeviceEnrollmentType() []string {
	return []string{
		string(ManagedDeviceDeviceEnrollmentTypeappleBulkWithUser),
		string(ManagedDeviceDeviceEnrollmentTypeappleBulkWithoutUser),
		string(ManagedDeviceDeviceEnrollmentTypeappleUserEnrollment),
		string(ManagedDeviceDeviceEnrollmentTypeappleUserEnrollmentWithServiceAccount),
		string(ManagedDeviceDeviceEnrollmentTypedeviceEnrollmentManager),
		string(ManagedDeviceDeviceEnrollmentTypeunknown),
		string(ManagedDeviceDeviceEnrollmentTypeuserEnrollment),
		string(ManagedDeviceDeviceEnrollmentTypewindowsAutoEnrollment),
		string(ManagedDeviceDeviceEnrollmentTypewindowsAzureADJoin),
		string(ManagedDeviceDeviceEnrollmentTypewindowsAzureADJoinUsingDeviceAuth),
		string(ManagedDeviceDeviceEnrollmentTypewindowsBulkAzureDomainJoin),
		string(ManagedDeviceDeviceEnrollmentTypewindowsBulkUserless),
		string(ManagedDeviceDeviceEnrollmentTypewindowsCoManagement),
	}
}

func parseManagedDeviceDeviceEnrollmentType(input string) (*ManagedDeviceDeviceEnrollmentType, error) {
	vals := map[string]ManagedDeviceDeviceEnrollmentType{
		"applebulkwithuser":                     ManagedDeviceDeviceEnrollmentTypeappleBulkWithUser,
		"applebulkwithoutuser":                  ManagedDeviceDeviceEnrollmentTypeappleBulkWithoutUser,
		"appleuserenrollment":                   ManagedDeviceDeviceEnrollmentTypeappleUserEnrollment,
		"appleuserenrollmentwithserviceaccount": ManagedDeviceDeviceEnrollmentTypeappleUserEnrollmentWithServiceAccount,
		"deviceenrollmentmanager":               ManagedDeviceDeviceEnrollmentTypedeviceEnrollmentManager,
		"unknown":                               ManagedDeviceDeviceEnrollmentTypeunknown,
		"userenrollment":                        ManagedDeviceDeviceEnrollmentTypeuserEnrollment,
		"windowsautoenrollment":                 ManagedDeviceDeviceEnrollmentTypewindowsAutoEnrollment,
		"windowsazureadjoin":                    ManagedDeviceDeviceEnrollmentTypewindowsAzureADJoin,
		"windowsazureadjoinusingdeviceauth":     ManagedDeviceDeviceEnrollmentTypewindowsAzureADJoinUsingDeviceAuth,
		"windowsbulkazuredomainjoin":            ManagedDeviceDeviceEnrollmentTypewindowsBulkAzureDomainJoin,
		"windowsbulkuserless":                   ManagedDeviceDeviceEnrollmentTypewindowsBulkUserless,
		"windowscomanagement":                   ManagedDeviceDeviceEnrollmentTypewindowsCoManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceDeviceEnrollmentType(input)
	return &out, nil
}
