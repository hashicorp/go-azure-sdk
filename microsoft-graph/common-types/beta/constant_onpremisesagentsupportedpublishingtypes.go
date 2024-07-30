package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAgentSupportedPublishingTypes string

const (
	OnPremisesAgentSupportedPublishingTypes_ApplicationProxy OnPremisesAgentSupportedPublishingTypes = "applicationProxy"
	OnPremisesAgentSupportedPublishingTypes_Authentication   OnPremisesAgentSupportedPublishingTypes = "authentication"
	OnPremisesAgentSupportedPublishingTypes_ExchangeOnline   OnPremisesAgentSupportedPublishingTypes = "exchangeOnline"
	OnPremisesAgentSupportedPublishingTypes_IntunePfx        OnPremisesAgentSupportedPublishingTypes = "intunePfx"
	OnPremisesAgentSupportedPublishingTypes_OflineDomainJoin OnPremisesAgentSupportedPublishingTypes = "oflineDomainJoin"
	OnPremisesAgentSupportedPublishingTypes_Provisioning     OnPremisesAgentSupportedPublishingTypes = "provisioning"
)

func PossibleValuesForOnPremisesAgentSupportedPublishingTypes() []string {
	return []string{
		string(OnPremisesAgentSupportedPublishingTypes_ApplicationProxy),
		string(OnPremisesAgentSupportedPublishingTypes_Authentication),
		string(OnPremisesAgentSupportedPublishingTypes_ExchangeOnline),
		string(OnPremisesAgentSupportedPublishingTypes_IntunePfx),
		string(OnPremisesAgentSupportedPublishingTypes_OflineDomainJoin),
		string(OnPremisesAgentSupportedPublishingTypes_Provisioning),
	}
}

func (s *OnPremisesAgentSupportedPublishingTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnPremisesAgentSupportedPublishingTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnPremisesAgentSupportedPublishingTypes(input string) (*OnPremisesAgentSupportedPublishingTypes, error) {
	vals := map[string]OnPremisesAgentSupportedPublishingTypes{
		"applicationproxy": OnPremisesAgentSupportedPublishingTypes_ApplicationProxy,
		"authentication":   OnPremisesAgentSupportedPublishingTypes_Authentication,
		"exchangeonline":   OnPremisesAgentSupportedPublishingTypes_ExchangeOnline,
		"intunepfx":        OnPremisesAgentSupportedPublishingTypes_IntunePfx,
		"oflinedomainjoin": OnPremisesAgentSupportedPublishingTypes_OflineDomainJoin,
		"provisioning":     OnPremisesAgentSupportedPublishingTypes_Provisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesAgentSupportedPublishingTypes(input)
	return &out, nil
}
