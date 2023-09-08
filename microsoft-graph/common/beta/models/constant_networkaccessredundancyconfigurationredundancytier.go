package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessRedundancyConfigurationRedundancyTier string

const (
	NetworkaccessRedundancyConfigurationRedundancyTiernoRedundancy   NetworkaccessRedundancyConfigurationRedundancyTier = "NoRedundancy"
	NetworkaccessRedundancyConfigurationRedundancyTierzoneRedundancy NetworkaccessRedundancyConfigurationRedundancyTier = "ZoneRedundancy"
)

func PossibleValuesForNetworkaccessRedundancyConfigurationRedundancyTier() []string {
	return []string{
		string(NetworkaccessRedundancyConfigurationRedundancyTiernoRedundancy),
		string(NetworkaccessRedundancyConfigurationRedundancyTierzoneRedundancy),
	}
}

func parseNetworkaccessRedundancyConfigurationRedundancyTier(input string) (*NetworkaccessRedundancyConfigurationRedundancyTier, error) {
	vals := map[string]NetworkaccessRedundancyConfigurationRedundancyTier{
		"noredundancy":   NetworkaccessRedundancyConfigurationRedundancyTiernoRedundancy,
		"zoneredundancy": NetworkaccessRedundancyConfigurationRedundancyTierzoneRedundancy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessRedundancyConfigurationRedundancyTier(input)
	return &out, nil
}
