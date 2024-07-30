package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationSafeSearchFilter string

const (
	Windows10GeneralConfigurationSafeSearchFilter_Moderate    Windows10GeneralConfigurationSafeSearchFilter = "moderate"
	Windows10GeneralConfigurationSafeSearchFilter_Strict      Windows10GeneralConfigurationSafeSearchFilter = "strict"
	Windows10GeneralConfigurationSafeSearchFilter_UserDefined Windows10GeneralConfigurationSafeSearchFilter = "userDefined"
)

func PossibleValuesForWindows10GeneralConfigurationSafeSearchFilter() []string {
	return []string{
		string(Windows10GeneralConfigurationSafeSearchFilter_Moderate),
		string(Windows10GeneralConfigurationSafeSearchFilter_Strict),
		string(Windows10GeneralConfigurationSafeSearchFilter_UserDefined),
	}
}

func (s *Windows10GeneralConfigurationSafeSearchFilter) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationSafeSearchFilter(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationSafeSearchFilter(input string) (*Windows10GeneralConfigurationSafeSearchFilter, error) {
	vals := map[string]Windows10GeneralConfigurationSafeSearchFilter{
		"moderate":    Windows10GeneralConfigurationSafeSearchFilter_Moderate,
		"strict":      Windows10GeneralConfigurationSafeSearchFilter_Strict,
		"userdefined": Windows10GeneralConfigurationSafeSearchFilter_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationSafeSearchFilter(input)
	return &out, nil
}
