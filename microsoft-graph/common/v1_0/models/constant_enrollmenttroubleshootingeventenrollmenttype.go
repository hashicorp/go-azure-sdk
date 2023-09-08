package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentTroubleshootingEventEnrollmentType string

const (
	EnrollmentTroubleshootingEventEnrollmentTypeappleBulkWithUser                     EnrollmentTroubleshootingEventEnrollmentType = "AppleBulkWithUser"
	EnrollmentTroubleshootingEventEnrollmentTypeappleBulkWithoutUser                  EnrollmentTroubleshootingEventEnrollmentType = "AppleBulkWithoutUser"
	EnrollmentTroubleshootingEventEnrollmentTypeappleUserEnrollment                   EnrollmentTroubleshootingEventEnrollmentType = "AppleUserEnrollment"
	EnrollmentTroubleshootingEventEnrollmentTypeappleUserEnrollmentWithServiceAccount EnrollmentTroubleshootingEventEnrollmentType = "AppleUserEnrollmentWithServiceAccount"
	EnrollmentTroubleshootingEventEnrollmentTypedeviceEnrollmentManager               EnrollmentTroubleshootingEventEnrollmentType = "DeviceEnrollmentManager"
	EnrollmentTroubleshootingEventEnrollmentTypeunknown                               EnrollmentTroubleshootingEventEnrollmentType = "Unknown"
	EnrollmentTroubleshootingEventEnrollmentTypeuserEnrollment                        EnrollmentTroubleshootingEventEnrollmentType = "UserEnrollment"
	EnrollmentTroubleshootingEventEnrollmentTypewindowsAutoEnrollment                 EnrollmentTroubleshootingEventEnrollmentType = "WindowsAutoEnrollment"
	EnrollmentTroubleshootingEventEnrollmentTypewindowsAzureADJoin                    EnrollmentTroubleshootingEventEnrollmentType = "WindowsAzureADJoin"
	EnrollmentTroubleshootingEventEnrollmentTypewindowsAzureADJoinUsingDeviceAuth     EnrollmentTroubleshootingEventEnrollmentType = "WindowsAzureADJoinUsingDeviceAuth"
	EnrollmentTroubleshootingEventEnrollmentTypewindowsBulkAzureDomainJoin            EnrollmentTroubleshootingEventEnrollmentType = "WindowsBulkAzureDomainJoin"
	EnrollmentTroubleshootingEventEnrollmentTypewindowsBulkUserless                   EnrollmentTroubleshootingEventEnrollmentType = "WindowsBulkUserless"
	EnrollmentTroubleshootingEventEnrollmentTypewindowsCoManagement                   EnrollmentTroubleshootingEventEnrollmentType = "WindowsCoManagement"
)

func PossibleValuesForEnrollmentTroubleshootingEventEnrollmentType() []string {
	return []string{
		string(EnrollmentTroubleshootingEventEnrollmentTypeappleBulkWithUser),
		string(EnrollmentTroubleshootingEventEnrollmentTypeappleBulkWithoutUser),
		string(EnrollmentTroubleshootingEventEnrollmentTypeappleUserEnrollment),
		string(EnrollmentTroubleshootingEventEnrollmentTypeappleUserEnrollmentWithServiceAccount),
		string(EnrollmentTroubleshootingEventEnrollmentTypedeviceEnrollmentManager),
		string(EnrollmentTroubleshootingEventEnrollmentTypeunknown),
		string(EnrollmentTroubleshootingEventEnrollmentTypeuserEnrollment),
		string(EnrollmentTroubleshootingEventEnrollmentTypewindowsAutoEnrollment),
		string(EnrollmentTroubleshootingEventEnrollmentTypewindowsAzureADJoin),
		string(EnrollmentTroubleshootingEventEnrollmentTypewindowsAzureADJoinUsingDeviceAuth),
		string(EnrollmentTroubleshootingEventEnrollmentTypewindowsBulkAzureDomainJoin),
		string(EnrollmentTroubleshootingEventEnrollmentTypewindowsBulkUserless),
		string(EnrollmentTroubleshootingEventEnrollmentTypewindowsCoManagement),
	}
}

func parseEnrollmentTroubleshootingEventEnrollmentType(input string) (*EnrollmentTroubleshootingEventEnrollmentType, error) {
	vals := map[string]EnrollmentTroubleshootingEventEnrollmentType{
		"applebulkwithuser":                     EnrollmentTroubleshootingEventEnrollmentTypeappleBulkWithUser,
		"applebulkwithoutuser":                  EnrollmentTroubleshootingEventEnrollmentTypeappleBulkWithoutUser,
		"appleuserenrollment":                   EnrollmentTroubleshootingEventEnrollmentTypeappleUserEnrollment,
		"appleuserenrollmentwithserviceaccount": EnrollmentTroubleshootingEventEnrollmentTypeappleUserEnrollmentWithServiceAccount,
		"deviceenrollmentmanager":               EnrollmentTroubleshootingEventEnrollmentTypedeviceEnrollmentManager,
		"unknown":                               EnrollmentTroubleshootingEventEnrollmentTypeunknown,
		"userenrollment":                        EnrollmentTroubleshootingEventEnrollmentTypeuserEnrollment,
		"windowsautoenrollment":                 EnrollmentTroubleshootingEventEnrollmentTypewindowsAutoEnrollment,
		"windowsazureadjoin":                    EnrollmentTroubleshootingEventEnrollmentTypewindowsAzureADJoin,
		"windowsazureadjoinusingdeviceauth":     EnrollmentTroubleshootingEventEnrollmentTypewindowsAzureADJoinUsingDeviceAuth,
		"windowsbulkazuredomainjoin":            EnrollmentTroubleshootingEventEnrollmentTypewindowsBulkAzureDomainJoin,
		"windowsbulkuserless":                   EnrollmentTroubleshootingEventEnrollmentTypewindowsBulkUserless,
		"windowscomanagement":                   EnrollmentTroubleshootingEventEnrollmentTypewindowsCoManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentTroubleshootingEventEnrollmentType(input)
	return &out, nil
}
