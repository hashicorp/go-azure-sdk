package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity string

const (
	AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity_High   AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity = "high"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity_Low    AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity = "low"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity_Medium AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity = "medium"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity_None   AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity = "none"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity_High),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity_Low),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity_Medium),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity_None),
	}
}

func (s *AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity(input string) (*AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity{
		"high":   AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity_High,
		"low":    AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity_Low,
		"medium": AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity_Medium,
		"none":   AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity(input)
	return &out, nil
}
