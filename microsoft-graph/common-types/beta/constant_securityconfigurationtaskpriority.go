package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConfigurationTaskPriority string

const (
	SecurityConfigurationTaskPriority_High SecurityConfigurationTaskPriority = "high"
	SecurityConfigurationTaskPriority_Low  SecurityConfigurationTaskPriority = "low"
	SecurityConfigurationTaskPriority_None SecurityConfigurationTaskPriority = "none"
)

func PossibleValuesForSecurityConfigurationTaskPriority() []string {
	return []string{
		string(SecurityConfigurationTaskPriority_High),
		string(SecurityConfigurationTaskPriority_Low),
		string(SecurityConfigurationTaskPriority_None),
	}
}

func (s *SecurityConfigurationTaskPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityConfigurationTaskPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityConfigurationTaskPriority(input string) (*SecurityConfigurationTaskPriority, error) {
	vals := map[string]SecurityConfigurationTaskPriority{
		"high": SecurityConfigurationTaskPriority_High,
		"low":  SecurityConfigurationTaskPriority_Low,
		"none": SecurityConfigurationTaskPriority_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityConfigurationTaskPriority(input)
	return &out, nil
}
