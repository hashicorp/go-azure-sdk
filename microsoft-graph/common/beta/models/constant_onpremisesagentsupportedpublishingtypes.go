package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAgentSupportedPublishingTypes string

const (
	OnPremisesAgentSupportedPublishingTypesapplicationProxy OnPremisesAgentSupportedPublishingTypes = "ApplicationProxy"
	OnPremisesAgentSupportedPublishingTypesauthentication   OnPremisesAgentSupportedPublishingTypes = "Authentication"
	OnPremisesAgentSupportedPublishingTypesexchangeOnline   OnPremisesAgentSupportedPublishingTypes = "ExchangeOnline"
	OnPremisesAgentSupportedPublishingTypesintunePfx        OnPremisesAgentSupportedPublishingTypes = "IntunePfx"
	OnPremisesAgentSupportedPublishingTypesoflineDomainJoin OnPremisesAgentSupportedPublishingTypes = "OflineDomainJoin"
	OnPremisesAgentSupportedPublishingTypesprovisioning     OnPremisesAgentSupportedPublishingTypes = "Provisioning"
)

func PossibleValuesForOnPremisesAgentSupportedPublishingTypes() []string {
	return []string{
		string(OnPremisesAgentSupportedPublishingTypesapplicationProxy),
		string(OnPremisesAgentSupportedPublishingTypesauthentication),
		string(OnPremisesAgentSupportedPublishingTypesexchangeOnline),
		string(OnPremisesAgentSupportedPublishingTypesintunePfx),
		string(OnPremisesAgentSupportedPublishingTypesoflineDomainJoin),
		string(OnPremisesAgentSupportedPublishingTypesprovisioning),
	}
}

func parseOnPremisesAgentSupportedPublishingTypes(input string) (*OnPremisesAgentSupportedPublishingTypes, error) {
	vals := map[string]OnPremisesAgentSupportedPublishingTypes{
		"applicationproxy": OnPremisesAgentSupportedPublishingTypesapplicationProxy,
		"authentication":   OnPremisesAgentSupportedPublishingTypesauthentication,
		"exchangeonline":   OnPremisesAgentSupportedPublishingTypesexchangeOnline,
		"intunepfx":        OnPremisesAgentSupportedPublishingTypesintunePfx,
		"oflinedomainjoin": OnPremisesAgentSupportedPublishingTypesoflineDomainJoin,
		"provisioning":     OnPremisesAgentSupportedPublishingTypesprovisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesAgentSupportedPublishingTypes(input)
	return &out, nil
}
