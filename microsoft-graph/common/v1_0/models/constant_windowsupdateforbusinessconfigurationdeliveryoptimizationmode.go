package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode string

const (
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationModebypassMode                  WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "BypassMode"
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationModehttpOnly                    WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "HttpOnly"
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationModehttpWithInternetPeering     WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "HttpWithInternetPeering"
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationModehttpWithPeeringNat          WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "HttpWithPeeringNat"
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationModehttpWithPeeringPrivateGroup WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "HttpWithPeeringPrivateGroup"
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationModesimpleDownload              WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "SimpleDownload"
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationModeuserDefined                 WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "UserDefined"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationDeliveryOptimizationMode() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationModebypassMode),
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationModehttpOnly),
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationModehttpWithInternetPeering),
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationModehttpWithPeeringNat),
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationModehttpWithPeeringPrivateGroup),
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationModesimpleDownload),
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationModeuserDefined),
	}
}

func parseWindowsUpdateForBusinessConfigurationDeliveryOptimizationMode(input string) (*WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode{
		"bypassmode":                  WindowsUpdateForBusinessConfigurationDeliveryOptimizationModebypassMode,
		"httponly":                    WindowsUpdateForBusinessConfigurationDeliveryOptimizationModehttpOnly,
		"httpwithinternetpeering":     WindowsUpdateForBusinessConfigurationDeliveryOptimizationModehttpWithInternetPeering,
		"httpwithpeeringnat":          WindowsUpdateForBusinessConfigurationDeliveryOptimizationModehttpWithPeeringNat,
		"httpwithpeeringprivategroup": WindowsUpdateForBusinessConfigurationDeliveryOptimizationModehttpWithPeeringPrivateGroup,
		"simpledownload":              WindowsUpdateForBusinessConfigurationDeliveryOptimizationModesimpleDownload,
		"userdefined":                 WindowsUpdateForBusinessConfigurationDeliveryOptimizationModeuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode(input)
	return &out, nil
}
