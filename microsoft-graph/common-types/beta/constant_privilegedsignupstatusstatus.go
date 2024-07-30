package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedSignupStatusStatus string

const (
	PrivilegedSignupStatusStatus_Disabled                      PrivilegedSignupStatusStatus = "disabled"
	PrivilegedSignupStatusStatus_NotRegisteredYet              PrivilegedSignupStatusStatus = "notRegisteredYet"
	PrivilegedSignupStatusStatus_RegisteredSetupInProgress     PrivilegedSignupStatusStatus = "registeredSetupInProgress"
	PrivilegedSignupStatusStatus_RegisteredSetupNotStarted     PrivilegedSignupStatusStatus = "registeredSetupNotStarted"
	PrivilegedSignupStatusStatus_RegistrationAndSetupCompleted PrivilegedSignupStatusStatus = "registrationAndSetupCompleted"
	PrivilegedSignupStatusStatus_RegistrationFailed            PrivilegedSignupStatusStatus = "registrationFailed"
	PrivilegedSignupStatusStatus_RegistrationTimedOut          PrivilegedSignupStatusStatus = "registrationTimedOut"
	PrivilegedSignupStatusStatus_Unknown                       PrivilegedSignupStatusStatus = "unknown"
)

func PossibleValuesForPrivilegedSignupStatusStatus() []string {
	return []string{
		string(PrivilegedSignupStatusStatus_Disabled),
		string(PrivilegedSignupStatusStatus_NotRegisteredYet),
		string(PrivilegedSignupStatusStatus_RegisteredSetupInProgress),
		string(PrivilegedSignupStatusStatus_RegisteredSetupNotStarted),
		string(PrivilegedSignupStatusStatus_RegistrationAndSetupCompleted),
		string(PrivilegedSignupStatusStatus_RegistrationFailed),
		string(PrivilegedSignupStatusStatus_RegistrationTimedOut),
		string(PrivilegedSignupStatusStatus_Unknown),
	}
}

func (s *PrivilegedSignupStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedSignupStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedSignupStatusStatus(input string) (*PrivilegedSignupStatusStatus, error) {
	vals := map[string]PrivilegedSignupStatusStatus{
		"disabled":                      PrivilegedSignupStatusStatus_Disabled,
		"notregisteredyet":              PrivilegedSignupStatusStatus_NotRegisteredYet,
		"registeredsetupinprogress":     PrivilegedSignupStatusStatus_RegisteredSetupInProgress,
		"registeredsetupnotstarted":     PrivilegedSignupStatusStatus_RegisteredSetupNotStarted,
		"registrationandsetupcompleted": PrivilegedSignupStatusStatus_RegistrationAndSetupCompleted,
		"registrationfailed":            PrivilegedSignupStatusStatus_RegistrationFailed,
		"registrationtimedout":          PrivilegedSignupStatusStatus_RegistrationTimedOut,
		"unknown":                       PrivilegedSignupStatusStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedSignupStatusStatus(input)
	return &out, nil
}
