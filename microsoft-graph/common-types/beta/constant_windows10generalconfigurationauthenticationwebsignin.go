package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationAuthenticationWebSignIn string

const (
	Windows10GeneralConfigurationAuthenticationWebSignIn_Disabled      Windows10GeneralConfigurationAuthenticationWebSignIn = "disabled"
	Windows10GeneralConfigurationAuthenticationWebSignIn_Enabled       Windows10GeneralConfigurationAuthenticationWebSignIn = "enabled"
	Windows10GeneralConfigurationAuthenticationWebSignIn_NotConfigured Windows10GeneralConfigurationAuthenticationWebSignIn = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationAuthenticationWebSignIn() []string {
	return []string{
		string(Windows10GeneralConfigurationAuthenticationWebSignIn_Disabled),
		string(Windows10GeneralConfigurationAuthenticationWebSignIn_Enabled),
		string(Windows10GeneralConfigurationAuthenticationWebSignIn_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationAuthenticationWebSignIn) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationAuthenticationWebSignIn(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationAuthenticationWebSignIn(input string) (*Windows10GeneralConfigurationAuthenticationWebSignIn, error) {
	vals := map[string]Windows10GeneralConfigurationAuthenticationWebSignIn{
		"disabled":      Windows10GeneralConfigurationAuthenticationWebSignIn_Disabled,
		"enabled":       Windows10GeneralConfigurationAuthenticationWebSignIn_Enabled,
		"notconfigured": Windows10GeneralConfigurationAuthenticationWebSignIn_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationAuthenticationWebSignIn(input)
	return &out, nil
}
