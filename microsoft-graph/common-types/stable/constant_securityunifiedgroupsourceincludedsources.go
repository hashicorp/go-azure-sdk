package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUnifiedGroupSourceIncludedSources string

const (
	SecurityUnifiedGroupSourceIncludedSources_Mailbox SecurityUnifiedGroupSourceIncludedSources = "mailbox"
	SecurityUnifiedGroupSourceIncludedSources_Site    SecurityUnifiedGroupSourceIncludedSources = "site"
)

func PossibleValuesForSecurityUnifiedGroupSourceIncludedSources() []string {
	return []string{
		string(SecurityUnifiedGroupSourceIncludedSources_Mailbox),
		string(SecurityUnifiedGroupSourceIncludedSources_Site),
	}
}

func (s *SecurityUnifiedGroupSourceIncludedSources) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUnifiedGroupSourceIncludedSources(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUnifiedGroupSourceIncludedSources(input string) (*SecurityUnifiedGroupSourceIncludedSources, error) {
	vals := map[string]SecurityUnifiedGroupSourceIncludedSources{
		"mailbox": SecurityUnifiedGroupSourceIncludedSources_Mailbox,
		"site":    SecurityUnifiedGroupSourceIncludedSources_Site,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUnifiedGroupSourceIncludedSources(input)
	return &out, nil
}
