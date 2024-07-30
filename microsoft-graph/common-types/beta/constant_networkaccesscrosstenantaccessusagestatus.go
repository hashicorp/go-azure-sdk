package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessCrossTenantAccessUsageStatus string

const (
	NetworkaccessCrossTenantAccessUsageStatus_FrequentlyUsed NetworkaccessCrossTenantAccessUsageStatus = "frequentlyUsed"
	NetworkaccessCrossTenantAccessUsageStatus_RarelyUsed     NetworkaccessCrossTenantAccessUsageStatus = "rarelyUsed"
)

func PossibleValuesForNetworkaccessCrossTenantAccessUsageStatus() []string {
	return []string{
		string(NetworkaccessCrossTenantAccessUsageStatus_FrequentlyUsed),
		string(NetworkaccessCrossTenantAccessUsageStatus_RarelyUsed),
	}
}

func (s *NetworkaccessCrossTenantAccessUsageStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessCrossTenantAccessUsageStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessCrossTenantAccessUsageStatus(input string) (*NetworkaccessCrossTenantAccessUsageStatus, error) {
	vals := map[string]NetworkaccessCrossTenantAccessUsageStatus{
		"frequentlyused": NetworkaccessCrossTenantAccessUsageStatus_FrequentlyUsed,
		"rarelyused":     NetworkaccessCrossTenantAccessUsageStatus_RarelyUsed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessCrossTenantAccessUsageStatus(input)
	return &out, nil
}
