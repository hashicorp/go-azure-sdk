package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPrivacyDataAccessControlItemAccessLevel string

const (
	WindowsPrivacyDataAccessControlItemAccessLevel_ForceAllow    WindowsPrivacyDataAccessControlItemAccessLevel = "forceAllow"
	WindowsPrivacyDataAccessControlItemAccessLevel_ForceDeny     WindowsPrivacyDataAccessControlItemAccessLevel = "forceDeny"
	WindowsPrivacyDataAccessControlItemAccessLevel_NotConfigured WindowsPrivacyDataAccessControlItemAccessLevel = "notConfigured"
	WindowsPrivacyDataAccessControlItemAccessLevel_UserInControl WindowsPrivacyDataAccessControlItemAccessLevel = "userInControl"
)

func PossibleValuesForWindowsPrivacyDataAccessControlItemAccessLevel() []string {
	return []string{
		string(WindowsPrivacyDataAccessControlItemAccessLevel_ForceAllow),
		string(WindowsPrivacyDataAccessControlItemAccessLevel_ForceDeny),
		string(WindowsPrivacyDataAccessControlItemAccessLevel_NotConfigured),
		string(WindowsPrivacyDataAccessControlItemAccessLevel_UserInControl),
	}
}

func (s *WindowsPrivacyDataAccessControlItemAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPrivacyDataAccessControlItemAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPrivacyDataAccessControlItemAccessLevel(input string) (*WindowsPrivacyDataAccessControlItemAccessLevel, error) {
	vals := map[string]WindowsPrivacyDataAccessControlItemAccessLevel{
		"forceallow":    WindowsPrivacyDataAccessControlItemAccessLevel_ForceAllow,
		"forcedeny":     WindowsPrivacyDataAccessControlItemAccessLevel_ForceDeny,
		"notconfigured": WindowsPrivacyDataAccessControlItemAccessLevel_NotConfigured,
		"userincontrol": WindowsPrivacyDataAccessControlItemAccessLevel_UserInControl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPrivacyDataAccessControlItemAccessLevel(input)
	return &out, nil
}
