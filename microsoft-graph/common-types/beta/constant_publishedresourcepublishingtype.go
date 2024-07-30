package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublishedResourcePublishingType string

const (
	PublishedResourcePublishingType_ApplicationProxy PublishedResourcePublishingType = "applicationProxy"
	PublishedResourcePublishingType_Authentication   PublishedResourcePublishingType = "authentication"
	PublishedResourcePublishingType_ExchangeOnline   PublishedResourcePublishingType = "exchangeOnline"
	PublishedResourcePublishingType_IntunePfx        PublishedResourcePublishingType = "intunePfx"
	PublishedResourcePublishingType_OflineDomainJoin PublishedResourcePublishingType = "oflineDomainJoin"
	PublishedResourcePublishingType_Provisioning     PublishedResourcePublishingType = "provisioning"
)

func PossibleValuesForPublishedResourcePublishingType() []string {
	return []string{
		string(PublishedResourcePublishingType_ApplicationProxy),
		string(PublishedResourcePublishingType_Authentication),
		string(PublishedResourcePublishingType_ExchangeOnline),
		string(PublishedResourcePublishingType_IntunePfx),
		string(PublishedResourcePublishingType_OflineDomainJoin),
		string(PublishedResourcePublishingType_Provisioning),
	}
}

func (s *PublishedResourcePublishingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePublishedResourcePublishingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePublishedResourcePublishingType(input string) (*PublishedResourcePublishingType, error) {
	vals := map[string]PublishedResourcePublishingType{
		"applicationproxy": PublishedResourcePublishingType_ApplicationProxy,
		"authentication":   PublishedResourcePublishingType_Authentication,
		"exchangeonline":   PublishedResourcePublishingType_ExchangeOnline,
		"intunepfx":        PublishedResourcePublishingType_IntunePfx,
		"oflinedomainjoin": PublishedResourcePublishingType_OflineDomainJoin,
		"provisioning":     PublishedResourcePublishingType_Provisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublishedResourcePublishingType(input)
	return &out, nil
}
