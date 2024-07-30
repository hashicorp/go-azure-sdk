package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIntelligenceProfileIndicatorSource string

const (
	SecurityIntelligenceProfileIndicatorSource_Microsoft SecurityIntelligenceProfileIndicatorSource = "microsoft"
	SecurityIntelligenceProfileIndicatorSource_Osint     SecurityIntelligenceProfileIndicatorSource = "osint"
	SecurityIntelligenceProfileIndicatorSource_Public    SecurityIntelligenceProfileIndicatorSource = "public"
)

func PossibleValuesForSecurityIntelligenceProfileIndicatorSource() []string {
	return []string{
		string(SecurityIntelligenceProfileIndicatorSource_Microsoft),
		string(SecurityIntelligenceProfileIndicatorSource_Osint),
		string(SecurityIntelligenceProfileIndicatorSource_Public),
	}
}

func (s *SecurityIntelligenceProfileIndicatorSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIntelligenceProfileIndicatorSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIntelligenceProfileIndicatorSource(input string) (*SecurityIntelligenceProfileIndicatorSource, error) {
	vals := map[string]SecurityIntelligenceProfileIndicatorSource{
		"microsoft": SecurityIntelligenceProfileIndicatorSource_Microsoft,
		"osint":     SecurityIntelligenceProfileIndicatorSource_Osint,
		"public":    SecurityIntelligenceProfileIndicatorSource_Public,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIntelligenceProfileIndicatorSource(input)
	return &out, nil
}
