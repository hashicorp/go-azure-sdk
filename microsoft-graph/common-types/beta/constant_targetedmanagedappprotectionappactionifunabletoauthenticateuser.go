package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser string

const (
	TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser = "block"
	TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn  TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser = "warn"
	TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe  TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser = "wipe"
)

func PossibleValuesForTargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser() []string {
	return []string{
		string(TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block),
		string(TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn),
		string(TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe),
	}
}

func (s *TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser(input string) (*TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser, error) {
	vals := map[string]TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser{
		"block": TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block,
		"warn":  TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn,
		"wipe":  TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser(input)
	return &out, nil
}
