package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfigurationAllowedAccounts string

const (
	SharedPCConfigurationAllowedAccounts_Domain        SharedPCConfigurationAllowedAccounts = "domain"
	SharedPCConfigurationAllowedAccounts_Guest         SharedPCConfigurationAllowedAccounts = "guest"
	SharedPCConfigurationAllowedAccounts_NotConfigured SharedPCConfigurationAllowedAccounts = "notConfigured"
)

func PossibleValuesForSharedPCConfigurationAllowedAccounts() []string {
	return []string{
		string(SharedPCConfigurationAllowedAccounts_Domain),
		string(SharedPCConfigurationAllowedAccounts_Guest),
		string(SharedPCConfigurationAllowedAccounts_NotConfigured),
	}
}

func (s *SharedPCConfigurationAllowedAccounts) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedPCConfigurationAllowedAccounts(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharedPCConfigurationAllowedAccounts(input string) (*SharedPCConfigurationAllowedAccounts, error) {
	vals := map[string]SharedPCConfigurationAllowedAccounts{
		"domain":        SharedPCConfigurationAllowedAccounts_Domain,
		"guest":         SharedPCConfigurationAllowedAccounts_Guest,
		"notconfigured": SharedPCConfigurationAllowedAccounts_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCConfigurationAllowedAccounts(input)
	return &out, nil
}
