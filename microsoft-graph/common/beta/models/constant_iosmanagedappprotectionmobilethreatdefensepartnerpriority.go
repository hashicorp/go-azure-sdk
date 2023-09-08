package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionMobileThreatDefensePartnerPriority string

const (
	IosManagedAppProtectionMobileThreatDefensePartnerPrioritydefenderOverThirdPartyPartner IosManagedAppProtectionMobileThreatDefensePartnerPriority = "DefenderOverThirdPartyPartner"
	IosManagedAppProtectionMobileThreatDefensePartnerPrioritythirdPartyPartnerOverDefender IosManagedAppProtectionMobileThreatDefensePartnerPriority = "ThirdPartyPartnerOverDefender"
)

func PossibleValuesForIosManagedAppProtectionMobileThreatDefensePartnerPriority() []string {
	return []string{
		string(IosManagedAppProtectionMobileThreatDefensePartnerPrioritydefenderOverThirdPartyPartner),
		string(IosManagedAppProtectionMobileThreatDefensePartnerPrioritythirdPartyPartnerOverDefender),
	}
}

func parseIosManagedAppProtectionMobileThreatDefensePartnerPriority(input string) (*IosManagedAppProtectionMobileThreatDefensePartnerPriority, error) {
	vals := map[string]IosManagedAppProtectionMobileThreatDefensePartnerPriority{
		"defenderoverthirdpartypartner": IosManagedAppProtectionMobileThreatDefensePartnerPrioritydefenderOverThirdPartyPartner,
		"thirdpartypartneroverdefender": IosManagedAppProtectionMobileThreatDefensePartnerPrioritythirdPartyPartnerOverDefender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionMobileThreatDefensePartnerPriority(input)
	return &out, nil
}
