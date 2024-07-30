package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAppActionIfUnableToAuthenticateUser string

const (
	ManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block ManagedAppProtectionAppActionIfUnableToAuthenticateUser = "block"
	ManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn  ManagedAppProtectionAppActionIfUnableToAuthenticateUser = "warn"
	ManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe  ManagedAppProtectionAppActionIfUnableToAuthenticateUser = "wipe"
)

func PossibleValuesForManagedAppProtectionAppActionIfUnableToAuthenticateUser() []string {
	return []string{
		string(ManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block),
		string(ManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn),
		string(ManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe),
	}
}

func (s *ManagedAppProtectionAppActionIfUnableToAuthenticateUser) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionAppActionIfUnableToAuthenticateUser(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionAppActionIfUnableToAuthenticateUser(input string) (*ManagedAppProtectionAppActionIfUnableToAuthenticateUser, error) {
	vals := map[string]ManagedAppProtectionAppActionIfUnableToAuthenticateUser{
		"block": ManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block,
		"warn":  ManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn,
		"wipe":  ManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAppActionIfUnableToAuthenticateUser(input)
	return &out, nil
}
