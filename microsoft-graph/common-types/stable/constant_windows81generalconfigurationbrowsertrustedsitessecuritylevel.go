package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel string

const (
	Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_High        Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel = "high"
	Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_Low         Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel = "low"
	Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_Medium      Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel = "medium"
	Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_MediumHigh  Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel = "mediumHigh"
	Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_MediumLow   Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel = "mediumLow"
	Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_UserDefined Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel = "userDefined"
)

func PossibleValuesForWindows81GeneralConfigurationBrowserTrustedSitesSecurityLevel() []string {
	return []string{
		string(Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_High),
		string(Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_Low),
		string(Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_Medium),
		string(Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_MediumHigh),
		string(Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_MediumLow),
		string(Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_UserDefined),
	}
}

func (s *Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81GeneralConfigurationBrowserTrustedSitesSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81GeneralConfigurationBrowserTrustedSitesSecurityLevel(input string) (*Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel, error) {
	vals := map[string]Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel{
		"high":        Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_High,
		"low":         Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_Low,
		"medium":      Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_Medium,
		"mediumhigh":  Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_MediumHigh,
		"mediumlow":   Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_MediumLow,
		"userdefined": Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationBrowserTrustedSitesSecurityLevel(input)
	return &out, nil
}
