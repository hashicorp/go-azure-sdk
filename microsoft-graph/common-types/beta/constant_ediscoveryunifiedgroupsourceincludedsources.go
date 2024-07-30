package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryUnifiedGroupSourceIncludedSources string

const (
	EdiscoveryUnifiedGroupSourceIncludedSources_Mailbox EdiscoveryUnifiedGroupSourceIncludedSources = "mailbox"
	EdiscoveryUnifiedGroupSourceIncludedSources_Site    EdiscoveryUnifiedGroupSourceIncludedSources = "site"
)

func PossibleValuesForEdiscoveryUnifiedGroupSourceIncludedSources() []string {
	return []string{
		string(EdiscoveryUnifiedGroupSourceIncludedSources_Mailbox),
		string(EdiscoveryUnifiedGroupSourceIncludedSources_Site),
	}
}

func (s *EdiscoveryUnifiedGroupSourceIncludedSources) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryUnifiedGroupSourceIncludedSources(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryUnifiedGroupSourceIncludedSources(input string) (*EdiscoveryUnifiedGroupSourceIncludedSources, error) {
	vals := map[string]EdiscoveryUnifiedGroupSourceIncludedSources{
		"mailbox": EdiscoveryUnifiedGroupSourceIncludedSources_Mailbox,
		"site":    EdiscoveryUnifiedGroupSourceIncludedSources_Site,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryUnifiedGroupSourceIncludedSources(input)
	return &out, nil
}
