package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BaseSitePagePageLayout string

const (
	BaseSitePagePageLayout_Article           BaseSitePagePageLayout = "article"
	BaseSitePagePageLayout_Home              BaseSitePagePageLayout = "home"
	BaseSitePagePageLayout_MicrosoftReserved BaseSitePagePageLayout = "microsoftReserved"
)

func PossibleValuesForBaseSitePagePageLayout() []string {
	return []string{
		string(BaseSitePagePageLayout_Article),
		string(BaseSitePagePageLayout_Home),
		string(BaseSitePagePageLayout_MicrosoftReserved),
	}
}

func (s *BaseSitePagePageLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBaseSitePagePageLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBaseSitePagePageLayout(input string) (*BaseSitePagePageLayout, error) {
	vals := map[string]BaseSitePagePageLayout{
		"article":           BaseSitePagePageLayout_Article,
		"home":              BaseSitePagePageLayout_Home,
		"microsoftreserved": BaseSitePagePageLayout_MicrosoftReserved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BaseSitePagePageLayout(input)
	return &out, nil
}
