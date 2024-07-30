package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePagePageLayout string

const (
	SitePagePageLayout_Article           SitePagePageLayout = "article"
	SitePagePageLayout_Home              SitePagePageLayout = "home"
	SitePagePageLayout_MicrosoftReserved SitePagePageLayout = "microsoftReserved"
)

func PossibleValuesForSitePagePageLayout() []string {
	return []string{
		string(SitePagePageLayout_Article),
		string(SitePagePageLayout_Home),
		string(SitePagePageLayout_MicrosoftReserved),
	}
}

func (s *SitePagePageLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSitePagePageLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSitePagePageLayout(input string) (*SitePagePageLayout, error) {
	vals := map[string]SitePagePageLayout{
		"article":           SitePagePageLayout_Article,
		"home":              SitePagePageLayout_Home,
		"microsoftreserved": SitePagePageLayout_MicrosoftReserved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SitePagePageLayout(input)
	return &out, nil
}
