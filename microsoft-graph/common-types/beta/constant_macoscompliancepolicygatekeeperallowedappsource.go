package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCompliancePolicyGatekeeperAllowedAppSource string

const (
	MacOSCompliancePolicyGatekeeperAllowedAppSource_Anywhere                           MacOSCompliancePolicyGatekeeperAllowedAppSource = "anywhere"
	MacOSCompliancePolicyGatekeeperAllowedAppSource_MacAppStore                        MacOSCompliancePolicyGatekeeperAllowedAppSource = "macAppStore"
	MacOSCompliancePolicyGatekeeperAllowedAppSource_MacAppStoreAndIdentifiedDevelopers MacOSCompliancePolicyGatekeeperAllowedAppSource = "macAppStoreAndIdentifiedDevelopers"
	MacOSCompliancePolicyGatekeeperAllowedAppSource_NotConfigured                      MacOSCompliancePolicyGatekeeperAllowedAppSource = "notConfigured"
)

func PossibleValuesForMacOSCompliancePolicyGatekeeperAllowedAppSource() []string {
	return []string{
		string(MacOSCompliancePolicyGatekeeperAllowedAppSource_Anywhere),
		string(MacOSCompliancePolicyGatekeeperAllowedAppSource_MacAppStore),
		string(MacOSCompliancePolicyGatekeeperAllowedAppSource_MacAppStoreAndIdentifiedDevelopers),
		string(MacOSCompliancePolicyGatekeeperAllowedAppSource_NotConfigured),
	}
}

func (s *MacOSCompliancePolicyGatekeeperAllowedAppSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSCompliancePolicyGatekeeperAllowedAppSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSCompliancePolicyGatekeeperAllowedAppSource(input string) (*MacOSCompliancePolicyGatekeeperAllowedAppSource, error) {
	vals := map[string]MacOSCompliancePolicyGatekeeperAllowedAppSource{
		"anywhere":                           MacOSCompliancePolicyGatekeeperAllowedAppSource_Anywhere,
		"macappstore":                        MacOSCompliancePolicyGatekeeperAllowedAppSource_MacAppStore,
		"macappstoreandidentifieddevelopers": MacOSCompliancePolicyGatekeeperAllowedAppSource_MacAppStoreAndIdentifiedDevelopers,
		"notconfigured":                      MacOSCompliancePolicyGatekeeperAllowedAppSource_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCompliancePolicyGatekeeperAllowedAppSource(input)
	return &out, nil
}
