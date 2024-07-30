package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser string

const (
	WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser = "block"
	WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn  WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser = "warn"
	WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe  WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser = "wipe"
)

func PossibleValuesForWindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser() []string {
	return []string{
		string(WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block),
		string(WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn),
		string(WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe),
	}
}

func (s *WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser(input string) (*WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser, error) {
	vals := map[string]WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser{
		"block": WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block,
		"warn":  WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn,
		"wipe":  WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser(input)
	return &out, nil
}
