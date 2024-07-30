package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode string

const (
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_BypassMode                  WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "bypassMode"
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_HttpOnly                    WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "httpOnly"
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_HttpWithInternetPeering     WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "httpWithInternetPeering"
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_HttpWithPeeringNat          WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "httpWithPeeringNat"
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_HttpWithPeeringPrivateGroup WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "httpWithPeeringPrivateGroup"
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_SimpleDownload              WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "simpleDownload"
	WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_UserDefined                 WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode = "userDefined"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationDeliveryOptimizationMode() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_BypassMode),
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_HttpOnly),
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_HttpWithInternetPeering),
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_HttpWithPeeringNat),
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_HttpWithPeeringPrivateGroup),
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_SimpleDownload),
		string(WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_UserDefined),
	}
}

func (s *WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateForBusinessConfigurationDeliveryOptimizationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateForBusinessConfigurationDeliveryOptimizationMode(input string) (*WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode{
		"bypassmode":                  WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_BypassMode,
		"httponly":                    WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_HttpOnly,
		"httpwithinternetpeering":     WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_HttpWithInternetPeering,
		"httpwithpeeringnat":          WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_HttpWithPeeringNat,
		"httpwithpeeringprivategroup": WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_HttpWithPeeringPrivateGroup,
		"simpledownload":              WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_SimpleDownload,
		"userdefined":                 WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode(input)
	return &out, nil
}
