package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode string

const (
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModebypassMode                  WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "BypassMode"
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModehttpOnly                    WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "HttpOnly"
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModehttpWithInternetPeering     WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "HttpWithInternetPeering"
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModehttpWithPeeringNat          WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "HttpWithPeeringNat"
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModehttpWithPeeringPrivateGroup WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "HttpWithPeeringPrivateGroup"
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModesimpleDownload              WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "SimpleDownload"
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModeuserDefined                 WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "UserDefined"
)

func PossibleValuesForWindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode() []string {
	return []string{
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModebypassMode),
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModehttpOnly),
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModehttpWithInternetPeering),
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModehttpWithPeeringNat),
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModehttpWithPeeringPrivateGroup),
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModesimpleDownload),
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModeuserDefined),
	}
}

func parseWindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode(input string) (*WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode, error) {
	vals := map[string]WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode{
		"bypassmode":                  WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModebypassMode,
		"httponly":                    WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModehttpOnly,
		"httpwithinternetpeering":     WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModehttpWithInternetPeering,
		"httpwithpeeringnat":          WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModehttpWithPeeringNat,
		"httpwithpeeringprivategroup": WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModehttpWithPeeringPrivateGroup,
		"simpledownload":              WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModesimpleDownload,
		"userdefined":                 WindowsDeliveryOptimizationConfigurationDeliveryOptimizationModeuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode(input)
	return &out, nil
}
