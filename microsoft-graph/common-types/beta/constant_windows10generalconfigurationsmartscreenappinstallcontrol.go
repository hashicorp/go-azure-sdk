package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationSmartScreenAppInstallControl string

const (
	Windows10GeneralConfigurationSmartScreenAppInstallControl_Anywhere        Windows10GeneralConfigurationSmartScreenAppInstallControl = "anywhere"
	Windows10GeneralConfigurationSmartScreenAppInstallControl_NotConfigured   Windows10GeneralConfigurationSmartScreenAppInstallControl = "notConfigured"
	Windows10GeneralConfigurationSmartScreenAppInstallControl_PreferStore     Windows10GeneralConfigurationSmartScreenAppInstallControl = "preferStore"
	Windows10GeneralConfigurationSmartScreenAppInstallControl_Recommendations Windows10GeneralConfigurationSmartScreenAppInstallControl = "recommendations"
	Windows10GeneralConfigurationSmartScreenAppInstallControl_StoreOnly       Windows10GeneralConfigurationSmartScreenAppInstallControl = "storeOnly"
)

func PossibleValuesForWindows10GeneralConfigurationSmartScreenAppInstallControl() []string {
	return []string{
		string(Windows10GeneralConfigurationSmartScreenAppInstallControl_Anywhere),
		string(Windows10GeneralConfigurationSmartScreenAppInstallControl_NotConfigured),
		string(Windows10GeneralConfigurationSmartScreenAppInstallControl_PreferStore),
		string(Windows10GeneralConfigurationSmartScreenAppInstallControl_Recommendations),
		string(Windows10GeneralConfigurationSmartScreenAppInstallControl_StoreOnly),
	}
}

func (s *Windows10GeneralConfigurationSmartScreenAppInstallControl) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationSmartScreenAppInstallControl(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationSmartScreenAppInstallControl(input string) (*Windows10GeneralConfigurationSmartScreenAppInstallControl, error) {
	vals := map[string]Windows10GeneralConfigurationSmartScreenAppInstallControl{
		"anywhere":        Windows10GeneralConfigurationSmartScreenAppInstallControl_Anywhere,
		"notconfigured":   Windows10GeneralConfigurationSmartScreenAppInstallControl_NotConfigured,
		"preferstore":     Windows10GeneralConfigurationSmartScreenAppInstallControl_PreferStore,
		"recommendations": Windows10GeneralConfigurationSmartScreenAppInstallControl_Recommendations,
		"storeonly":       Windows10GeneralConfigurationSmartScreenAppInstallControl_StoreOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationSmartScreenAppInstallControl(input)
	return &out, nil
}
