package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType string

const (
	AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_AccountDrivenUserEnrollment AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType = "accountDrivenUserEnrollment"
	AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_Device                      AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType = "device"
	AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_Unknown                     AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType = "unknown"
	AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_User                        AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType = "user"
	AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_WebDeviceEnrollment         AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType = "webDeviceEnrollment"
)

func PossibleValuesForAppleUserInitiatedEnrollmentProfileDefaultEnrollmentType() []string {
	return []string{
		string(AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_AccountDrivenUserEnrollment),
		string(AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_Device),
		string(AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_Unknown),
		string(AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_User),
		string(AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_WebDeviceEnrollment),
	}
}

func (s *AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleUserInitiatedEnrollmentProfileDefaultEnrollmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleUserInitiatedEnrollmentProfileDefaultEnrollmentType(input string) (*AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType, error) {
	vals := map[string]AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType{
		"accountdrivenuserenrollment": AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_AccountDrivenUserEnrollment,
		"device":                      AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_Device,
		"unknown":                     AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_Unknown,
		"user":                        AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_User,
		"webdeviceenrollment":         AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType_WebDeviceEnrollment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleUserInitiatedEnrollmentProfileDefaultEnrollmentType(input)
	return &out, nil
}
