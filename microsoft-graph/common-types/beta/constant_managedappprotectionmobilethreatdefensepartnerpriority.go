package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionMobileThreatDefensePartnerPriority string

const (
	ManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner ManagedAppProtectionMobileThreatDefensePartnerPriority = "defenderOverThirdPartyPartner"
	ManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender ManagedAppProtectionMobileThreatDefensePartnerPriority = "thirdPartyPartnerOverDefender"
)

func PossibleValuesForManagedAppProtectionMobileThreatDefensePartnerPriority() []string {
	return []string{
		string(ManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner),
		string(ManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender),
	}
}

func (s *ManagedAppProtectionMobileThreatDefensePartnerPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionMobileThreatDefensePartnerPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionMobileThreatDefensePartnerPriority(input string) (*ManagedAppProtectionMobileThreatDefensePartnerPriority, error) {
	vals := map[string]ManagedAppProtectionMobileThreatDefensePartnerPriority{
		"defenderoverthirdpartypartner": ManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner,
		"thirdpartypartneroverdefender": ManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionMobileThreatDefensePartnerPriority(input)
	return &out, nil
}
