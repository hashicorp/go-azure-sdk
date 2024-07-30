package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfigurationSignInOnResume string

const (
	SharedPCConfigurationSignInOnResume_Disabled      SharedPCConfigurationSignInOnResume = "disabled"
	SharedPCConfigurationSignInOnResume_Enabled       SharedPCConfigurationSignInOnResume = "enabled"
	SharedPCConfigurationSignInOnResume_NotConfigured SharedPCConfigurationSignInOnResume = "notConfigured"
)

func PossibleValuesForSharedPCConfigurationSignInOnResume() []string {
	return []string{
		string(SharedPCConfigurationSignInOnResume_Disabled),
		string(SharedPCConfigurationSignInOnResume_Enabled),
		string(SharedPCConfigurationSignInOnResume_NotConfigured),
	}
}

func (s *SharedPCConfigurationSignInOnResume) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedPCConfigurationSignInOnResume(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharedPCConfigurationSignInOnResume(input string) (*SharedPCConfigurationSignInOnResume, error) {
	vals := map[string]SharedPCConfigurationSignInOnResume{
		"disabled":      SharedPCConfigurationSignInOnResume_Disabled,
		"enabled":       SharedPCConfigurationSignInOnResume_Enabled,
		"notconfigured": SharedPCConfigurationSignInOnResume_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCConfigurationSignInOnResume(input)
	return &out, nil
}
