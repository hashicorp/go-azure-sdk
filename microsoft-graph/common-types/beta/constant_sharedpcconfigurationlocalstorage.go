package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfigurationLocalStorage string

const (
	SharedPCConfigurationLocalStorage_Disabled      SharedPCConfigurationLocalStorage = "disabled"
	SharedPCConfigurationLocalStorage_Enabled       SharedPCConfigurationLocalStorage = "enabled"
	SharedPCConfigurationLocalStorage_NotConfigured SharedPCConfigurationLocalStorage = "notConfigured"
)

func PossibleValuesForSharedPCConfigurationLocalStorage() []string {
	return []string{
		string(SharedPCConfigurationLocalStorage_Disabled),
		string(SharedPCConfigurationLocalStorage_Enabled),
		string(SharedPCConfigurationLocalStorage_NotConfigured),
	}
}

func (s *SharedPCConfigurationLocalStorage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedPCConfigurationLocalStorage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharedPCConfigurationLocalStorage(input string) (*SharedPCConfigurationLocalStorage, error) {
	vals := map[string]SharedPCConfigurationLocalStorage{
		"disabled":      SharedPCConfigurationLocalStorage_Disabled,
		"enabled":       SharedPCConfigurationLocalStorage_Enabled,
		"notconfigured": SharedPCConfigurationLocalStorage_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCConfigurationLocalStorage(input)
	return &out, nil
}
