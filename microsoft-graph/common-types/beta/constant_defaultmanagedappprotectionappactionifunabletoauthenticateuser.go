package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser string

const (
	DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser = "block"
	DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn  DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser = "warn"
	DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe  DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block),
		string(DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn),
		string(DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe),
	}
}

func (s *DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser(input string) (*DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser{
		"block": DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block,
		"warn":  DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn,
		"wipe":  DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser(input)
	return &out, nil
}
