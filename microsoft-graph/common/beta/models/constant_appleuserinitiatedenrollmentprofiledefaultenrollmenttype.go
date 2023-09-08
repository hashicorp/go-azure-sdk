package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType string

const (
	AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypeaccountDrivenUserEnrollment AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType = "AccountDrivenUserEnrollment"
	AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypedevice                      AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType = "Device"
	AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypeunknown                     AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType = "Unknown"
	AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypeuser                        AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType = "User"
	AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypewebDeviceEnrollment         AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType = "WebDeviceEnrollment"
)

func PossibleValuesForAppleUserInitiatedEnrollmentProfileDefaultEnrollmentType() []string {
	return []string{
		string(AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypeaccountDrivenUserEnrollment),
		string(AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypedevice),
		string(AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypeunknown),
		string(AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypeuser),
		string(AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypewebDeviceEnrollment),
	}
}

func parseAppleUserInitiatedEnrollmentProfileDefaultEnrollmentType(input string) (*AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType, error) {
	vals := map[string]AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType{
		"accountdrivenuserenrollment": AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypeaccountDrivenUserEnrollment,
		"device":                      AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypedevice,
		"unknown":                     AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypeunknown,
		"user":                        AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypeuser,
		"webdeviceenrollment":         AppleUserInitiatedEnrollmentProfileDefaultEnrollmentTypewebDeviceEnrollment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType(input)
	return &out, nil
}
