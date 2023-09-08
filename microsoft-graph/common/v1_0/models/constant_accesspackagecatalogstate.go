package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageCatalogState string

const (
	AccessPackageCatalogStatepublished   AccessPackageCatalogState = "Published"
	AccessPackageCatalogStateunpublished AccessPackageCatalogState = "Unpublished"
)

func PossibleValuesForAccessPackageCatalogState() []string {
	return []string{
		string(AccessPackageCatalogStatepublished),
		string(AccessPackageCatalogStateunpublished),
	}
}

func parseAccessPackageCatalogState(input string) (*AccessPackageCatalogState, error) {
	vals := map[string]AccessPackageCatalogState{
		"published":   AccessPackageCatalogStatepublished,
		"unpublished": AccessPackageCatalogStateunpublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageCatalogState(input)
	return &out, nil
}
