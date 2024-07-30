package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePagePromotionKind string

const (
	SitePagePromotionKind_MicrosoftReserved SitePagePromotionKind = "microsoftReserved"
	SitePagePromotionKind_NewsPost          SitePagePromotionKind = "newsPost"
	SitePagePromotionKind_Page              SitePagePromotionKind = "page"
)

func PossibleValuesForSitePagePromotionKind() []string {
	return []string{
		string(SitePagePromotionKind_MicrosoftReserved),
		string(SitePagePromotionKind_NewsPost),
		string(SitePagePromotionKind_Page),
	}
}

func (s *SitePagePromotionKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSitePagePromotionKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSitePagePromotionKind(input string) (*SitePagePromotionKind, error) {
	vals := map[string]SitePagePromotionKind{
		"microsoftreserved": SitePagePromotionKind_MicrosoftReserved,
		"newspost":          SitePagePromotionKind_NewsPost,
		"page":              SitePagePromotionKind_Page,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SitePagePromotionKind(input)
	return &out, nil
}
