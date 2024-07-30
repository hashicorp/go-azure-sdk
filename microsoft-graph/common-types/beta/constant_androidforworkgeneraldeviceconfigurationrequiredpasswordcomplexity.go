package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity string

const (
	AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity_High   AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity = "high"
	AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity_Low    AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity = "low"
	AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity_Medium AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity = "medium"
	AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity_None   AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity = "none"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity_High),
		string(AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity_Low),
		string(AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity_Medium),
		string(AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity_None),
	}
}

func (s *AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity(input string) (*AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity{
		"high":   AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity_High,
		"low":    AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity_Low,
		"medium": AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity_Medium,
		"none":   AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity(input)
	return &out, nil
}
