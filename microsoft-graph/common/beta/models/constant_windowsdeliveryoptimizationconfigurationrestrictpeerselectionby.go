package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy string

const (
	WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBynotConfigured WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy = "NotConfigured"
	WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBysubnetMask    WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy = "SubnetMask"
)

func PossibleValuesForWindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy() []string {
	return []string{
		string(WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBynotConfigured),
		string(WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBysubnetMask),
	}
}

func parseWindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy(input string) (*WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy, error) {
	vals := map[string]WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy{
		"notconfigured": WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBynotConfigured,
		"subnetmask":    WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBysubnetMask,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeliveryOptimizationConfigurationRestrictPeerSelectionBy(input)
	return &out, nil
}
