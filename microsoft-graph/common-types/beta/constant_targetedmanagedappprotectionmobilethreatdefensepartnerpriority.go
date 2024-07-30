package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionMobileThreatDefensePartnerPriority string

const (
	TargetedManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner TargetedManagedAppProtectionMobileThreatDefensePartnerPriority = "defenderOverThirdPartyPartner"
	TargetedManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender TargetedManagedAppProtectionMobileThreatDefensePartnerPriority = "thirdPartyPartnerOverDefender"
)

func PossibleValuesForTargetedManagedAppProtectionMobileThreatDefensePartnerPriority() []string {
	return []string{
		string(TargetedManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner),
		string(TargetedManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender),
	}
}

func (s *TargetedManagedAppProtectionMobileThreatDefensePartnerPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionMobileThreatDefensePartnerPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionMobileThreatDefensePartnerPriority(input string) (*TargetedManagedAppProtectionMobileThreatDefensePartnerPriority, error) {
	vals := map[string]TargetedManagedAppProtectionMobileThreatDefensePartnerPriority{
		"defenderoverthirdpartypartner": TargetedManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner,
		"thirdpartypartneroverdefender": TargetedManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionMobileThreatDefensePartnerPriority(input)
	return &out, nil
}
