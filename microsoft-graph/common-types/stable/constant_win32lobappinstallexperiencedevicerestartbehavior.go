package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppInstallExperienceDeviceRestartBehavior string

const (
	Win32LobAppInstallExperienceDeviceRestartBehavior_Allow             Win32LobAppInstallExperienceDeviceRestartBehavior = "allow"
	Win32LobAppInstallExperienceDeviceRestartBehavior_BasedOnReturnCode Win32LobAppInstallExperienceDeviceRestartBehavior = "basedOnReturnCode"
	Win32LobAppInstallExperienceDeviceRestartBehavior_Force             Win32LobAppInstallExperienceDeviceRestartBehavior = "force"
	Win32LobAppInstallExperienceDeviceRestartBehavior_Suppress          Win32LobAppInstallExperienceDeviceRestartBehavior = "suppress"
)

func PossibleValuesForWin32LobAppInstallExperienceDeviceRestartBehavior() []string {
	return []string{
		string(Win32LobAppInstallExperienceDeviceRestartBehavior_Allow),
		string(Win32LobAppInstallExperienceDeviceRestartBehavior_BasedOnReturnCode),
		string(Win32LobAppInstallExperienceDeviceRestartBehavior_Force),
		string(Win32LobAppInstallExperienceDeviceRestartBehavior_Suppress),
	}
}

func (s *Win32LobAppInstallExperienceDeviceRestartBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppInstallExperienceDeviceRestartBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppInstallExperienceDeviceRestartBehavior(input string) (*Win32LobAppInstallExperienceDeviceRestartBehavior, error) {
	vals := map[string]Win32LobAppInstallExperienceDeviceRestartBehavior{
		"allow":             Win32LobAppInstallExperienceDeviceRestartBehavior_Allow,
		"basedonreturncode": Win32LobAppInstallExperienceDeviceRestartBehavior_BasedOnReturnCode,
		"force":             Win32LobAppInstallExperienceDeviceRestartBehavior_Force,
		"suppress":          Win32LobAppInstallExperienceDeviceRestartBehavior_Suppress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppInstallExperienceDeviceRestartBehavior(input)
	return &out, nil
}
