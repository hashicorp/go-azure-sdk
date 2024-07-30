package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityArticleIndicatorSource string

const (
	SecurityArticleIndicatorSource_Microsoft SecurityArticleIndicatorSource = "microsoft"
	SecurityArticleIndicatorSource_Osint     SecurityArticleIndicatorSource = "osint"
	SecurityArticleIndicatorSource_Public    SecurityArticleIndicatorSource = "public"
)

func PossibleValuesForSecurityArticleIndicatorSource() []string {
	return []string{
		string(SecurityArticleIndicatorSource_Microsoft),
		string(SecurityArticleIndicatorSource_Osint),
		string(SecurityArticleIndicatorSource_Public),
	}
}

func (s *SecurityArticleIndicatorSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityArticleIndicatorSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityArticleIndicatorSource(input string) (*SecurityArticleIndicatorSource, error) {
	vals := map[string]SecurityArticleIndicatorSource{
		"microsoft": SecurityArticleIndicatorSource_Microsoft,
		"osint":     SecurityArticleIndicatorSource_Osint,
		"public":    SecurityArticleIndicatorSource_Public,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityArticleIndicatorSource(input)
	return &out, nil
}
