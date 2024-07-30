package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfigurationSetPowerPolicies string

const (
	SharedPCConfigurationSetPowerPolicies_Disabled      SharedPCConfigurationSetPowerPolicies = "disabled"
	SharedPCConfigurationSetPowerPolicies_Enabled       SharedPCConfigurationSetPowerPolicies = "enabled"
	SharedPCConfigurationSetPowerPolicies_NotConfigured SharedPCConfigurationSetPowerPolicies = "notConfigured"
)

func PossibleValuesForSharedPCConfigurationSetPowerPolicies() []string {
	return []string{
		string(SharedPCConfigurationSetPowerPolicies_Disabled),
		string(SharedPCConfigurationSetPowerPolicies_Enabled),
		string(SharedPCConfigurationSetPowerPolicies_NotConfigured),
	}
}

func (s *SharedPCConfigurationSetPowerPolicies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedPCConfigurationSetPowerPolicies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharedPCConfigurationSetPowerPolicies(input string) (*SharedPCConfigurationSetPowerPolicies, error) {
	vals := map[string]SharedPCConfigurationSetPowerPolicies{
		"disabled":      SharedPCConfigurationSetPowerPolicies_Disabled,
		"enabled":       SharedPCConfigurationSetPowerPolicies_Enabled,
		"notconfigured": SharedPCConfigurationSetPowerPolicies_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCConfigurationSetPowerPolicies(input)
	return &out, nil
}
