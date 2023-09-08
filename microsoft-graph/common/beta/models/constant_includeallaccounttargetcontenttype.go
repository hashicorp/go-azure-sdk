package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncludeAllAccountTargetContentType string

const (
	IncludeAllAccountTargetContentTypeaddressBook IncludeAllAccountTargetContentType = "AddressBook"
	IncludeAllAccountTargetContentTypeincludeAll  IncludeAllAccountTargetContentType = "IncludeAll"
	IncludeAllAccountTargetContentTypeunknown     IncludeAllAccountTargetContentType = "Unknown"
)

func PossibleValuesForIncludeAllAccountTargetContentType() []string {
	return []string{
		string(IncludeAllAccountTargetContentTypeaddressBook),
		string(IncludeAllAccountTargetContentTypeincludeAll),
		string(IncludeAllAccountTargetContentTypeunknown),
	}
}

func parseIncludeAllAccountTargetContentType(input string) (*IncludeAllAccountTargetContentType, error) {
	vals := map[string]IncludeAllAccountTargetContentType{
		"addressbook": IncludeAllAccountTargetContentTypeaddressBook,
		"includeall":  IncludeAllAccountTargetContentTypeincludeAll,
		"unknown":     IncludeAllAccountTargetContentTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncludeAllAccountTargetContentType(input)
	return &out, nil
}
