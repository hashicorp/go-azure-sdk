package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleOwnerTypeEnrollmentTypeEnrollmentType string

const (
	AppleOwnerTypeEnrollmentTypeEnrollmentTypeaccountDrivenUserEnrollment AppleOwnerTypeEnrollmentTypeEnrollmentType = "AccountDrivenUserEnrollment"
	AppleOwnerTypeEnrollmentTypeEnrollmentTypedevice                      AppleOwnerTypeEnrollmentTypeEnrollmentType = "Device"
	AppleOwnerTypeEnrollmentTypeEnrollmentTypeunknown                     AppleOwnerTypeEnrollmentTypeEnrollmentType = "Unknown"
	AppleOwnerTypeEnrollmentTypeEnrollmentTypeuser                        AppleOwnerTypeEnrollmentTypeEnrollmentType = "User"
	AppleOwnerTypeEnrollmentTypeEnrollmentTypewebDeviceEnrollment         AppleOwnerTypeEnrollmentTypeEnrollmentType = "WebDeviceEnrollment"
)

func PossibleValuesForAppleOwnerTypeEnrollmentTypeEnrollmentType() []string {
	return []string{
		string(AppleOwnerTypeEnrollmentTypeEnrollmentTypeaccountDrivenUserEnrollment),
		string(AppleOwnerTypeEnrollmentTypeEnrollmentTypedevice),
		string(AppleOwnerTypeEnrollmentTypeEnrollmentTypeunknown),
		string(AppleOwnerTypeEnrollmentTypeEnrollmentTypeuser),
		string(AppleOwnerTypeEnrollmentTypeEnrollmentTypewebDeviceEnrollment),
	}
}

func parseAppleOwnerTypeEnrollmentTypeEnrollmentType(input string) (*AppleOwnerTypeEnrollmentTypeEnrollmentType, error) {
	vals := map[string]AppleOwnerTypeEnrollmentTypeEnrollmentType{
		"accountdrivenuserenrollment": AppleOwnerTypeEnrollmentTypeEnrollmentTypeaccountDrivenUserEnrollment,
		"device":                      AppleOwnerTypeEnrollmentTypeEnrollmentTypedevice,
		"unknown":                     AppleOwnerTypeEnrollmentTypeEnrollmentTypeunknown,
		"user":                        AppleOwnerTypeEnrollmentTypeEnrollmentTypeuser,
		"webdeviceenrollment":         AppleOwnerTypeEnrollmentTypeEnrollmentTypewebDeviceEnrollment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleOwnerTypeEnrollmentTypeEnrollmentType(input)
	return &out, nil
}
