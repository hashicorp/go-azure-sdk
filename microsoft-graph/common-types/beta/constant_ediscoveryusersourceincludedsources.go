package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryUserSourceIncludedSources string

const (
	EdiscoveryUserSourceIncludedSources_Mailbox EdiscoveryUserSourceIncludedSources = "mailbox"
	EdiscoveryUserSourceIncludedSources_Site    EdiscoveryUserSourceIncludedSources = "site"
)

func PossibleValuesForEdiscoveryUserSourceIncludedSources() []string {
	return []string{
		string(EdiscoveryUserSourceIncludedSources_Mailbox),
		string(EdiscoveryUserSourceIncludedSources_Site),
	}
}

func (s *EdiscoveryUserSourceIncludedSources) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryUserSourceIncludedSources(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryUserSourceIncludedSources(input string) (*EdiscoveryUserSourceIncludedSources, error) {
	vals := map[string]EdiscoveryUserSourceIncludedSources{
		"mailbox": EdiscoveryUserSourceIncludedSources_Mailbox,
		"site":    EdiscoveryUserSourceIncludedSources_Site,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryUserSourceIncludedSources(input)
	return &out, nil
}
