package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BaseSitePagePageLayout string

const (
	BaseSitePagePageLayoutarticle           BaseSitePagePageLayout = "Article"
	BaseSitePagePageLayouthome              BaseSitePagePageLayout = "Home"
	BaseSitePagePageLayoutmicrosoftReserved BaseSitePagePageLayout = "MicrosoftReserved"
)

func PossibleValuesForBaseSitePagePageLayout() []string {
	return []string{
		string(BaseSitePagePageLayoutarticle),
		string(BaseSitePagePageLayouthome),
		string(BaseSitePagePageLayoutmicrosoftReserved),
	}
}

func parseBaseSitePagePageLayout(input string) (*BaseSitePagePageLayout, error) {
	vals := map[string]BaseSitePagePageLayout{
		"article":           BaseSitePagePageLayoutarticle,
		"home":              BaseSitePagePageLayouthome,
		"microsoftreserved": BaseSitePagePageLayoutmicrosoftReserved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BaseSitePagePageLayout(input)
	return &out, nil
}
