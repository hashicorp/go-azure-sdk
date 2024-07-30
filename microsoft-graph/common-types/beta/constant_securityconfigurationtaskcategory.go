package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConfigurationTaskCategory string

const (
	SecurityConfigurationTaskCategory_AdvancedThreatProtection SecurityConfigurationTaskCategory = "advancedThreatProtection"
	SecurityConfigurationTaskCategory_Unknown                  SecurityConfigurationTaskCategory = "unknown"
)

func PossibleValuesForSecurityConfigurationTaskCategory() []string {
	return []string{
		string(SecurityConfigurationTaskCategory_AdvancedThreatProtection),
		string(SecurityConfigurationTaskCategory_Unknown),
	}
}

func (s *SecurityConfigurationTaskCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityConfigurationTaskCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityConfigurationTaskCategory(input string) (*SecurityConfigurationTaskCategory, error) {
	vals := map[string]SecurityConfigurationTaskCategory{
		"advancedthreatprotection": SecurityConfigurationTaskCategory_AdvancedThreatProtection,
		"unknown":                  SecurityConfigurationTaskCategory_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityConfigurationTaskCategory(input)
	return &out, nil
}
