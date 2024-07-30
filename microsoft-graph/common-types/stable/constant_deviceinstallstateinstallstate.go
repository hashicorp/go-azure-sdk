package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceInstallStateInstallState string

const (
	DeviceInstallStateInstallState_Failed          DeviceInstallStateInstallState = "failed"
	DeviceInstallStateInstallState_Installed       DeviceInstallStateInstallState = "installed"
	DeviceInstallStateInstallState_NotApplicable   DeviceInstallStateInstallState = "notApplicable"
	DeviceInstallStateInstallState_NotInstalled    DeviceInstallStateInstallState = "notInstalled"
	DeviceInstallStateInstallState_UninstallFailed DeviceInstallStateInstallState = "uninstallFailed"
	DeviceInstallStateInstallState_Unknown         DeviceInstallStateInstallState = "unknown"
)

func PossibleValuesForDeviceInstallStateInstallState() []string {
	return []string{
		string(DeviceInstallStateInstallState_Failed),
		string(DeviceInstallStateInstallState_Installed),
		string(DeviceInstallStateInstallState_NotApplicable),
		string(DeviceInstallStateInstallState_NotInstalled),
		string(DeviceInstallStateInstallState_UninstallFailed),
		string(DeviceInstallStateInstallState_Unknown),
	}
}

func (s *DeviceInstallStateInstallState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceInstallStateInstallState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceInstallStateInstallState(input string) (*DeviceInstallStateInstallState, error) {
	vals := map[string]DeviceInstallStateInstallState{
		"failed":          DeviceInstallStateInstallState_Failed,
		"installed":       DeviceInstallStateInstallState_Installed,
		"notapplicable":   DeviceInstallStateInstallState_NotApplicable,
		"notinstalled":    DeviceInstallStateInstallState_NotInstalled,
		"uninstallfailed": DeviceInstallStateInstallState_UninstallFailed,
		"unknown":         DeviceInstallStateInstallState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceInstallStateInstallState(input)
	return &out, nil
}
