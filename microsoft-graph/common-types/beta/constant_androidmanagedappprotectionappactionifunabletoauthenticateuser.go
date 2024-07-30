package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser string

const (
	AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser = "block"
	AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn  AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser = "warn"
	AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe  AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block),
		string(AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn),
		string(AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser(input string) (*AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser{
		"block": AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser(input)
	return &out, nil
}
