package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptRunAsAccount string

const (
	DeviceManagementScriptRunAsAccount_System DeviceManagementScriptRunAsAccount = "system"
	DeviceManagementScriptRunAsAccount_User   DeviceManagementScriptRunAsAccount = "user"
)

func PossibleValuesForDeviceManagementScriptRunAsAccount() []string {
	return []string{
		string(DeviceManagementScriptRunAsAccount_System),
		string(DeviceManagementScriptRunAsAccount_User),
	}
}

func (s *DeviceManagementScriptRunAsAccount) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementScriptRunAsAccount(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementScriptRunAsAccount(input string) (*DeviceManagementScriptRunAsAccount, error) {
	vals := map[string]DeviceManagementScriptRunAsAccount{
		"system": DeviceManagementScriptRunAsAccount_System,
		"user":   DeviceManagementScriptRunAsAccount_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementScriptRunAsAccount(input)
	return &out, nil
}
