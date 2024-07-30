package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly string

const (
	WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_All                        WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly = "all"
	WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_BusinessReadyOnly          WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly = "businessReadyOnly"
	WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_UserDefined                WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly = "userDefined"
	WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_WindowsInsiderBuildFast    WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly = "windowsInsiderBuildFast"
	WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_WindowsInsiderBuildRelease WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly = "windowsInsiderBuildRelease"
	WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_WindowsInsiderBuildSlow    WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly = "windowsInsiderBuildSlow"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_All),
		string(WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_BusinessReadyOnly),
		string(WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_UserDefined),
		string(WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_WindowsInsiderBuildFast),
		string(WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_WindowsInsiderBuildRelease),
		string(WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_WindowsInsiderBuildSlow),
	}
}

func (s *WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly(input string) (*WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly{
		"all":                        WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_All,
		"businessreadyonly":          WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_BusinessReadyOnly,
		"userdefined":                WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_UserDefined,
		"windowsinsiderbuildfast":    WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_WindowsInsiderBuildFast,
		"windowsinsiderbuildrelease": WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_WindowsInsiderBuildRelease,
		"windowsinsiderbuildslow":    WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly_WindowsInsiderBuildSlow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly(input)
	return &out, nil
}
