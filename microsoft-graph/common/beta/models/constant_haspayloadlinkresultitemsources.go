package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HasPayloadLinkResultItemSources string

const (
	HasPayloadLinkResultItemSourcesdirect     HasPayloadLinkResultItemSources = "Direct"
	HasPayloadLinkResultItemSourcespolicySets HasPayloadLinkResultItemSources = "PolicySets"
)

func PossibleValuesForHasPayloadLinkResultItemSources() []string {
	return []string{
		string(HasPayloadLinkResultItemSourcesdirect),
		string(HasPayloadLinkResultItemSourcespolicySets),
	}
}

func parseHasPayloadLinkResultItemSources(input string) (*HasPayloadLinkResultItemSources, error) {
	vals := map[string]HasPayloadLinkResultItemSources{
		"direct":     HasPayloadLinkResultItemSourcesdirect,
		"policysets": HasPayloadLinkResultItemSourcespolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HasPayloadLinkResultItemSources(input)
	return &out, nil
}
