package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectSiteAccessType string

const (
	ProtectSiteAccessTypeblock   ProtectSiteAccessType = "Block"
	ProtectSiteAccessTypefull    ProtectSiteAccessType = "Full"
	ProtectSiteAccessTypelimited ProtectSiteAccessType = "Limited"
)

func PossibleValuesForProtectSiteAccessType() []string {
	return []string{
		string(ProtectSiteAccessTypeblock),
		string(ProtectSiteAccessTypefull),
		string(ProtectSiteAccessTypelimited),
	}
}

func parseProtectSiteAccessType(input string) (*ProtectSiteAccessType, error) {
	vals := map[string]ProtectSiteAccessType{
		"block":   ProtectSiteAccessTypeblock,
		"full":    ProtectSiteAccessTypefull,
		"limited": ProtectSiteAccessTypelimited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectSiteAccessType(input)
	return &out, nil
}
