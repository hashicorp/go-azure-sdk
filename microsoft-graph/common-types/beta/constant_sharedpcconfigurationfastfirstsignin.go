package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfigurationFastFirstSignIn string

const (
	SharedPCConfigurationFastFirstSignIn_Disabled      SharedPCConfigurationFastFirstSignIn = "disabled"
	SharedPCConfigurationFastFirstSignIn_Enabled       SharedPCConfigurationFastFirstSignIn = "enabled"
	SharedPCConfigurationFastFirstSignIn_NotConfigured SharedPCConfigurationFastFirstSignIn = "notConfigured"
)

func PossibleValuesForSharedPCConfigurationFastFirstSignIn() []string {
	return []string{
		string(SharedPCConfigurationFastFirstSignIn_Disabled),
		string(SharedPCConfigurationFastFirstSignIn_Enabled),
		string(SharedPCConfigurationFastFirstSignIn_NotConfigured),
	}
}

func (s *SharedPCConfigurationFastFirstSignIn) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedPCConfigurationFastFirstSignIn(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharedPCConfigurationFastFirstSignIn(input string) (*SharedPCConfigurationFastFirstSignIn, error) {
	vals := map[string]SharedPCConfigurationFastFirstSignIn{
		"disabled":      SharedPCConfigurationFastFirstSignIn_Disabled,
		"enabled":       SharedPCConfigurationFastFirstSignIn_Enabled,
		"notconfigured": SharedPCConfigurationFastFirstSignIn_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCConfigurationFastFirstSignIn(input)
	return &out, nil
}
