package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnAlwaysOnConfigurationCellularExceptionAction string

const (
	AppleVpnAlwaysOnConfigurationCellularExceptionAction_AllowTrafficOutside AppleVpnAlwaysOnConfigurationCellularExceptionAction = "allowTrafficOutside"
	AppleVpnAlwaysOnConfigurationCellularExceptionAction_DropTraffic         AppleVpnAlwaysOnConfigurationCellularExceptionAction = "dropTraffic"
	AppleVpnAlwaysOnConfigurationCellularExceptionAction_ForceTrafficViaVPN  AppleVpnAlwaysOnConfigurationCellularExceptionAction = "forceTrafficViaVPN"
)

func PossibleValuesForAppleVpnAlwaysOnConfigurationCellularExceptionAction() []string {
	return []string{
		string(AppleVpnAlwaysOnConfigurationCellularExceptionAction_AllowTrafficOutside),
		string(AppleVpnAlwaysOnConfigurationCellularExceptionAction_DropTraffic),
		string(AppleVpnAlwaysOnConfigurationCellularExceptionAction_ForceTrafficViaVPN),
	}
}

func (s *AppleVpnAlwaysOnConfigurationCellularExceptionAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleVpnAlwaysOnConfigurationCellularExceptionAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleVpnAlwaysOnConfigurationCellularExceptionAction(input string) (*AppleVpnAlwaysOnConfigurationCellularExceptionAction, error) {
	vals := map[string]AppleVpnAlwaysOnConfigurationCellularExceptionAction{
		"allowtrafficoutside": AppleVpnAlwaysOnConfigurationCellularExceptionAction_AllowTrafficOutside,
		"droptraffic":         AppleVpnAlwaysOnConfigurationCellularExceptionAction_DropTraffic,
		"forcetrafficviavpn":  AppleVpnAlwaysOnConfigurationCellularExceptionAction_ForceTrafficViaVPN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnAlwaysOnConfigurationCellularExceptionAction(input)
	return &out, nil
}
