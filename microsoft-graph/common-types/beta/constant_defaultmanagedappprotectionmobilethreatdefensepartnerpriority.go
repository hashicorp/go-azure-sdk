package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionMobileThreatDefensePartnerPriority string

const (
	DefaultManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner DefaultManagedAppProtectionMobileThreatDefensePartnerPriority = "defenderOverThirdPartyPartner"
	DefaultManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender DefaultManagedAppProtectionMobileThreatDefensePartnerPriority = "thirdPartyPartnerOverDefender"
)

func PossibleValuesForDefaultManagedAppProtectionMobileThreatDefensePartnerPriority() []string {
	return []string{
		string(DefaultManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner),
		string(DefaultManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender),
	}
}

func (s *DefaultManagedAppProtectionMobileThreatDefensePartnerPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionMobileThreatDefensePartnerPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionMobileThreatDefensePartnerPriority(input string) (*DefaultManagedAppProtectionMobileThreatDefensePartnerPriority, error) {
	vals := map[string]DefaultManagedAppProtectionMobileThreatDefensePartnerPriority{
		"defenderoverthirdpartypartner": DefaultManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner,
		"thirdpartypartneroverdefender": DefaultManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionMobileThreatDefensePartnerPriority(input)
	return &out, nil
}
