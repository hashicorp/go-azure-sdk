package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileThreatDefenseConnectorPartnerState string

const (
	MobileThreatDefenseConnectorPartnerStateavailable    MobileThreatDefenseConnectorPartnerState = "Available"
	MobileThreatDefenseConnectorPartnerStateenabled      MobileThreatDefenseConnectorPartnerState = "Enabled"
	MobileThreatDefenseConnectorPartnerStateunavailable  MobileThreatDefenseConnectorPartnerState = "Unavailable"
	MobileThreatDefenseConnectorPartnerStateunresponsive MobileThreatDefenseConnectorPartnerState = "Unresponsive"
)

func PossibleValuesForMobileThreatDefenseConnectorPartnerState() []string {
	return []string{
		string(MobileThreatDefenseConnectorPartnerStateavailable),
		string(MobileThreatDefenseConnectorPartnerStateenabled),
		string(MobileThreatDefenseConnectorPartnerStateunavailable),
		string(MobileThreatDefenseConnectorPartnerStateunresponsive),
	}
}

func parseMobileThreatDefenseConnectorPartnerState(input string) (*MobileThreatDefenseConnectorPartnerState, error) {
	vals := map[string]MobileThreatDefenseConnectorPartnerState{
		"available":    MobileThreatDefenseConnectorPartnerStateavailable,
		"enabled":      MobileThreatDefenseConnectorPartnerStateenabled,
		"unavailable":  MobileThreatDefenseConnectorPartnerStateunavailable,
		"unresponsive": MobileThreatDefenseConnectorPartnerStateunresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileThreatDefenseConnectorPartnerState(input)
	return &out, nil
}
