package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationCountRegistrationStatus string

const (
	UserRegistrationCountRegistrationStatus_Capable       UserRegistrationCountRegistrationStatus = "capable"
	UserRegistrationCountRegistrationStatus_Enabled       UserRegistrationCountRegistrationStatus = "enabled"
	UserRegistrationCountRegistrationStatus_MfaRegistered UserRegistrationCountRegistrationStatus = "mfaRegistered"
	UserRegistrationCountRegistrationStatus_Registered    UserRegistrationCountRegistrationStatus = "registered"
)

func PossibleValuesForUserRegistrationCountRegistrationStatus() []string {
	return []string{
		string(UserRegistrationCountRegistrationStatus_Capable),
		string(UserRegistrationCountRegistrationStatus_Enabled),
		string(UserRegistrationCountRegistrationStatus_MfaRegistered),
		string(UserRegistrationCountRegistrationStatus_Registered),
	}
}

func (s *UserRegistrationCountRegistrationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserRegistrationCountRegistrationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserRegistrationCountRegistrationStatus(input string) (*UserRegistrationCountRegistrationStatus, error) {
	vals := map[string]UserRegistrationCountRegistrationStatus{
		"capable":       UserRegistrationCountRegistrationStatus_Capable,
		"enabled":       UserRegistrationCountRegistrationStatus_Enabled,
		"mfaregistered": UserRegistrationCountRegistrationStatus_MfaRegistered,
		"registered":    UserRegistrationCountRegistrationStatus_Registered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationCountRegistrationStatus(input)
	return &out, nil
}
