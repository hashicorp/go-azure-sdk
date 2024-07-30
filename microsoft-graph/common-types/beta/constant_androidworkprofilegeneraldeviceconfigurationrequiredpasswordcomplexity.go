package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity string

const (
	AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity_High   AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity = "high"
	AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity_Low    AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity = "low"
	AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity_Medium AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity = "medium"
	AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity_None   AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity = "none"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity_High),
		string(AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity_Low),
		string(AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity_Medium),
		string(AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity_None),
	}
}

func (s *AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity(input string) (*AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity{
		"high":   AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity_High,
		"low":    AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity_Low,
		"medium": AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity_Medium,
		"none":   AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity(input)
	return &out, nil
}
