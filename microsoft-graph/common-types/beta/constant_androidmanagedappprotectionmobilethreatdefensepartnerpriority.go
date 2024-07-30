package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionMobileThreatDefensePartnerPriority string

const (
	AndroidManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner AndroidManagedAppProtectionMobileThreatDefensePartnerPriority = "defenderOverThirdPartyPartner"
	AndroidManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender AndroidManagedAppProtectionMobileThreatDefensePartnerPriority = "thirdPartyPartnerOverDefender"
)

func PossibleValuesForAndroidManagedAppProtectionMobileThreatDefensePartnerPriority() []string {
	return []string{
		string(AndroidManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner),
		string(AndroidManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender),
	}
}

func (s *AndroidManagedAppProtectionMobileThreatDefensePartnerPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionMobileThreatDefensePartnerPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionMobileThreatDefensePartnerPriority(input string) (*AndroidManagedAppProtectionMobileThreatDefensePartnerPriority, error) {
	vals := map[string]AndroidManagedAppProtectionMobileThreatDefensePartnerPriority{
		"defenderoverthirdpartypartner": AndroidManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner,
		"thirdpartypartneroverdefender": AndroidManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionMobileThreatDefensePartnerPriority(input)
	return &out, nil
}
