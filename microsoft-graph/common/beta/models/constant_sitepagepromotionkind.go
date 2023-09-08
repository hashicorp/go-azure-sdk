package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePagePromotionKind string

const (
	SitePagePromotionKindmicrosoftReserved SitePagePromotionKind = "MicrosoftReserved"
	SitePagePromotionKindnewsPost          SitePagePromotionKind = "NewsPost"
	SitePagePromotionKindpage              SitePagePromotionKind = "Page"
)

func PossibleValuesForSitePagePromotionKind() []string {
	return []string{
		string(SitePagePromotionKindmicrosoftReserved),
		string(SitePagePromotionKindnewsPost),
		string(SitePagePromotionKindpage),
	}
}

func parseSitePagePromotionKind(input string) (*SitePagePromotionKind, error) {
	vals := map[string]SitePagePromotionKind{
		"microsoftreserved": SitePagePromotionKindmicrosoftReserved,
		"newspost":          SitePagePromotionKindnewsPost,
		"page":              SitePagePromotionKindpage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SitePagePromotionKind(input)
	return &out, nil
}
