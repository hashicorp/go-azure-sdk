package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountTargetContentType string

const (
	AccountTargetContentTypeaddressBook AccountTargetContentType = "AddressBook"
	AccountTargetContentTypeincludeAll  AccountTargetContentType = "IncludeAll"
	AccountTargetContentTypeunknown     AccountTargetContentType = "Unknown"
)

func PossibleValuesForAccountTargetContentType() []string {
	return []string{
		string(AccountTargetContentTypeaddressBook),
		string(AccountTargetContentTypeincludeAll),
		string(AccountTargetContentTypeunknown),
	}
}

func parseAccountTargetContentType(input string) (*AccountTargetContentType, error) {
	vals := map[string]AccountTargetContentType{
		"addressbook": AccountTargetContentTypeaddressBook,
		"includeall":  AccountTargetContentTypeincludeAll,
		"unknown":     AccountTargetContentTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccountTargetContentType(input)
	return &out, nil
}
