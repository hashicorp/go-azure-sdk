package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeliveryOptimizationConfigurationVpnPeerCaching string

const (
	WindowsDeliveryOptimizationConfigurationVpnPeerCachingdisabled      WindowsDeliveryOptimizationConfigurationVpnPeerCaching = "Disabled"
	WindowsDeliveryOptimizationConfigurationVpnPeerCachingenabled       WindowsDeliveryOptimizationConfigurationVpnPeerCaching = "Enabled"
	WindowsDeliveryOptimizationConfigurationVpnPeerCachingnotConfigured WindowsDeliveryOptimizationConfigurationVpnPeerCaching = "NotConfigured"
)

func PossibleValuesForWindowsDeliveryOptimizationConfigurationVpnPeerCaching() []string {
	return []string{
		string(WindowsDeliveryOptimizationConfigurationVpnPeerCachingdisabled),
		string(WindowsDeliveryOptimizationConfigurationVpnPeerCachingenabled),
		string(WindowsDeliveryOptimizationConfigurationVpnPeerCachingnotConfigured),
	}
}

func parseWindowsDeliveryOptimizationConfigurationVpnPeerCaching(input string) (*WindowsDeliveryOptimizationConfigurationVpnPeerCaching, error) {
	vals := map[string]WindowsDeliveryOptimizationConfigurationVpnPeerCaching{
		"disabled":      WindowsDeliveryOptimizationConfigurationVpnPeerCachingdisabled,
		"enabled":       WindowsDeliveryOptimizationConfigurationVpnPeerCachingenabled,
		"notconfigured": WindowsDeliveryOptimizationConfigurationVpnPeerCachingnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeliveryOptimizationConfigurationVpnPeerCaching(input)
	return &out, nil
}
