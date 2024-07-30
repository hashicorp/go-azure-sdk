package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryReviewTagChildSelectability string

const (
	SecurityEdiscoveryReviewTagChildSelectability_Many SecurityEdiscoveryReviewTagChildSelectability = "Many"
	SecurityEdiscoveryReviewTagChildSelectability_One  SecurityEdiscoveryReviewTagChildSelectability = "One"
)

func PossibleValuesForSecurityEdiscoveryReviewTagChildSelectability() []string {
	return []string{
		string(SecurityEdiscoveryReviewTagChildSelectability_Many),
		string(SecurityEdiscoveryReviewTagChildSelectability_One),
	}
}

func (s *SecurityEdiscoveryReviewTagChildSelectability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryReviewTagChildSelectability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryReviewTagChildSelectability(input string) (*SecurityEdiscoveryReviewTagChildSelectability, error) {
	vals := map[string]SecurityEdiscoveryReviewTagChildSelectability{
		"many": SecurityEdiscoveryReviewTagChildSelectability_Many,
		"one":  SecurityEdiscoveryReviewTagChildSelectability_One,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryReviewTagChildSelectability(input)
	return &out, nil
}
