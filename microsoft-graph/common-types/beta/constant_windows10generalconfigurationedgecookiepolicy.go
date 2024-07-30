package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationEdgeCookiePolicy string

const (
	Windows10GeneralConfigurationEdgeCookiePolicy_Allow           Windows10GeneralConfigurationEdgeCookiePolicy = "allow"
	Windows10GeneralConfigurationEdgeCookiePolicy_BlockAll        Windows10GeneralConfigurationEdgeCookiePolicy = "blockAll"
	Windows10GeneralConfigurationEdgeCookiePolicy_BlockThirdParty Windows10GeneralConfigurationEdgeCookiePolicy = "blockThirdParty"
	Windows10GeneralConfigurationEdgeCookiePolicy_UserDefined     Windows10GeneralConfigurationEdgeCookiePolicy = "userDefined"
)

func PossibleValuesForWindows10GeneralConfigurationEdgeCookiePolicy() []string {
	return []string{
		string(Windows10GeneralConfigurationEdgeCookiePolicy_Allow),
		string(Windows10GeneralConfigurationEdgeCookiePolicy_BlockAll),
		string(Windows10GeneralConfigurationEdgeCookiePolicy_BlockThirdParty),
		string(Windows10GeneralConfigurationEdgeCookiePolicy_UserDefined),
	}
}

func (s *Windows10GeneralConfigurationEdgeCookiePolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationEdgeCookiePolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationEdgeCookiePolicy(input string) (*Windows10GeneralConfigurationEdgeCookiePolicy, error) {
	vals := map[string]Windows10GeneralConfigurationEdgeCookiePolicy{
		"allow":           Windows10GeneralConfigurationEdgeCookiePolicy_Allow,
		"blockall":        Windows10GeneralConfigurationEdgeCookiePolicy_BlockAll,
		"blockthirdparty": Windows10GeneralConfigurationEdgeCookiePolicy_BlockThirdParty,
		"userdefined":     Windows10GeneralConfigurationEdgeCookiePolicy_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationEdgeCookiePolicy(input)
	return &out, nil
}
