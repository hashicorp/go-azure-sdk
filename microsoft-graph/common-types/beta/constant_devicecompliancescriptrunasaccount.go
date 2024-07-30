package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRunAsAccount string

const (
	DeviceComplianceScriptRunAsAccount_System DeviceComplianceScriptRunAsAccount = "system"
	DeviceComplianceScriptRunAsAccount_User   DeviceComplianceScriptRunAsAccount = "user"
)

func PossibleValuesForDeviceComplianceScriptRunAsAccount() []string {
	return []string{
		string(DeviceComplianceScriptRunAsAccount_System),
		string(DeviceComplianceScriptRunAsAccount_User),
	}
}

func (s *DeviceComplianceScriptRunAsAccount) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceScriptRunAsAccount(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceScriptRunAsAccount(input string) (*DeviceComplianceScriptRunAsAccount, error) {
	vals := map[string]DeviceComplianceScriptRunAsAccount{
		"system": DeviceComplianceScriptRunAsAccount_System,
		"user":   DeviceComplianceScriptRunAsAccount_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRunAsAccount(input)
	return &out, nil
}
