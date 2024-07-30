package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationUserAccountControlSettings string

const (
	Windows81GeneralConfigurationUserAccountControlSettings_AlwaysNotify                     Windows81GeneralConfigurationUserAccountControlSettings = "alwaysNotify"
	Windows81GeneralConfigurationUserAccountControlSettings_NeverNotify                      Windows81GeneralConfigurationUserAccountControlSettings = "neverNotify"
	Windows81GeneralConfigurationUserAccountControlSettings_NotifyOnAppChanges               Windows81GeneralConfigurationUserAccountControlSettings = "notifyOnAppChanges"
	Windows81GeneralConfigurationUserAccountControlSettings_NotifyOnAppChangesWithoutDimming Windows81GeneralConfigurationUserAccountControlSettings = "notifyOnAppChangesWithoutDimming"
	Windows81GeneralConfigurationUserAccountControlSettings_UserDefined                      Windows81GeneralConfigurationUserAccountControlSettings = "userDefined"
)

func PossibleValuesForWindows81GeneralConfigurationUserAccountControlSettings() []string {
	return []string{
		string(Windows81GeneralConfigurationUserAccountControlSettings_AlwaysNotify),
		string(Windows81GeneralConfigurationUserAccountControlSettings_NeverNotify),
		string(Windows81GeneralConfigurationUserAccountControlSettings_NotifyOnAppChanges),
		string(Windows81GeneralConfigurationUserAccountControlSettings_NotifyOnAppChangesWithoutDimming),
		string(Windows81GeneralConfigurationUserAccountControlSettings_UserDefined),
	}
}

func (s *Windows81GeneralConfigurationUserAccountControlSettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81GeneralConfigurationUserAccountControlSettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81GeneralConfigurationUserAccountControlSettings(input string) (*Windows81GeneralConfigurationUserAccountControlSettings, error) {
	vals := map[string]Windows81GeneralConfigurationUserAccountControlSettings{
		"alwaysnotify":                     Windows81GeneralConfigurationUserAccountControlSettings_AlwaysNotify,
		"nevernotify":                      Windows81GeneralConfigurationUserAccountControlSettings_NeverNotify,
		"notifyonappchanges":               Windows81GeneralConfigurationUserAccountControlSettings_NotifyOnAppChanges,
		"notifyonappchangeswithoutdimming": Windows81GeneralConfigurationUserAccountControlSettings_NotifyOnAppChangesWithoutDimming,
		"userdefined":                      Windows81GeneralConfigurationUserAccountControlSettings_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationUserAccountControlSettings(input)
	return &out, nil
}
