package openapis

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapacitySkuTier string

const (
	CapacitySkuTierAutoPremiumHost CapacitySkuTier = "AutoPremiumHost"
	CapacitySkuTierPBIEAzure       CapacitySkuTier = "PBIE_Azure"
	CapacitySkuTierPremium         CapacitySkuTier = "Premium"
)

func PossibleValuesForCapacitySkuTier() []string {
	return []string{
		string(CapacitySkuTierAutoPremiumHost),
		string(CapacitySkuTierPBIEAzure),
		string(CapacitySkuTierPremium),
	}
}

func (s *CapacitySkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCapacitySkuTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCapacitySkuTier(input string) (*CapacitySkuTier, error) {
	vals := map[string]CapacitySkuTier{
		"autopremiumhost": CapacitySkuTierAutoPremiumHost,
		"pbie_azure":      CapacitySkuTierPBIEAzure,
		"premium":         CapacitySkuTierPremium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CapacitySkuTier(input)
	return &out, nil
}
