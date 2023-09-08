package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePagePageLayout string

const (
	SitePagePageLayoutarticle           SitePagePageLayout = "Article"
	SitePagePageLayouthome              SitePagePageLayout = "Home"
	SitePagePageLayoutmicrosoftReserved SitePagePageLayout = "MicrosoftReserved"
)

func PossibleValuesForSitePagePageLayout() []string {
	return []string{
		string(SitePagePageLayoutarticle),
		string(SitePagePageLayouthome),
		string(SitePagePageLayoutmicrosoftReserved),
	}
}

func parseSitePagePageLayout(input string) (*SitePagePageLayout, error) {
	vals := map[string]SitePagePageLayout{
		"article":           SitePagePageLayoutarticle,
		"home":              SitePagePageLayouthome,
		"microsoftreserved": SitePagePageLayoutmicrosoftReserved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SitePagePageLayout(input)
	return &out, nil
}
