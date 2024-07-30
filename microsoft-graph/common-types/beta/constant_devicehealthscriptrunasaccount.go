package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptRunAsAccount string

const (
	DeviceHealthScriptRunAsAccount_System DeviceHealthScriptRunAsAccount = "system"
	DeviceHealthScriptRunAsAccount_User   DeviceHealthScriptRunAsAccount = "user"
)

func PossibleValuesForDeviceHealthScriptRunAsAccount() []string {
	return []string{
		string(DeviceHealthScriptRunAsAccount_System),
		string(DeviceHealthScriptRunAsAccount_User),
	}
}

func (s *DeviceHealthScriptRunAsAccount) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthScriptRunAsAccount(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthScriptRunAsAccount(input string) (*DeviceHealthScriptRunAsAccount, error) {
	vals := map[string]DeviceHealthScriptRunAsAccount{
		"system": DeviceHealthScriptRunAsAccount_System,
		"user":   DeviceHealthScriptRunAsAccount_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptRunAsAccount(input)
	return &out, nil
}
