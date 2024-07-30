package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceShellScriptRunAsAccount string

const (
	DeviceShellScriptRunAsAccount_System DeviceShellScriptRunAsAccount = "system"
	DeviceShellScriptRunAsAccount_User   DeviceShellScriptRunAsAccount = "user"
)

func PossibleValuesForDeviceShellScriptRunAsAccount() []string {
	return []string{
		string(DeviceShellScriptRunAsAccount_System),
		string(DeviceShellScriptRunAsAccount_User),
	}
}

func (s *DeviceShellScriptRunAsAccount) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceShellScriptRunAsAccount(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceShellScriptRunAsAccount(input string) (*DeviceShellScriptRunAsAccount, error) {
	vals := map[string]DeviceShellScriptRunAsAccount{
		"system": DeviceShellScriptRunAsAccount_System,
		"user":   DeviceShellScriptRunAsAccount_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceShellScriptRunAsAccount(input)
	return &out, nil
}
