package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreGroupScope string

const (
	TermStoreGroupScopeglobal         TermStoreGroupScope = "Global"
	TermStoreGroupScopesiteCollection TermStoreGroupScope = "SiteCollection"
	TermStoreGroupScopesystem         TermStoreGroupScope = "System"
)

func PossibleValuesForTermStoreGroupScope() []string {
	return []string{
		string(TermStoreGroupScopeglobal),
		string(TermStoreGroupScopesiteCollection),
		string(TermStoreGroupScopesystem),
	}
}

func parseTermStoreGroupScope(input string) (*TermStoreGroupScope, error) {
	vals := map[string]TermStoreGroupScope{
		"global":         TermStoreGroupScopeglobal,
		"sitecollection": TermStoreGroupScopesiteCollection,
		"system":         TermStoreGroupScopesystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TermStoreGroupScope(input)
	return &out, nil
}
