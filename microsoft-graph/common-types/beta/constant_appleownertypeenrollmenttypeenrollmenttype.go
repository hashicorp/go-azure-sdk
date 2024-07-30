package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleOwnerTypeEnrollmentTypeEnrollmentType string

const (
	AppleOwnerTypeEnrollmentTypeEnrollmentType_AccountDrivenUserEnrollment AppleOwnerTypeEnrollmentTypeEnrollmentType = "accountDrivenUserEnrollment"
	AppleOwnerTypeEnrollmentTypeEnrollmentType_Device                      AppleOwnerTypeEnrollmentTypeEnrollmentType = "device"
	AppleOwnerTypeEnrollmentTypeEnrollmentType_Unknown                     AppleOwnerTypeEnrollmentTypeEnrollmentType = "unknown"
	AppleOwnerTypeEnrollmentTypeEnrollmentType_User                        AppleOwnerTypeEnrollmentTypeEnrollmentType = "user"
	AppleOwnerTypeEnrollmentTypeEnrollmentType_WebDeviceEnrollment         AppleOwnerTypeEnrollmentTypeEnrollmentType = "webDeviceEnrollment"
)

func PossibleValuesForAppleOwnerTypeEnrollmentTypeEnrollmentType() []string {
	return []string{
		string(AppleOwnerTypeEnrollmentTypeEnrollmentType_AccountDrivenUserEnrollment),
		string(AppleOwnerTypeEnrollmentTypeEnrollmentType_Device),
		string(AppleOwnerTypeEnrollmentTypeEnrollmentType_Unknown),
		string(AppleOwnerTypeEnrollmentTypeEnrollmentType_User),
		string(AppleOwnerTypeEnrollmentTypeEnrollmentType_WebDeviceEnrollment),
	}
}

func (s *AppleOwnerTypeEnrollmentTypeEnrollmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleOwnerTypeEnrollmentTypeEnrollmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleOwnerTypeEnrollmentTypeEnrollmentType(input string) (*AppleOwnerTypeEnrollmentTypeEnrollmentType, error) {
	vals := map[string]AppleOwnerTypeEnrollmentTypeEnrollmentType{
		"accountdrivenuserenrollment": AppleOwnerTypeEnrollmentTypeEnrollmentType_AccountDrivenUserEnrollment,
		"device":                      AppleOwnerTypeEnrollmentTypeEnrollmentType_Device,
		"unknown":                     AppleOwnerTypeEnrollmentTypeEnrollmentType_Unknown,
		"user":                        AppleOwnerTypeEnrollmentTypeEnrollmentType_User,
		"webdeviceenrollment":         AppleOwnerTypeEnrollmentTypeEnrollmentType_WebDeviceEnrollment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleOwnerTypeEnrollmentTypeEnrollmentType(input)
	return &out, nil
}
