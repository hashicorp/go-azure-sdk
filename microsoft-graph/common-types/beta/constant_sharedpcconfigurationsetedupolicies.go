package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfigurationSetEduPolicies string

const (
	SharedPCConfigurationSetEduPolicies_Disabled      SharedPCConfigurationSetEduPolicies = "disabled"
	SharedPCConfigurationSetEduPolicies_Enabled       SharedPCConfigurationSetEduPolicies = "enabled"
	SharedPCConfigurationSetEduPolicies_NotConfigured SharedPCConfigurationSetEduPolicies = "notConfigured"
)

func PossibleValuesForSharedPCConfigurationSetEduPolicies() []string {
	return []string{
		string(SharedPCConfigurationSetEduPolicies_Disabled),
		string(SharedPCConfigurationSetEduPolicies_Enabled),
		string(SharedPCConfigurationSetEduPolicies_NotConfigured),
	}
}

func (s *SharedPCConfigurationSetEduPolicies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedPCConfigurationSetEduPolicies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharedPCConfigurationSetEduPolicies(input string) (*SharedPCConfigurationSetEduPolicies, error) {
	vals := map[string]SharedPCConfigurationSetEduPolicies{
		"disabled":      SharedPCConfigurationSetEduPolicies_Disabled,
		"enabled":       SharedPCConfigurationSetEduPolicies_Enabled,
		"notconfigured": SharedPCConfigurationSetEduPolicies_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCConfigurationSetEduPolicies(input)
	return &out, nil
}
