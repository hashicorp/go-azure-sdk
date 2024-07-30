package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserSourceIncludedSources string

const (
	SecurityUserSourceIncludedSources_Mailbox SecurityUserSourceIncludedSources = "mailbox"
	SecurityUserSourceIncludedSources_Site    SecurityUserSourceIncludedSources = "site"
)

func PossibleValuesForSecurityUserSourceIncludedSources() []string {
	return []string{
		string(SecurityUserSourceIncludedSources_Mailbox),
		string(SecurityUserSourceIncludedSources_Site),
	}
}

func (s *SecurityUserSourceIncludedSources) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUserSourceIncludedSources(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUserSourceIncludedSources(input string) (*SecurityUserSourceIncludedSources, error) {
	vals := map[string]SecurityUserSourceIncludedSources{
		"mailbox": SecurityUserSourceIncludedSources_Mailbox,
		"site":    SecurityUserSourceIncludedSources_Site,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUserSourceIncludedSources(input)
	return &out, nil
}
