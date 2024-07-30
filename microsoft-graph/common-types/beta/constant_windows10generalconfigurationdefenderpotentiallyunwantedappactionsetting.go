package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting string

const (
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_AuditMode     Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting = "auditMode"
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_Enable        Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting = "enable"
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_NotConfigured Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting = "notConfigured"
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_UserDefined   Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting = "userDefined"
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_Warn          Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting = "warn"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_AuditMode),
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_Enable),
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_NotConfigured),
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_UserDefined),
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_Warn),
	}
}

func (s *Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting(input string) (*Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting{
		"auditmode":     Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_AuditMode,
		"enable":        Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_Enable,
		"notconfigured": Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_NotConfigured,
		"userdefined":   Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_UserDefined,
		"warn":          Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting(input)
	return &out, nil
}
