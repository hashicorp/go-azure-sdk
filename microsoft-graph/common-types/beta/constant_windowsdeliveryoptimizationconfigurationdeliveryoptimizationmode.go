package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode string

const (
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_BypassMode                  WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "bypassMode"
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_HttpOnly                    WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "httpOnly"
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_HttpWithInternetPeering     WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "httpWithInternetPeering"
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_HttpWithPeeringNat          WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "httpWithPeeringNat"
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_HttpWithPeeringPrivateGroup WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "httpWithPeeringPrivateGroup"
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_SimpleDownload              WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "simpleDownload"
	WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_UserDefined                 WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode = "userDefined"
)

func PossibleValuesForWindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode() []string {
	return []string{
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_BypassMode),
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_HttpOnly),
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_HttpWithInternetPeering),
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_HttpWithPeeringNat),
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_HttpWithPeeringPrivateGroup),
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_SimpleDownload),
		string(WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_UserDefined),
	}
}

func (s *WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode(input string) (*WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode, error) {
	vals := map[string]WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode{
		"bypassmode":                  WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_BypassMode,
		"httponly":                    WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_HttpOnly,
		"httpwithinternetpeering":     WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_HttpWithInternetPeering,
		"httpwithpeeringnat":          WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_HttpWithPeeringNat,
		"httpwithpeeringprivategroup": WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_HttpWithPeeringPrivateGroup,
		"simpledownload":              WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_SimpleDownload,
		"userdefined":                 WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeliveryOptimizationConfigurationDeliveryOptimizationMode(input)
	return &out, nil
}
