package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationBrowserInternetSecurityLevel string

const (
	Windows81GeneralConfigurationBrowserInternetSecurityLevel_High        Windows81GeneralConfigurationBrowserInternetSecurityLevel = "high"
	Windows81GeneralConfigurationBrowserInternetSecurityLevel_Medium      Windows81GeneralConfigurationBrowserInternetSecurityLevel = "medium"
	Windows81GeneralConfigurationBrowserInternetSecurityLevel_MediumHigh  Windows81GeneralConfigurationBrowserInternetSecurityLevel = "mediumHigh"
	Windows81GeneralConfigurationBrowserInternetSecurityLevel_UserDefined Windows81GeneralConfigurationBrowserInternetSecurityLevel = "userDefined"
)

func PossibleValuesForWindows81GeneralConfigurationBrowserInternetSecurityLevel() []string {
	return []string{
		string(Windows81GeneralConfigurationBrowserInternetSecurityLevel_High),
		string(Windows81GeneralConfigurationBrowserInternetSecurityLevel_Medium),
		string(Windows81GeneralConfigurationBrowserInternetSecurityLevel_MediumHigh),
		string(Windows81GeneralConfigurationBrowserInternetSecurityLevel_UserDefined),
	}
}

func (s *Windows81GeneralConfigurationBrowserInternetSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81GeneralConfigurationBrowserInternetSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81GeneralConfigurationBrowserInternetSecurityLevel(input string) (*Windows81GeneralConfigurationBrowserInternetSecurityLevel, error) {
	vals := map[string]Windows81GeneralConfigurationBrowserInternetSecurityLevel{
		"high":        Windows81GeneralConfigurationBrowserInternetSecurityLevel_High,
		"medium":      Windows81GeneralConfigurationBrowserInternetSecurityLevel_Medium,
		"mediumhigh":  Windows81GeneralConfigurationBrowserInternetSecurityLevel_MediumHigh,
		"userdefined": Windows81GeneralConfigurationBrowserInternetSecurityLevel_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationBrowserInternetSecurityLevel(input)
	return &out, nil
}
