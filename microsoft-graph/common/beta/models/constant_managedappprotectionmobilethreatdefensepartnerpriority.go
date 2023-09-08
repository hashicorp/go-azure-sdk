package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionMobileThreatDefensePartnerPriority string

const (
	ManagedAppProtectionMobileThreatDefensePartnerPrioritydefenderOverThirdPartyPartner ManagedAppProtectionMobileThreatDefensePartnerPriority = "DefenderOverThirdPartyPartner"
	ManagedAppProtectionMobileThreatDefensePartnerPrioritythirdPartyPartnerOverDefender ManagedAppProtectionMobileThreatDefensePartnerPriority = "ThirdPartyPartnerOverDefender"
)

func PossibleValuesForManagedAppProtectionMobileThreatDefensePartnerPriority() []string {
	return []string{
		string(ManagedAppProtectionMobileThreatDefensePartnerPrioritydefenderOverThirdPartyPartner),
		string(ManagedAppProtectionMobileThreatDefensePartnerPrioritythirdPartyPartnerOverDefender),
	}
}

func parseManagedAppProtectionMobileThreatDefensePartnerPriority(input string) (*ManagedAppProtectionMobileThreatDefensePartnerPriority, error) {
	vals := map[string]ManagedAppProtectionMobileThreatDefensePartnerPriority{
		"defenderoverthirdpartypartner": ManagedAppProtectionMobileThreatDefensePartnerPrioritydefenderOverThirdPartyPartner,
		"thirdpartypartneroverdefender": ManagedAppProtectionMobileThreatDefensePartnerPrioritythirdPartyPartnerOverDefender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionMobileThreatDefensePartnerPriority(input)
	return &out, nil
}
