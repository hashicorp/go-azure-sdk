package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption string

const (
	DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_AdSite                 DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption = "adSite"
	DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_AuthenticatedDomainSid DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption = "authenticatedDomainSid"
	DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_DhcpUserOption         DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption = "dhcpUserOption"
	DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_DnsSuffix              DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption = "dnsSuffix"
	DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_NotConfigured          DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption = "notConfigured"
)

func PossibleValuesForDeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption() []string {
	return []string{
		string(DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_AdSite),
		string(DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_AuthenticatedDomainSid),
		string(DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_DhcpUserOption),
		string(DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_DnsSuffix),
		string(DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_NotConfigured),
	}
}

func (s *DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption(input string) (*DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption, error) {
	vals := map[string]DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption{
		"adsite":                 DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_AdSite,
		"authenticateddomainsid": DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_AuthenticatedDomainSid,
		"dhcpuseroption":         DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_DhcpUserOption,
		"dnssuffix":              DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_DnsSuffix,
		"notconfigured":          DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption(input)
	return &out, nil
}
