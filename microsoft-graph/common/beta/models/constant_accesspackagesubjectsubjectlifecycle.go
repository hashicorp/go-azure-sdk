package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageSubjectSubjectLifecycle string

const (
	AccessPackageSubjectSubjectLifecyclegoverned    AccessPackageSubjectSubjectLifecycle = "Governed"
	AccessPackageSubjectSubjectLifecyclenotDefined  AccessPackageSubjectSubjectLifecycle = "NotDefined"
	AccessPackageSubjectSubjectLifecyclenotGoverned AccessPackageSubjectSubjectLifecycle = "NotGoverned"
)

func PossibleValuesForAccessPackageSubjectSubjectLifecycle() []string {
	return []string{
		string(AccessPackageSubjectSubjectLifecyclegoverned),
		string(AccessPackageSubjectSubjectLifecyclenotDefined),
		string(AccessPackageSubjectSubjectLifecyclenotGoverned),
	}
}

func parseAccessPackageSubjectSubjectLifecycle(input string) (*AccessPackageSubjectSubjectLifecycle, error) {
	vals := map[string]AccessPackageSubjectSubjectLifecycle{
		"governed":    AccessPackageSubjectSubjectLifecyclegoverned,
		"notdefined":  AccessPackageSubjectSubjectLifecyclenotDefined,
		"notgoverned": AccessPackageSubjectSubjectLifecyclenotGoverned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageSubjectSubjectLifecycle(input)
	return &out, nil
}
