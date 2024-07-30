package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAppActionIfUnableToAuthenticateUser string

const (
	IosManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block IosManagedAppProtectionAppActionIfUnableToAuthenticateUser = "block"
	IosManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn  IosManagedAppProtectionAppActionIfUnableToAuthenticateUser = "warn"
	IosManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe  IosManagedAppProtectionAppActionIfUnableToAuthenticateUser = "wipe"
)

func PossibleValuesForIosManagedAppProtectionAppActionIfUnableToAuthenticateUser() []string {
	return []string{
		string(IosManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block),
		string(IosManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn),
		string(IosManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe),
	}
}

func (s *IosManagedAppProtectionAppActionIfUnableToAuthenticateUser) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionAppActionIfUnableToAuthenticateUser(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionAppActionIfUnableToAuthenticateUser(input string) (*IosManagedAppProtectionAppActionIfUnableToAuthenticateUser, error) {
	vals := map[string]IosManagedAppProtectionAppActionIfUnableToAuthenticateUser{
		"block": IosManagedAppProtectionAppActionIfUnableToAuthenticateUser_Block,
		"warn":  IosManagedAppProtectionAppActionIfUnableToAuthenticateUser_Warn,
		"wipe":  IosManagedAppProtectionAppActionIfUnableToAuthenticateUser_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAppActionIfUnableToAuthenticateUser(input)
	return &out, nil
}
