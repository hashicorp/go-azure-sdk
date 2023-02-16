package tenantaccessgit

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessIdName string

const (
	AccessIdNameAccess    AccessIdName = "access"
	AccessIdNameGitAccess AccessIdName = "gitAccess"
)

func PossibleValuesForAccessIdName() []string {
	return []string{
		string(AccessIdNameAccess),
		string(AccessIdNameGitAccess),
	}
}

func parseAccessIdName(input string) (*AccessIdName, error) {
	vals := map[string]AccessIdName{
		"access":    AccessIdNameAccess,
		"gitaccess": AccessIdNameGitAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessIdName(input)
	return &out, nil
}
