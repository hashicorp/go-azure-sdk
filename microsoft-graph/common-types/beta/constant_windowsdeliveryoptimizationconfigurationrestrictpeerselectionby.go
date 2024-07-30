package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy string

const (
	WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy_NotConfigured WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy = "notConfigured"
	WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy_SubnetMask    WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy = "subnetMask"
)

func PossibleValuesForWindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy() []string {
	return []string{
		string(WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy_NotConfigured),
		string(WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy_SubnetMask),
	}
}

func (s *WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy(input string) (*WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy, error) {
	vals := map[string]WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy{
		"notconfigured": WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy_NotConfigured,
		"subnetmask":    WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy_SubnetMask,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy(input)
	return &out, nil
}
