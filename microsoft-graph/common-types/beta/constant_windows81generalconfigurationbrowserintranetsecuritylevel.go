package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationBrowserIntranetSecurityLevel string

const (
	Windows81GeneralConfigurationBrowserIntranetSecurityLevel_High        Windows81GeneralConfigurationBrowserIntranetSecurityLevel = "high"
	Windows81GeneralConfigurationBrowserIntranetSecurityLevel_Low         Windows81GeneralConfigurationBrowserIntranetSecurityLevel = "low"
	Windows81GeneralConfigurationBrowserIntranetSecurityLevel_Medium      Windows81GeneralConfigurationBrowserIntranetSecurityLevel = "medium"
	Windows81GeneralConfigurationBrowserIntranetSecurityLevel_MediumHigh  Windows81GeneralConfigurationBrowserIntranetSecurityLevel = "mediumHigh"
	Windows81GeneralConfigurationBrowserIntranetSecurityLevel_MediumLow   Windows81GeneralConfigurationBrowserIntranetSecurityLevel = "mediumLow"
	Windows81GeneralConfigurationBrowserIntranetSecurityLevel_UserDefined Windows81GeneralConfigurationBrowserIntranetSecurityLevel = "userDefined"
)

func PossibleValuesForWindows81GeneralConfigurationBrowserIntranetSecurityLevel() []string {
	return []string{
		string(Windows81GeneralConfigurationBrowserIntranetSecurityLevel_High),
		string(Windows81GeneralConfigurationBrowserIntranetSecurityLevel_Low),
		string(Windows81GeneralConfigurationBrowserIntranetSecurityLevel_Medium),
		string(Windows81GeneralConfigurationBrowserIntranetSecurityLevel_MediumHigh),
		string(Windows81GeneralConfigurationBrowserIntranetSecurityLevel_MediumLow),
		string(Windows81GeneralConfigurationBrowserIntranetSecurityLevel_UserDefined),
	}
}

func (s *Windows81GeneralConfigurationBrowserIntranetSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81GeneralConfigurationBrowserIntranetSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81GeneralConfigurationBrowserIntranetSecurityLevel(input string) (*Windows81GeneralConfigurationBrowserIntranetSecurityLevel, error) {
	vals := map[string]Windows81GeneralConfigurationBrowserIntranetSecurityLevel{
		"high":        Windows81GeneralConfigurationBrowserIntranetSecurityLevel_High,
		"low":         Windows81GeneralConfigurationBrowserIntranetSecurityLevel_Low,
		"medium":      Windows81GeneralConfigurationBrowserIntranetSecurityLevel_Medium,
		"mediumhigh":  Windows81GeneralConfigurationBrowserIntranetSecurityLevel_MediumHigh,
		"mediumlow":   Windows81GeneralConfigurationBrowserIntranetSecurityLevel_MediumLow,
		"userdefined": Windows81GeneralConfigurationBrowserIntranetSecurityLevel_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationBrowserIntranetSecurityLevel(input)
	return &out, nil
}
