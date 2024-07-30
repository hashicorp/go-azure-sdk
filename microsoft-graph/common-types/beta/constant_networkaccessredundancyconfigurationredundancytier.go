package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessRedundancyConfigurationRedundancyTier string

const (
	NetworkaccessRedundancyConfigurationRedundancyTier_NoRedundancy   NetworkaccessRedundancyConfigurationRedundancyTier = "noRedundancy"
	NetworkaccessRedundancyConfigurationRedundancyTier_ZoneRedundancy NetworkaccessRedundancyConfigurationRedundancyTier = "zoneRedundancy"
)

func PossibleValuesForNetworkaccessRedundancyConfigurationRedundancyTier() []string {
	return []string{
		string(NetworkaccessRedundancyConfigurationRedundancyTier_NoRedundancy),
		string(NetworkaccessRedundancyConfigurationRedundancyTier_ZoneRedundancy),
	}
}

func (s *NetworkaccessRedundancyConfigurationRedundancyTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessRedundancyConfigurationRedundancyTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessRedundancyConfigurationRedundancyTier(input string) (*NetworkaccessRedundancyConfigurationRedundancyTier, error) {
	vals := map[string]NetworkaccessRedundancyConfigurationRedundancyTier{
		"noredundancy":   NetworkaccessRedundancyConfigurationRedundancyTier_NoRedundancy,
		"zoneredundancy": NetworkaccessRedundancyConfigurationRedundancyTier_ZoneRedundancy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessRedundancyConfigurationRedundancyTier(input)
	return &out, nil
}
