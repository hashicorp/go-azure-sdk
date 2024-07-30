package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy string

const (
	MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_FirstAvailable  MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy = "firstAvailable"
	MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_NotConfigured   MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy = "notConfigured"
	MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_Random          MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy = "random"
	MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_RoundRobin      MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy = "roundRobin"
	MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_StickyAvailable MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy = "stickyAvailable"
	MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_UrlPathHash     MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy = "urlPathHash"
)

func PossibleValuesForMacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy() []string {
	return []string{
		string(MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_FirstAvailable),
		string(MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_NotConfigured),
		string(MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_Random),
		string(MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_RoundRobin),
		string(MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_StickyAvailable),
		string(MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_UrlPathHash),
	}
}

func (s *MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy(input string) (*MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy, error) {
	vals := map[string]MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy{
		"firstavailable":  MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_FirstAvailable,
		"notconfigured":   MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_NotConfigured,
		"random":          MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_Random,
		"roundrobin":      MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_RoundRobin,
		"stickyavailable": MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_StickyAvailable,
		"urlpathhash":     MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy_UrlPathHash,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSDeviceFeaturesConfigurationContentCachingParentSelectionPolicy(input)
	return &out, nil
}
