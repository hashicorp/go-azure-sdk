package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionMobileThreatDefensePartnerPriority string

const (
	DefaultManagedAppProtectionMobileThreatDefensePartnerPrioritydefenderOverThirdPartyPartner DefaultManagedAppProtectionMobileThreatDefensePartnerPriority = "DefenderOverThirdPartyPartner"
	DefaultManagedAppProtectionMobileThreatDefensePartnerPrioritythirdPartyPartnerOverDefender DefaultManagedAppProtectionMobileThreatDefensePartnerPriority = "ThirdPartyPartnerOverDefender"
)

func PossibleValuesForDefaultManagedAppProtectionMobileThreatDefensePartnerPriority() []string {
	return []string{
		string(DefaultManagedAppProtectionMobileThreatDefensePartnerPrioritydefenderOverThirdPartyPartner),
		string(DefaultManagedAppProtectionMobileThreatDefensePartnerPrioritythirdPartyPartnerOverDefender),
	}
}

func parseDefaultManagedAppProtectionMobileThreatDefensePartnerPriority(input string) (*DefaultManagedAppProtectionMobileThreatDefensePartnerPriority, error) {
	vals := map[string]DefaultManagedAppProtectionMobileThreatDefensePartnerPriority{
		"defenderoverthirdpartypartner": DefaultManagedAppProtectionMobileThreatDefensePartnerPrioritydefenderOverThirdPartyPartner,
		"thirdpartypartneroverdefender": DefaultManagedAppProtectionMobileThreatDefensePartnerPrioritythirdPartyPartnerOverDefender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionMobileThreatDefensePartnerPriority(input)
	return &out, nil
}
