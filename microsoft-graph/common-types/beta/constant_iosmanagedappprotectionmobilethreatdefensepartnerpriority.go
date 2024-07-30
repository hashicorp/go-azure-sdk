package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionMobileThreatDefensePartnerPriority string

const (
	IosManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner IosManagedAppProtectionMobileThreatDefensePartnerPriority = "defenderOverThirdPartyPartner"
	IosManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender IosManagedAppProtectionMobileThreatDefensePartnerPriority = "thirdPartyPartnerOverDefender"
)

func PossibleValuesForIosManagedAppProtectionMobileThreatDefensePartnerPriority() []string {
	return []string{
		string(IosManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner),
		string(IosManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender),
	}
}

func (s *IosManagedAppProtectionMobileThreatDefensePartnerPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionMobileThreatDefensePartnerPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionMobileThreatDefensePartnerPriority(input string) (*IosManagedAppProtectionMobileThreatDefensePartnerPriority, error) {
	vals := map[string]IosManagedAppProtectionMobileThreatDefensePartnerPriority{
		"defenderoverthirdpartypartner": IosManagedAppProtectionMobileThreatDefensePartnerPriority_DefenderOverThirdPartyPartner,
		"thirdpartypartneroverdefender": IosManagedAppProtectionMobileThreatDefensePartnerPriority_ThirdPartyPartnerOverDefender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionMobileThreatDefensePartnerPriority(input)
	return &out, nil
}
