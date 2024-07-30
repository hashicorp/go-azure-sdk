package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnAlwaysOnConfigurationVoicemailExceptionAction string

const (
	AppleVpnAlwaysOnConfigurationVoicemailExceptionAction_AllowTrafficOutside AppleVpnAlwaysOnConfigurationVoicemailExceptionAction = "allowTrafficOutside"
	AppleVpnAlwaysOnConfigurationVoicemailExceptionAction_DropTraffic         AppleVpnAlwaysOnConfigurationVoicemailExceptionAction = "dropTraffic"
	AppleVpnAlwaysOnConfigurationVoicemailExceptionAction_ForceTrafficViaVPN  AppleVpnAlwaysOnConfigurationVoicemailExceptionAction = "forceTrafficViaVPN"
)

func PossibleValuesForAppleVpnAlwaysOnConfigurationVoicemailExceptionAction() []string {
	return []string{
		string(AppleVpnAlwaysOnConfigurationVoicemailExceptionAction_AllowTrafficOutside),
		string(AppleVpnAlwaysOnConfigurationVoicemailExceptionAction_DropTraffic),
		string(AppleVpnAlwaysOnConfigurationVoicemailExceptionAction_ForceTrafficViaVPN),
	}
}

func (s *AppleVpnAlwaysOnConfigurationVoicemailExceptionAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleVpnAlwaysOnConfigurationVoicemailExceptionAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleVpnAlwaysOnConfigurationVoicemailExceptionAction(input string) (*AppleVpnAlwaysOnConfigurationVoicemailExceptionAction, error) {
	vals := map[string]AppleVpnAlwaysOnConfigurationVoicemailExceptionAction{
		"allowtrafficoutside": AppleVpnAlwaysOnConfigurationVoicemailExceptionAction_AllowTrafficOutside,
		"droptraffic":         AppleVpnAlwaysOnConfigurationVoicemailExceptionAction_DropTraffic,
		"forcetrafficviavpn":  AppleVpnAlwaysOnConfigurationVoicemailExceptionAction_ForceTrafficViaVPN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnAlwaysOnConfigurationVoicemailExceptionAction(input)
	return &out, nil
}
