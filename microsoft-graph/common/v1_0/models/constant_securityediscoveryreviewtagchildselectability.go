package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryReviewTagChildSelectability string

const (
	SecurityEdiscoveryReviewTagChildSelectabilityMany SecurityEdiscoveryReviewTagChildSelectability = "Many"
	SecurityEdiscoveryReviewTagChildSelectabilityOne  SecurityEdiscoveryReviewTagChildSelectability = "One"
)

func PossibleValuesForSecurityEdiscoveryReviewTagChildSelectability() []string {
	return []string{
		string(SecurityEdiscoveryReviewTagChildSelectabilityMany),
		string(SecurityEdiscoveryReviewTagChildSelectabilityOne),
	}
}

func parseSecurityEdiscoveryReviewTagChildSelectability(input string) (*SecurityEdiscoveryReviewTagChildSelectability, error) {
	vals := map[string]SecurityEdiscoveryReviewTagChildSelectability{
		"many": SecurityEdiscoveryReviewTagChildSelectabilityMany,
		"one":  SecurityEdiscoveryReviewTagChildSelectabilityOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryReviewTagChildSelectability(input)
	return &out, nil
}
