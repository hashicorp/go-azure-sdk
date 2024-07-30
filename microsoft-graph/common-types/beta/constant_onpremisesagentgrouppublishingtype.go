package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAgentGroupPublishingType string

const (
	OnPremisesAgentGroupPublishingType_ApplicationProxy OnPremisesAgentGroupPublishingType = "applicationProxy"
	OnPremisesAgentGroupPublishingType_Authentication   OnPremisesAgentGroupPublishingType = "authentication"
	OnPremisesAgentGroupPublishingType_ExchangeOnline   OnPremisesAgentGroupPublishingType = "exchangeOnline"
	OnPremisesAgentGroupPublishingType_IntunePfx        OnPremisesAgentGroupPublishingType = "intunePfx"
	OnPremisesAgentGroupPublishingType_OflineDomainJoin OnPremisesAgentGroupPublishingType = "oflineDomainJoin"
	OnPremisesAgentGroupPublishingType_Provisioning     OnPremisesAgentGroupPublishingType = "provisioning"
)

func PossibleValuesForOnPremisesAgentGroupPublishingType() []string {
	return []string{
		string(OnPremisesAgentGroupPublishingType_ApplicationProxy),
		string(OnPremisesAgentGroupPublishingType_Authentication),
		string(OnPremisesAgentGroupPublishingType_ExchangeOnline),
		string(OnPremisesAgentGroupPublishingType_IntunePfx),
		string(OnPremisesAgentGroupPublishingType_OflineDomainJoin),
		string(OnPremisesAgentGroupPublishingType_Provisioning),
	}
}

func (s *OnPremisesAgentGroupPublishingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnPremisesAgentGroupPublishingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnPremisesAgentGroupPublishingType(input string) (*OnPremisesAgentGroupPublishingType, error) {
	vals := map[string]OnPremisesAgentGroupPublishingType{
		"applicationproxy": OnPremisesAgentGroupPublishingType_ApplicationProxy,
		"authentication":   OnPremisesAgentGroupPublishingType_Authentication,
		"exchangeonline":   OnPremisesAgentGroupPublishingType_ExchangeOnline,
		"intunepfx":        OnPremisesAgentGroupPublishingType_IntunePfx,
		"oflinedomainjoin": OnPremisesAgentGroupPublishingType_OflineDomainJoin,
		"provisioning":     OnPremisesAgentGroupPublishingType_Provisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesAgentGroupPublishingType(input)
	return &out, nil
}
