package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidGeneralDeviceConfigurationRequiredPasswordComplexity string

const (
	AndroidGeneralDeviceConfigurationRequiredPasswordComplexity_High   AndroidGeneralDeviceConfigurationRequiredPasswordComplexity = "high"
	AndroidGeneralDeviceConfigurationRequiredPasswordComplexity_Low    AndroidGeneralDeviceConfigurationRequiredPasswordComplexity = "low"
	AndroidGeneralDeviceConfigurationRequiredPasswordComplexity_Medium AndroidGeneralDeviceConfigurationRequiredPasswordComplexity = "medium"
	AndroidGeneralDeviceConfigurationRequiredPasswordComplexity_None   AndroidGeneralDeviceConfigurationRequiredPasswordComplexity = "none"
)

func PossibleValuesForAndroidGeneralDeviceConfigurationRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidGeneralDeviceConfigurationRequiredPasswordComplexity_High),
		string(AndroidGeneralDeviceConfigurationRequiredPasswordComplexity_Low),
		string(AndroidGeneralDeviceConfigurationRequiredPasswordComplexity_Medium),
		string(AndroidGeneralDeviceConfigurationRequiredPasswordComplexity_None),
	}
}

func (s *AndroidGeneralDeviceConfigurationRequiredPasswordComplexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidGeneralDeviceConfigurationRequiredPasswordComplexity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidGeneralDeviceConfigurationRequiredPasswordComplexity(input string) (*AndroidGeneralDeviceConfigurationRequiredPasswordComplexity, error) {
	vals := map[string]AndroidGeneralDeviceConfigurationRequiredPasswordComplexity{
		"high":   AndroidGeneralDeviceConfigurationRequiredPasswordComplexity_High,
		"low":    AndroidGeneralDeviceConfigurationRequiredPasswordComplexity_Low,
		"medium": AndroidGeneralDeviceConfigurationRequiredPasswordComplexity_Medium,
		"none":   AndroidGeneralDeviceConfigurationRequiredPasswordComplexity_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidGeneralDeviceConfigurationRequiredPasswordComplexity(input)
	return &out, nil
}
