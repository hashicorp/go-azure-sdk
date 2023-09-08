package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAgentGroupPublishingType string

const (
	OnPremisesAgentGroupPublishingTypeapplicationProxy OnPremisesAgentGroupPublishingType = "ApplicationProxy"
	OnPremisesAgentGroupPublishingTypeauthentication   OnPremisesAgentGroupPublishingType = "Authentication"
	OnPremisesAgentGroupPublishingTypeexchangeOnline   OnPremisesAgentGroupPublishingType = "ExchangeOnline"
	OnPremisesAgentGroupPublishingTypeintunePfx        OnPremisesAgentGroupPublishingType = "IntunePfx"
	OnPremisesAgentGroupPublishingTypeoflineDomainJoin OnPremisesAgentGroupPublishingType = "OflineDomainJoin"
	OnPremisesAgentGroupPublishingTypeprovisioning     OnPremisesAgentGroupPublishingType = "Provisioning"
)

func PossibleValuesForOnPremisesAgentGroupPublishingType() []string {
	return []string{
		string(OnPremisesAgentGroupPublishingTypeapplicationProxy),
		string(OnPremisesAgentGroupPublishingTypeauthentication),
		string(OnPremisesAgentGroupPublishingTypeexchangeOnline),
		string(OnPremisesAgentGroupPublishingTypeintunePfx),
		string(OnPremisesAgentGroupPublishingTypeoflineDomainJoin),
		string(OnPremisesAgentGroupPublishingTypeprovisioning),
	}
}

func parseOnPremisesAgentGroupPublishingType(input string) (*OnPremisesAgentGroupPublishingType, error) {
	vals := map[string]OnPremisesAgentGroupPublishingType{
		"applicationproxy": OnPremisesAgentGroupPublishingTypeapplicationProxy,
		"authentication":   OnPremisesAgentGroupPublishingTypeauthentication,
		"exchangeonline":   OnPremisesAgentGroupPublishingTypeexchangeOnline,
		"intunepfx":        OnPremisesAgentGroupPublishingTypeintunePfx,
		"oflinedomainjoin": OnPremisesAgentGroupPublishingTypeoflineDomainJoin,
		"provisioning":     OnPremisesAgentGroupPublishingTypeprovisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesAgentGroupPublishingType(input)
	return &out, nil
}
