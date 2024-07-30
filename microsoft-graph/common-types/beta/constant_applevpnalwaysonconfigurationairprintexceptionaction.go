package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnAlwaysOnConfigurationAirPrintExceptionAction string

const (
	AppleVpnAlwaysOnConfigurationAirPrintExceptionAction_AllowTrafficOutside AppleVpnAlwaysOnConfigurationAirPrintExceptionAction = "allowTrafficOutside"
	AppleVpnAlwaysOnConfigurationAirPrintExceptionAction_DropTraffic         AppleVpnAlwaysOnConfigurationAirPrintExceptionAction = "dropTraffic"
	AppleVpnAlwaysOnConfigurationAirPrintExceptionAction_ForceTrafficViaVPN  AppleVpnAlwaysOnConfigurationAirPrintExceptionAction = "forceTrafficViaVPN"
)

func PossibleValuesForAppleVpnAlwaysOnConfigurationAirPrintExceptionAction() []string {
	return []string{
		string(AppleVpnAlwaysOnConfigurationAirPrintExceptionAction_AllowTrafficOutside),
		string(AppleVpnAlwaysOnConfigurationAirPrintExceptionAction_DropTraffic),
		string(AppleVpnAlwaysOnConfigurationAirPrintExceptionAction_ForceTrafficViaVPN),
	}
}

func (s *AppleVpnAlwaysOnConfigurationAirPrintExceptionAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleVpnAlwaysOnConfigurationAirPrintExceptionAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleVpnAlwaysOnConfigurationAirPrintExceptionAction(input string) (*AppleVpnAlwaysOnConfigurationAirPrintExceptionAction, error) {
	vals := map[string]AppleVpnAlwaysOnConfigurationAirPrintExceptionAction{
		"allowtrafficoutside": AppleVpnAlwaysOnConfigurationAirPrintExceptionAction_AllowTrafficOutside,
		"droptraffic":         AppleVpnAlwaysOnConfigurationAirPrintExceptionAction_DropTraffic,
		"forcetrafficviavpn":  AppleVpnAlwaysOnConfigurationAirPrintExceptionAction_ForceTrafficViaVPN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnAlwaysOnConfigurationAirPrintExceptionAction(input)
	return &out, nil
}
