package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfigurationSetAccountManager string

const (
	SharedPCConfigurationSetAccountManager_Disabled      SharedPCConfigurationSetAccountManager = "disabled"
	SharedPCConfigurationSetAccountManager_Enabled       SharedPCConfigurationSetAccountManager = "enabled"
	SharedPCConfigurationSetAccountManager_NotConfigured SharedPCConfigurationSetAccountManager = "notConfigured"
)

func PossibleValuesForSharedPCConfigurationSetAccountManager() []string {
	return []string{
		string(SharedPCConfigurationSetAccountManager_Disabled),
		string(SharedPCConfigurationSetAccountManager_Enabled),
		string(SharedPCConfigurationSetAccountManager_NotConfigured),
	}
}

func (s *SharedPCConfigurationSetAccountManager) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedPCConfigurationSetAccountManager(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharedPCConfigurationSetAccountManager(input string) (*SharedPCConfigurationSetAccountManager, error) {
	vals := map[string]SharedPCConfigurationSetAccountManager{
		"disabled":      SharedPCConfigurationSetAccountManager_Disabled,
		"enabled":       SharedPCConfigurationSetAccountManager_Enabled,
		"notconfigured": SharedPCConfigurationSetAccountManager_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCConfigurationSetAccountManager(input)
	return &out, nil
}
