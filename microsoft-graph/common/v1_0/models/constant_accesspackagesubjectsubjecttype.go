package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageSubjectSubjectType string

const (
	AccessPackageSubjectSubjectTypenotSpecified     AccessPackageSubjectSubjectType = "NotSpecified"
	AccessPackageSubjectSubjectTypeservicePrincipal AccessPackageSubjectSubjectType = "ServicePrincipal"
	AccessPackageSubjectSubjectTypeuser             AccessPackageSubjectSubjectType = "User"
)

func PossibleValuesForAccessPackageSubjectSubjectType() []string {
	return []string{
		string(AccessPackageSubjectSubjectTypenotSpecified),
		string(AccessPackageSubjectSubjectTypeservicePrincipal),
		string(AccessPackageSubjectSubjectTypeuser),
	}
}

func parseAccessPackageSubjectSubjectType(input string) (*AccessPackageSubjectSubjectType, error) {
	vals := map[string]AccessPackageSubjectSubjectType{
		"notspecified":     AccessPackageSubjectSubjectTypenotSpecified,
		"serviceprincipal": AccessPackageSubjectSubjectTypeservicePrincipal,
		"user":             AccessPackageSubjectSubjectTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageSubjectSubjectType(input)
	return &out, nil
}
