package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HasPayloadLinkResultItemSources string

const (
	HasPayloadLinkResultItemSources_Direct     HasPayloadLinkResultItemSources = "direct"
	HasPayloadLinkResultItemSources_PolicySets HasPayloadLinkResultItemSources = "policySets"
)

func PossibleValuesForHasPayloadLinkResultItemSources() []string {
	return []string{
		string(HasPayloadLinkResultItemSources_Direct),
		string(HasPayloadLinkResultItemSources_PolicySets),
	}
}

func (s *HasPayloadLinkResultItemSources) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHasPayloadLinkResultItemSources(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHasPayloadLinkResultItemSources(input string) (*HasPayloadLinkResultItemSources, error) {
	vals := map[string]HasPayloadLinkResultItemSources{
		"direct":     HasPayloadLinkResultItemSources_Direct,
		"policysets": HasPayloadLinkResultItemSources_PolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HasPayloadLinkResultItemSources(input)
	return &out, nil
}
