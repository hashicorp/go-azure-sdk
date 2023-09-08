package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionMobileThreatDefensePartnerPriority string

const (
	AndroidManagedAppProtectionMobileThreatDefensePartnerPrioritydefenderOverThirdPartyPartner AndroidManagedAppProtectionMobileThreatDefensePartnerPriority = "DefenderOverThirdPartyPartner"
	AndroidManagedAppProtectionMobileThreatDefensePartnerPrioritythirdPartyPartnerOverDefender AndroidManagedAppProtectionMobileThreatDefensePartnerPriority = "ThirdPartyPartnerOverDefender"
)

func PossibleValuesForAndroidManagedAppProtectionMobileThreatDefensePartnerPriority() []string {
	return []string{
		string(AndroidManagedAppProtectionMobileThreatDefensePartnerPrioritydefenderOverThirdPartyPartner),
		string(AndroidManagedAppProtectionMobileThreatDefensePartnerPrioritythirdPartyPartnerOverDefender),
	}
}

func parseAndroidManagedAppProtectionMobileThreatDefensePartnerPriority(input string) (*AndroidManagedAppProtectionMobileThreatDefensePartnerPriority, error) {
	vals := map[string]AndroidManagedAppProtectionMobileThreatDefensePartnerPriority{
		"defenderoverthirdpartypartner": AndroidManagedAppProtectionMobileThreatDefensePartnerPrioritydefenderOverThirdPartyPartner,
		"thirdpartypartneroverdefender": AndroidManagedAppProtectionMobileThreatDefensePartnerPrioritythirdPartyPartnerOverDefender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionMobileThreatDefensePartnerPriority(input)
	return &out, nil
}
