package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantSetupInfoSetupStatus string

const (
	TenantSetupInfoSetupStatus_Disabled                      TenantSetupInfoSetupStatus = "disabled"
	TenantSetupInfoSetupStatus_NotRegisteredYet              TenantSetupInfoSetupStatus = "notRegisteredYet"
	TenantSetupInfoSetupStatus_RegisteredSetupInProgress     TenantSetupInfoSetupStatus = "registeredSetupInProgress"
	TenantSetupInfoSetupStatus_RegisteredSetupNotStarted     TenantSetupInfoSetupStatus = "registeredSetupNotStarted"
	TenantSetupInfoSetupStatus_RegistrationAndSetupCompleted TenantSetupInfoSetupStatus = "registrationAndSetupCompleted"
	TenantSetupInfoSetupStatus_RegistrationFailed            TenantSetupInfoSetupStatus = "registrationFailed"
	TenantSetupInfoSetupStatus_RegistrationTimedOut          TenantSetupInfoSetupStatus = "registrationTimedOut"
	TenantSetupInfoSetupStatus_Unknown                       TenantSetupInfoSetupStatus = "unknown"
)

func PossibleValuesForTenantSetupInfoSetupStatus() []string {
	return []string{
		string(TenantSetupInfoSetupStatus_Disabled),
		string(TenantSetupInfoSetupStatus_NotRegisteredYet),
		string(TenantSetupInfoSetupStatus_RegisteredSetupInProgress),
		string(TenantSetupInfoSetupStatus_RegisteredSetupNotStarted),
		string(TenantSetupInfoSetupStatus_RegistrationAndSetupCompleted),
		string(TenantSetupInfoSetupStatus_RegistrationFailed),
		string(TenantSetupInfoSetupStatus_RegistrationTimedOut),
		string(TenantSetupInfoSetupStatus_Unknown),
	}
}

func (s *TenantSetupInfoSetupStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTenantSetupInfoSetupStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTenantSetupInfoSetupStatus(input string) (*TenantSetupInfoSetupStatus, error) {
	vals := map[string]TenantSetupInfoSetupStatus{
		"disabled":                      TenantSetupInfoSetupStatus_Disabled,
		"notregisteredyet":              TenantSetupInfoSetupStatus_NotRegisteredYet,
		"registeredsetupinprogress":     TenantSetupInfoSetupStatus_RegisteredSetupInProgress,
		"registeredsetupnotstarted":     TenantSetupInfoSetupStatus_RegisteredSetupNotStarted,
		"registrationandsetupcompleted": TenantSetupInfoSetupStatus_RegistrationAndSetupCompleted,
		"registrationfailed":            TenantSetupInfoSetupStatus_RegistrationFailed,
		"registrationtimedout":          TenantSetupInfoSetupStatus_RegistrationTimedOut,
		"unknown":                       TenantSetupInfoSetupStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TenantSetupInfoSetupStatus(input)
	return &out, nil
}
